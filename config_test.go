package config_test

import (
	"testing"

	"os"

	"github.com/dairaga/config"
)

type product struct {
	Name  string `json:"name"`
	Sku   uint64 `json:"sku"`
	Color string `json:"color"`
}

func TestConfig(t *testing.T) {
	os.Setenv("ENV_MYSQL_HOST", "localhost")
	config.BindEnv("ENV")
	t.Log(config.GetString("mysql.host"))

	t.Log(config.GetString("title"))
	t.Logf("%v", config.GetBool("database.enabled"))
	t.Logf("%v", config.GetIntSlice("database.ports"))

	products := []product{}

	if err := config.GetObject("products", &products); err != nil {
		t.Errorf("get products: %v", err)
	} else {
		t.Log(products)
	}

	t.Logf("%v, %T", config.Get("clients.data"), config.Get("clients.data"))
	t.Logf("%v, %T", config.GetMap("clients"), config.GetMap("clients"))

}
