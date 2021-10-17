package configurationManager

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gobuffalo/packr/v2"
	"strings"
)

var DockerPrefix string = "nickthetramp/"

type ConfigurationManager struct {
	Box    *packr.Box
	Client *client.Client
}

type Configuration struct {
	ID        string
	Name      string
	IsRunning bool
	Path      string
}

func (cm *ConfigurationManager) GetConfigurations() []Configuration {
	configurations := make([]Configuration, 0)

	for _, file := range cm.Box.List() {
		// The parent folder of the docker-compose.yml is set as the name of the Configuration
		fileNames := strings.Split(file, "/")
		if fileNames[1] != "Dockerfile" && fileNames[1] != "docker-compose.yml" {
			continue
		}

		name := strings.ToLower(fileNames[0])

		configuration := Configuration{Name: name, Path: file}
		configuration.GetStatus(*cm)

		configurations = append(configurations, configuration)
	}

	return configurations
}

func (cm *ConfigurationManager) FindConfiguration(name string) (Configuration, error) {
	var configuration Configuration

	for _, cf := range cm.GetConfigurations() {
		if cf.Name == name {
			configuration = cf
		}
	}
	configuration.GetStatus(*cm)

	if &configuration.Path != nil {
		return configuration, nil
	} else {
		return configuration, fmt.Errorf("no configuration found with name: %s", name)
	}
}

func (c *Configuration) GetStatus(cm ConfigurationManager) {
	ctx := context.Background()
	containers, _ := cm.Client.ContainerList(ctx, types.ContainerListOptions{})

	for _, v := range containers {
		// Set default settings
		c.ID = ""
		c.IsRunning = false

		if v.Image == DockerPrefix+c.Name {
			c.IsRunning = true
			c.ID = v.ID
		}
	}
}
