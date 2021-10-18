package configurationManager

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/archive"
	"log"
)

var configurationsFolder = "./configurations/"

func (c *Configuration) BuildContainer(cm ConfigurationManager) {
	ctx := context.Background()

	containerDir, err := archive.TarWithOptions(configurationsFolder + c.Name, &archive.TarOptions{})
	if err != nil {
		log.Println(err, " :unable to write tar body")
	}

	imageBuildResponse, err := cm.Client.ImageBuild(
		ctx,
		containerDir,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Remove:     true,
			Tags: []string{DockerPrefix + c.Name},
		},
	)
	defer imageBuildResponse.Body.Close()
	//_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	//if err != nil {
	//	log.Println(err, " :unable to read image build response")
	//}
	//log.Println(imageBuildResponse.Body)

	c.CreateContainer(cm)
}

func (c *Configuration) CreateContainer(cm ConfigurationManager) {
	ctx := context.Background()

	createdContainer, err := cm.Client.ContainerCreate(ctx, &container.Config{
		Image: DockerPrefix + c.Name,
	}, nil, nil, nil, "")
	if err != nil {
		log.Println(err.Error())
	}

	if err := cm.Client.ContainerStart(ctx, createdContainer.ID, types.ContainerStartOptions{}); err != nil {
		log.Println(err.Error())
	}
}

func (c *Configuration) StopContainer(cm ConfigurationManager) {
	ctx := context.Background()

	if c.ID == "" {
		return
	}

	if err := cm.Client.ContainerStop(ctx, c.ID, nil); err != nil {
		log.Println(err.Error())
	}
	c.GetStatus(cm)
}
