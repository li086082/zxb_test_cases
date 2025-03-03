package config

import "testing"

func TestNew(t *testing.T) {
	New(".", "config")
	t.Log(Cfg)
}
