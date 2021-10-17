package configurationManager

import (
	"archive/tar"
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"io"
	"log"
	"os"
)

func (c *Configuration) BuildContainer(cm ConfigurationManager) {
	ctx := context.Background()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	readDockerFile, err := cm.Box.Find(c.Path)
	if err != nil {
		log.Println(err, " :unable to read Dockerfile")
	}

	tarHeader := &tar.Header{
		Name: DockerPrefix + c.Name,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Println(err, " :unable to write tar header")
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Println(err, " :unable to write tar body")
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cm.Client.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context:    dockerFileTarReader,
			Dockerfile: DockerPrefix + c.Name,
			Remove:     true,
		},
	)
	if err != nil {
		log.Println(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}

	log.Println("build image")
	log.Println(imageBuildResponse.Body)
	//c.CreateContainer(cm)
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
