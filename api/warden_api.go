package api

import "github.com/gin-gonic/gin"

type WardenApi struct{
	HostApi HostApi
	DockerApi DockerApi
}

func (api WardenApi) Serve() {
	r := gin.Default()
	r.GET("/probe", probe)
	r.GET("/host", api.HostApi.Handler)
	r.GET("/docker", api.DockerApi.Handler)
	r.Run()
}

