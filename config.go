package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"
)

const (
	ENV_DELIM  = "_"
	FLAG_DELIM = "-"
	KEY_DELIM  = "."

	DEFAULT_FILE = `config.toml`
)

var (
	ErrNoCandidates = errors.New("no candidates")
	ErrKeyNotFound  = errors.New("key not found")
)

// -----------------------------------------------------------------------------

type Object = map[string]any

// -----------------------------------------------------------------------------

type Config struct {
	src   Object
	delim string
}

// -----------------------------------------------------------------------------

func ConfigGet[T any](fn func(any, ...T) T, config map[string]any, path string, delim string, def ...T) T {
	val := Get(config, path, delim)
	return fn(val, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetBool(key string, def ...bool) bool {
	return ConfigGet(CastBool, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetString(key string, def ...string) string {
	return ConfigGet(CastString, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetDuration(key string, def ...time.Duration) time.Duration {
	return ConfigGet(CastDuration, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetInt(key string, def ...int) int {
	return ConfigGet(CastInt, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetInt8(key string, def ...int8) int8 {
	return ConfigGet(CastInt8, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetInt16(key string, def ...int16) int16 {
	return ConfigGet(CastInt16, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetInt32(key string, def ...int32) int32 {
	return ConfigGet(CastInt32, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetInt64(key string, def ...int64) int64 {
	return ConfigGet(CastInt64, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetUint(key string, def ...uint) uint {
	return ConfigGet(CastUint, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetUint8(key string, def ...uint8) uint8 {
	return ConfigGet(CastUint8, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetUint16(key string, def ...uint16) uint16 {
	return ConfigGet(CastUint16, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetUint32(key string, def ...uint32) uint32 {
	return ConfigGet(CastUint32, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetUint64(key string, def ...uint64) uint64 {
	return ConfigGet(CastUint64, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetFloat32(key string, def ...float32) float32 {
	return ConfigGet(CastFloat32, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetFloat64(key string, def ...float64) float64 {
	return ConfigGet(CastFloat64, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetObject(key string, def ...Object) Object {
	return ConfigGet(CastObject, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetTime(key string, def ...time.Time) time.Time {
	return ConfigGet(CastTime, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetIntSlice(key string, def ...[]int) []int {
	return ConfigGet(CastIntSlice, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) GetSlice(key string, def ...[]any) []any {
	return ConfigGet(CastSlice, c.src, key, c.delim, def...)
}

// -----------------------------------------------------------------------------

func (c *Config) Unmarshal(key string, v any) error {
	obj := c.GetObject(key)
	if obj == nil {
		return ErrKeyNotFound
	}

	objBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(objBytes, v)
}

// -----------------------------------------------------------------------------

func Load(bindEnvAndFlag ...bool) (*Config, error) {
	candidates := []string{
		DEFAULT_FILE,
		filepath.Join("/etc", strings.ToLower(os.Args[0]), DEFAULT_FILE),
	}

	fileBytes, err := OpenCandidate(candidates)
	if err != nil {
		return nil, err
	}

	src := make(map[string]any)
	err = toml.Unmarshal(fileBytes, &src)
	if err != nil {
		return nil, err
	}

	// default bind env and flags
	bind := len(bindEnvAndFlag) <= 0 || bindEnvAndFlag[0]

	if bind {
		// bind env first
		BindEnv(src)

		// bind flags to override env
		BindFlag(src)
	}

	return &Config{
		src:   src,
		delim: KEY_DELIM,
	}, nil
}
