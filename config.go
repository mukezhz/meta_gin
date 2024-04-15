package meta_gin

import "github.com/BurntSushi/toml"

type RouteConf struct {
	Role    string   `toml:"role"`
	Path    string   `toml:"path"`
	Methods []string `toml:"methods"`
}
type Config struct {
	Roles  map[string][]string `toml:"roles"`
	Routes []RouteConf         `toml:"routes"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
