package util

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func JoinIntSlice(slice []int, sep string) string {
	return strings.Join(lo.Map(slice, func(n int, _ int) string {
		return strconv.Itoa(n)
	}), sep)
}
