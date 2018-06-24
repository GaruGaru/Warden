package api

import (
	"github.com/GaruGaru/Warden/agent"
	"github.com/gin-gonic/gin"
)

type HostApi struct {
	HostInfo agent.HostInfoFetcher
}

func (api HostApi) Handler(c *gin.Context) {
	info, err := api.HostInfo.Fetch()
	if err != nil {
		c.JSON(500, gin.H{"info": agent.HostInfo{}, "error": err.Error()})
	} else {
		c.JSON(200, gin.H{"info": info, "error": nil})
	}
}

func probe(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
