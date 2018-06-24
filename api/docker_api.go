package api

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/GaruGaru/Warden/agent"
)

type DockerApi struct {
	DockerClient client.Client
}

func NewDockerApi() (DockerApi, error) {
	clint, err := client.NewEnvClient()

	if err != nil {
		return DockerApi{}, err
	}

	return DockerApi{DockerClient: *clint,}, nil
}

type DockerNode struct {
	ID         string
	Name       string
	Role       string
	Status     string
	Ip         string
	Containers []DockerContainer
}

type DockerContainer struct {
	ID        string
	Name      string
	Image     string
	State     string
	CreatedAt int64
}

func (a DockerApi) Handler(c *gin.Context) {
	ctx := context.Background()

	nodesMap, err := a.GetNodeMapFromLocal(ctx)

	if err != nil {
		c.JSON(500, gin.H{"info": agent.HostInfo{}, "error": err.Error()})
	} else {
		c.JSON(200, gin.H{"nodes": nodesMap, "error": nil})
	}

}

func (a DockerApi) GetNodeMapFromLocal(ctx context.Context) ([]DockerNode, error) {
	nodes := make([]DockerNode, 1)

	containers, err := a.DockerClient.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		return []DockerNode{}, err
	}

	containersList := make([]DockerContainer, len(containers))

	for i, container := range containers {
		containersList[i] = DockerContainer{
			ID:        container.ID,
			Name:      container.Names[0],
			Image:     container.Image,
			State:     container.State,
			CreatedAt: container.Created,
		}
	}

	nodes[0] = DockerNode{
		ID:         "0",
		Name:       "localhost",
		Role:       "master",
		Status:     "RUNNING",
		Ip:         "127.0.0.1",
		Containers: containersList,
	}
	return nodes, nil

}
