package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
)

func HttpListener(ctx context.Context, lc fx.Lifecycle, config *env.Config, r *gin.Engine) {
	if config.Server.Listen == "" {
		logrus.WithContext(ctx).Debug("no http listener configured")
		return
	}

	h := &http.Server{Handler: r}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "http-server.start")
			defer span.End()

			serviceLn, err := net.Listen("tcp", config.Server.Listen)
			if err != nil {
				logrus.
					WithContext(ctx).
					WithError(err).
					WithField("listen", config.Server.Listen).
					Fatal("failed to listen on http")
				return err
			}

			go func() {
				if err := h.Serve(serviceLn); err != nil && err != http.ErrServerClosed {
					logrus.WithContext(ctx).
						WithError(err).
						Fatal("failed to serve")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "http-server.stop")
			defer span.End()

			return h.Shutdown(ctx)
		},
	})
}

func HttpsListener(ctx context.Context, lc fx.Lifecycle, config *env.Config, r *gin.Engine) {
	if config.Server.CertFile == "" || config.Server.KeyFile == "" {
		logrus.WithContext(ctx).Debug("no cert or key path provided, skipping https listener")
		return
	}

	cert, err := tls.LoadX509KeyPair(config.Server.CertFile, config.Server.KeyFile)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("failed to load cert/key pair")
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			cert,
		},
	}

	h := &http.Server{Handler: r, TLSConfig: tlsConfig}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "http-server.start")
			defer span.End()

			serviceLn, err := tls.Listen("tcp", config.Server.SSLListen, tlsConfig)
			if err != nil {
				logrus.
					WithContext(ctx).
					WithError(err).
					WithField("listen", config.Server.SSLListen).
					Fatal("failed to listen for https")
				return err
			}

			go func() {
				if err := h.Serve(serviceLn); err != nil && err != http.ErrServerClosed {
					logrus.WithContext(ctx).WithError(err).Fatal("failed to serve")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "http-server.stop")
			defer span.End()

			return h.Shutdown(ctx)
		},
	})
}

func UnixListener(ctx context.Context, r *gin.Engine, config *env.Config, lc fx.Lifecycle) {
	if config.Server.UnixSock == "" {
		logrus.WithContext(ctx).Debug("no unix listener configured")
		return
	}

	h := &http.Server{Handler: r}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "http-server.start")
			defer span.End()

			serviceLn, err := net.Listen("unix", config.Server.UnixSock)
			if err != nil {
				logrus.
					WithContext(ctx).
					WithError(err).
					WithField("listen", config.Server.UnixSock).
					Fatal("failed to listen on unix")
				return err
			}

			go func() {
				if err := h.Serve(serviceLn); err != nil && err != http.ErrServerClosed {
					logrus.WithContext(ctx).
						WithError(err).
						Fatal("failed to serve")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, span := tracer.Start(ctx, "http-server.stop")
			defer span.End()

			return h.Shutdown(ctx)
		},
	})
}
