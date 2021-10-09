package configurationManager

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"strings"
)

type ConfigurationManager struct {
	Box *packr.Box
}

type Configuration struct {
	Name      string `json:"name"`
	IsRunning bool   `json:"isRunning"`
	Path      string `json:"path"`
}

func (cm *ConfigurationManager) GetConfigurations() []Configuration {
	configurations := make([]Configuration, 0)

	for _, file  := range cm.Box.List() {
		// The parent folder of the docker-compose.yml is set as the name of the Configuration
		name := strings.Title(strings.Split(file, "/")[0])
		configurations = append(configurations, Configuration{Name: name, Path: file})
	}

	return configurations
}

func (cm *ConfigurationManager) FindConfiguration(name string) (Configuration, error) {
	var configuration Configuration

	for _, cf := range cm.GetConfigurations(){
		if cf.Name == name {
			configuration = cf
		}
	}

	if &configuration.Path != nil {
		return configuration, nil
	} else {
		return configuration, fmt.Errorf("no configuration found with name: %s", name)
	}
}
