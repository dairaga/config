package config

import (
	"time"

	"github.com/spf13/cast"
)

func _cast[T any](fn func(any) (T, error), val any, def ...T) T {
	result, err := fn(val)
	if err != nil && len(def) > 0 {
		return def[0]
	}
	return result
}

// -----------------------------------------------------------------------------

func CastInt(val any, def ...int) int {
	return _cast(cast.ToIntE, val, def...)
}

// -----------------------------------------------------------------------------

func CastInt8(val any, def ...int8) int8 {
	return _cast(cast.ToInt8E, val, def...)
}

// -----------------------------------------------------------------------------

func CastInt16(val any, def ...int16) int16 {
	return _cast(cast.ToInt16E, val, def...)
}

// -----------------------------------------------------------------------------

func CastInt32(val any, def ...int32) int32 {
	return _cast(cast.ToInt32E, val, def...)
}

// -----------------------------------------------------------------------------

func CastInt64(val any, def ...int64) int64 {
	return _cast(cast.ToInt64E, val, def...)
}

// -----------------------------------------------------------------------------

func CastDuration(val any, def ...time.Duration) time.Duration {
	return _cast(cast.ToDurationE, val, def...)
}

// -----------------------------------------------------------------------------

func CastUint(val any, def ...uint) uint {
	return _cast(cast.ToUintE, val, def...)
}

// -----------------------------------------------------------------------------

func CastUint8(val any, def ...uint8) uint8 {
	return _cast(cast.ToUint8E, val, def...)
}

// -----------------------------------------------------------------------------

func CastUint16(val any, def ...uint16) uint16 {
	return _cast(cast.ToUint16E, val, def...)
}

// -----------------------------------------------------------------------------

func CastUint32(val any, def ...uint32) uint32 {
	return _cast(cast.ToUint32E, val, def...)
}

// -----------------------------------------------------------------------------

func CastUint64(val any, def ...uint64) uint64 {
	return _cast(cast.ToUint64E, val, def...)
}

// -----------------------------------------------------------------------------

func CastFloat32(val any, def ...float32) float32 {
	return _cast(cast.ToFloat32E, val, def...)
}

// -----------------------------------------------------------------------------

func CastFloat64(val any, def ...float64) float64 {
	return _cast(cast.ToFloat64E, val, def...)
}

// -----------------------------------------------------------------------------

func CastBool(val any, def ...bool) bool {
	return _cast(cast.ToBoolE, val, def...)
}

// -----------------------------------------------------------------------------

func CastString(val any, def ...string) string {
	return _cast(cast.ToStringE, val, def...)
}

// -----------------------------------------------------------------------------

func CastTime(val any, def ...time.Time) time.Time {
	return _cast(cast.ToTimeE, val, def...)
}

// -----------------------------------------------------------------------------

func CastStringMapString(val any, def ...map[string]string) map[string]string {
	return _cast(cast.ToStringMapStringE, val, def...)
}

// -----------------------------------------------------------------------------

func CastStringMapStringSlice(val any, def ...map[string][]string) map[string][]string {
	return _cast(cast.ToStringMapStringSliceE, val, def...)
}

// -----------------------------------------------------------------------------

func CastStringMapBool(val any, def ...map[string]bool) map[string]bool {
	return _cast(cast.ToStringMapBoolE, val, def...)
}

// -----------------------------------------------------------------------------

func CastStringMap(val any, def ...map[string]any) map[string]any {
	return _cast(cast.ToStringMapE, val, def...)
}

// -----------------------------------------------------------------------------

func CastObject(val any, def ...Object) Object {
	return CastStringMap(val, def...)
}

// -----------------------------------------------------------------------------

func CastStringMapInt(val any, def ...map[string]int) map[string]int {
	return _cast(cast.ToStringMapIntE, val, def...)
}

// -----------------------------------------------------------------------------

func CastStringMapInt64(val any, def ...map[string]int64) map[string]int64 {
	return _cast(cast.ToStringMapInt64E, val, def...)
}

// -----------------------------------------------------------------------------

func CastSlice(val any, def ...[]any) []any {
	return _cast(cast.ToSliceE, val, def...)
}

// -----------------------------------------------------------------------------

func CastBoolSlice(val any, def ...[]bool) []bool {
	return _cast(cast.ToBoolSliceE, val, def...)
}

// -----------------------------------------------------------------------------

func CastStringSlice(val any, def ...[]string) []string {
	return _cast(cast.ToStringSliceE, val, def...)
}

// -----------------------------------------------------------------------------

func CastIntSlice(val any, def ...[]int) []int {
	return _cast(cast.ToIntSliceE, val, def...)
}

// -----------------------------------------------------------------------------

func CastDurationSlice(val any, def ...[]time.Duration) []time.Duration {
	return _cast(cast.ToDurationSliceE, val, def...)
}

// -----------------------------------------------------------------------------

func StringToDate(s string) (time.Time, error) {
	return cast.StringToDate(s)
}

// -----------------------------------------------------------------------------

func StringToDateInDefaultLocation(s string, loc *time.Location) (time.Time, error) {
	return cast.StringToDateInDefaultLocation(s, loc)
}
