package config

import (
	"encoding/json"
	"fmt"
	"time"

	toml "github.com/pelletier/go-toml"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	x *toml.Tree
}

var (
	_config *Config
)

func init() {
	t, err := toml.LoadFile("config.toml")
	if err == nil {
		_config = &Config{t}
	}
}

// ----------------------------------------------------------------------------

// Get ...
func (c *Config) Get(key string) interface{} {
	return c.x.Get(key)
}

// GetInt ...
func (c *Config) GetInt(key string, def ...int) int {
	z := int(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToInt(x)
}

// GetIntSlice ...
func (c *Config) GetIntSlice(key string) []int {
	return cast.ToIntSlice(c.x.Get(key))

}

// GetUint ...
func (c *Config) GetUint(key string, def ...uint) uint {
	z := uint(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToUint(x)
}

// GetInt8 ...
func (c *Config) GetInt8(key string, def ...int8) int8 {
	z := int8(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToInt8(x)
}

// GetUint8 ...
func (c *Config) GetUint8(key string, def ...uint8) uint8 {
	z := uint8(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToUint8(x)
}

// GetInt16 ...
func (c *Config) GetInt16(key string, def ...int16) int16 {
	z := int16(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToInt16(x)
}

// GetUint16 ...
func (c *Config) GetUint16(key string, def ...uint16) uint16 {
	z := uint16(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToUint16(x)
}

// GetInt32 ...
func (c *Config) GetInt32(key string, def ...int32) int32 {
	z := int32(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToInt32(x)
}

// GetUint32 ...
func (c *Config) GetUint32(key string, def ...uint32) uint32 {
	z := uint32(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToUint32(x)
}

// GetInt64 ...
func (c *Config) GetInt64(key string, def ...int64) int64 {
	z := int64(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToInt64(x)
}

// GetUint64 ...
func (c *Config) GetUint64(key string, def ...uint64) uint64 {
	z := uint64(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToUint64(x)
}

// GetFloat32 ...
func (c *Config) GetFloat32(key string, def ...float32) float32 {
	z := float32(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToFloat32(x)
}

// GetFloat64 ...
func (c *Config) GetFloat64(key string, def ...float64) float64 {
	z := float64(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToFloat64(x)
}

// GetString ...
func (c *Config) GetString(key string, def ...string) string {
	z := ""
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToString(x)
}

// GetStringSlice ...
func (c *Config) GetStringSlice(key string) []string {
	return cast.ToStringSlice(c.x.Get(key))
}

// GetBool ...
func (c *Config) GetBool(key string, def ...bool) bool {
	z := false
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToBool(x)
}

// GetBoolSlice ...
func (c *Config) GetBoolSlice(key string) []bool {
	return cast.ToBoolSlice(c.x.Get(key))
}

// GetTime ...
func (c *Config) GetTime(key string, def ...time.Time) time.Time {
	z := time.Time{}
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToTime(x)
}

// GetDuration ...
func (c *Config) GetDuration(key string, def ...time.Duration) time.Duration {
	z := time.Duration(0)
	if len(def) > 0 {
		z = def[0]
	}

	x := c.x.GetDefault(key, z)
	return cast.ToDuration(x)
}

// GetDurationSlice ...
func (c *Config) GetDurationSlice(key string) []time.Duration {
	return cast.ToDurationSlice(c.x.Get(key))
}

// GetMap ...
func (c *Config) GetMap(key string) map[string]interface{} {
	x := c.x.Get(key)
	if x == nil {
		return nil
	}
	switch v := x.(type) {
	case *toml.Tree:
		return v.ToMap()
	default:
		return nil
	}
}

// GetObject ...
func (c *Config) GetObject(key string, x interface{}) error {
	a := c.x.Get(key)
	if a == nil {
		return fmt.Errorf("%s not found", key)
	}

	var tmp interface{}

	switch v := a.(type) {
	case *toml.Tree:
		tmp = v.ToMap()
	case []*toml.Tree:
		s := make([]map[string]interface{}, len(v))
		for i, m := range v {
			s[i] = m.ToMap()
		}

		tmp = s
	default:
		tmp = v
	}

	tmpBytes, err := json.Marshal(tmp)
	if err != nil {
		return err
	}

	return json.Unmarshal(tmpBytes, x)

}

// ----------------------------------------------------------------------------

// Get ...
func Get(key string) interface{} {
	return _config.Get(key)
}

// GetInt ...
func GetInt(key string, def ...int) int {
	return _config.GetInt(key, def...)
}

// GetIntSlice ...
func GetIntSlice(key string) []int {
	return _config.GetIntSlice(key)

}

// GetUint ...
func GetUint(key string, def ...uint) uint {
	return _config.GetUint(key, def...)
}

// GetInt8 ...
func GetInt8(key string, def ...int8) int8 {
	return _config.GetInt8(key, def...)
}

// GetUint8 ...
func GetUint8(key string, def ...uint8) uint8 {
	return _config.GetUint8(key, def...)
}

// GetInt16 ...
func GetInt16(key string, def ...int16) int16 {
	return _config.GetInt16(key, def...)
}

// GetUint16 ...
func GetUint16(key string, def ...uint16) uint16 {
	return _config.GetUint16(key, def...)
}

// GetInt32 ...
func GetInt32(key string, def ...int32) int32 {
	return _config.GetInt32(key, def...)
}

// GetUint32 ...
func GetUint32(key string, def ...uint32) uint32 {
	return _config.GetUint32(key, def...)
}

// GetInt64 ...
func GetInt64(key string, def ...int64) int64 {
	return _config.GetInt64(key, def...)
}

// GetUint64 ...
func GetUint64(key string, def ...uint64) uint64 {
	return _config.GetUint64(key, def...)
}

// GetFloat32 ...
func GetFloat32(key string, def ...float32) float32 {
	return _config.GetFloat32(key, def...)
}

// GetFloat64 ...
func GetFloat64(key string, def ...float64) float64 {
	return _config.GetFloat64(key, def...)
}

// GetString ...
func GetString(key string, def ...string) string {
	return _config.GetString(key, def...)
}

// GetStringSlice ...
func GetStringSlice(key string) []string {
	return _config.GetStringSlice(key)
}

// GetBool ...
func GetBool(key string, def ...bool) bool {
	return _config.GetBool(key, def...)
}

// GetBoolSlice ...
func GetBoolSlice(key string) []bool {
	return _config.GetBoolSlice(key)
}

// GetTime ...
func GetTime(key string, def ...time.Time) time.Time {
	return _config.GetTime(key, def...)
}

// GetDuration ...
func GetDuration(key string, def ...time.Duration) time.Duration {
	return _config.GetDuration(key, def...)
}

// GetDurationSlice ...
func GetDurationSlice(key string) []time.Duration {
	return _config.GetDurationSlice(key)
}

// GetMap ...
func GetMap(key string) map[string]interface{} {
	return _config.GetMap(key)
}

// GetObject ...
func GetObject(key string, x interface{}) error {
	return _config.GetObject(key, x)
}
