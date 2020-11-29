package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfiguration(t *testing.T) {
	t.Run("Load YAML Config", func(t *testing.T) {
		fn := "config_stub.yaml"
		conf := new(Config)
		err := Load(fn, conf)
		assert.NoError(t, err)
	})
	t.Run("Load TOML Config", func(t *testing.T) {
		t.Skip("TODO: Skipping test for now")

		fn := "config_stub.toml"
		conf := new(Config)
		err := Load(fn, conf)
		assert.NoError(t, err)
	})
	t.Run("Load nonexistent Config", func(t *testing.T) {
		t.Skip("TODO: Skipping test for now")
		fn := "config_stub"
		conf := new(Config)
		err := Load(fn, conf)
		assert.Error(t, err)
	})
}
