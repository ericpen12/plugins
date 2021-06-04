package db

import (
	"github.com/quzhen12/plugins/settings"
	"testing"
)

func TestConnect(t *testing.T) {
	settings.Init()
	err := Connect()
	if err != nil {
		t.Error(err)
	}
}
