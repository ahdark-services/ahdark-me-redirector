package trace

import (
	"context"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/uptrace-go/extra/otellogrus"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
)

func NewTracerProvider(ctx context.Context, config *env.Config) (*tracesdk.TracerProvider, error) {
	ctx, span := tracer.Start(ctx, "new-tracer-provider")
	defer span.End()

	logger := logrus.WithContext(ctx)

	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.System.TracerDSN)))
	if err != nil {
		logger.WithError(err).Fatal("failed to create jaeger exporter")
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.System.Name),
			semconv.DeploymentEnvironmentKey.String(lo.If(config.System.Debug, "development").Else("production")),
			semconv.ServiceInstanceIDKey.String(env.InstanceID),
		)),
	)

	return tp, nil
}

func UseTracerProvider(ctx context.Context, tp *tracesdk.TracerProvider, logger *logrus.Logger) {
	ctx, span := tracer.Start(ctx, "use-tracer-provider")
	defer span.End()

	otel.SetTracerProvider(tp)

	logger.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	)))
}
