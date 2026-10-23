package config

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestParseMaxConns(t *testing.T) {
	const key = "EVOGO_TEST_MAX_CONNS"

	cases := map[string]int{
		"":    defaultMaxConns,
		"40":  40,
		"0":   defaultMaxConns,
		"-3":  defaultMaxConns,
		"abc": defaultMaxConns,
	}
	for value, want := range cases {
		t.Setenv(key, value)
		if got := parseMaxConns(key); got != want {
			t.Errorf("parseMaxConns(%q) = %d, want %d", value, got, want)
		}
	}
}

func TestConfigurePoolLimitsOpenConns(t *testing.T) {
	// sql.Open não conecta; só o pool é configurado.
	db, err := sql.Open("postgres", "host=127.0.0.1 dbname=none sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ConfigurePool(db, 7)
	if got := db.Stats().MaxOpenConnections; got != 7 {
		t.Errorf("MaxOpenConnections = %d, want 7", got)
	}

	ConfigurePool(db, 0)
	if got := db.Stats().MaxOpenConnections; got != defaultMaxConns {
		t.Errorf("MaxOpenConnections = %d, want %d", got, defaultMaxConns)
	}
}
