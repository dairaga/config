package config_test

import (
	"flag"
	"os"
	"testing"
	"time"

	config2 "github.com/dairaga/config/v2"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	os.Setenv("HELLO_KEY1", "1000")
	os.Setenv("HELLO_KEY2", "2000")

	flag.String("hello-key2", "102", "key2")
	flag.Int("hello-key3", 0, "key3")
	flag.Set("hello-key3", "103")

	config, err := config2.Load()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "TOML Example", config.GetString("title"))

	t.Log(config.GetTime("owner.dob"))

	ansTime, err := time.Parse("2006-01-02T15:04:05-07:00", "1979-05-27T07:32:00+08:00")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, ansTime.UnixMilli(), config.GetTime("owner.dob").UnixMilli())

	assert.Equal(t, []int{8001, 8001, 8002}, config.GetIntSlice("database.ports"))
	assert.Equal(t, int64(5000), config.GetInt64("database.connection_max"))

	t.Log(config.GetSlice("clients.data"))

	assert.Equal(t, []string{"gamma", "delta"}, config2.CastStringSlice(config.GetSlice("clients.data")[0]))

	t.Log(config.GetObject("servers.alpha"))
	t.Log(config.GetSlice("products"))
	t.Log(config.GetSlice("fruit"))

	dur := config.GetDuration("owner.duration")
	assert.Equal(t, 100*time.Second, dur)

	assert.Equal(t, 1000, config.GetInt("hello.key1"))
	assert.Equal(t, 102, config.GetInt("hello.key2"))
	assert.Equal(t, 103, config.GetInt("hello.key3"))

	assert.Equal(t, true, config.GetBool("database.enabled"))
	assert.Equal(t, false, config.GetBool("database.enabled2"))
}
