package settings

import "testing"

func TestSetConfigPath(t *testing.T) {
	tests := []struct {
		path       string
		configName string
		configType string
		configPath string
	}{
		{"config.yaml", "config", "yaml", "."},
		{"./config.yaml", "config", "yaml", "."},
		{"../config.yaml", "config", "yaml", ".."},
		{"./1/config.yaml", "config", "yaml", "1"},
		{"./1/2/config.yaml", "config", "yaml", "1/2"},
	}
	for _, tt := range tests {
		SetConfigPath(tt.path)
		if configName != tt.configName {
			t.Errorf("unexcept configName: %s, tt: %+v", configName, tt)
		}
		if configType != tt.configType {
			t.Errorf("unexcept configType: %s, tt: %+v", configType, tt)
		}
		if configPath != tt.configPath {
			t.Errorf("unexcept configPath: %s, tt: %+v", configPath, tt)
		}
	}
}
