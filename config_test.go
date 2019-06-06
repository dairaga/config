package config

import "testing"

type product struct {
	Name  string `json:"name"`
	Sku   uint64 `json:"sku"`
	Color string `json:"color"`
}

func TestConfig(t *testing.T) {
	if _config == nil {
		t.Fatal("default config is nil")
	}
	t.Log(_config.GetString("title"))
	t.Logf("%v", _config.GetBool("database.enabled"))
	t.Logf("%v", _config.GetIntSlice("database.ports"))

	products := []product{}

	if err := _config.GetObject("products", &products); err != nil {
		t.Errorf("get products: %v", err)
	} else {
		t.Log(products)
	}

	t.Logf("%v, %T", _config.Get("clients.data"), _config.Get("clients.data"))
}
