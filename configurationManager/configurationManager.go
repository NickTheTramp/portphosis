package configurationManager

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gobuffalo/packr/v2"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

var DockerPrefix = "nickthetramp/"

type ConfigurationManager struct {
	Box    *packr.Box
	Client *client.Client
}

type Configuration struct {
	ID        string
	Name      string
	IsRunning bool
	Path      string
	Ports     []int `yaml:"ports"`
	Volumes   []struct {
		Target string `yaml:"target"`
		Source string `yaml:"source"`
	} `yaml:"volumes"`
}

func (cm *ConfigurationManager) GetConfigurations() []Configuration {
	configurations := make([]Configuration, 0)

	for _, file := range cm.Box.List() {
		// The parent folder of the docker-compose.yml is set as the name of the Configuration
		fileNames := strings.Split(file, "/")

		if len(fileNames) <= 1 {
			continue
		}

		if fileNames[1] != "Dockerfile" && fileNames[1] != "docker-compose.yml" {
			continue
		}

		name := strings.ToLower(fileNames[0])

		configuration := Configuration{Name: name, Path: file}
		configuration.Populate(*cm)

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
	configuration.Populate(*cm)

	if &configuration.Path != nil {
		return configuration, nil
	} else {
		return configuration, fmt.Errorf("no configuration found with name: %s", name)
	}
}

func (c *Configuration) Populate(cm ConfigurationManager) {
	ctx := context.Background()
	containers, _ := cm.Client.ContainerList(ctx, types.ContainerListOptions{})

	// Get Porthosis Configuration File
	configurationYaml, err := cm.Box.FindString(c.Name + "/porthosis.yml")
	if err != nil {
		log.Println("Error while finding file", err.Error())
	}

	err = yaml.Unmarshal([]byte(configurationYaml), &c)
	if err != nil {
		log.Println("Error while unmarshalling", err.Error())
	}

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
