package cmd

import (
	"github.com/GaruGaru/Warden/agent"
	"github.com/GaruGaru/Warden/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Info("Api server started")

		dockerApi, err := api.NewDockerApi()

		if err != nil {
			panic(err)
		}

		hostApi := api.HostApi{HostInfo: agent.DefaultHostInfoFetcher{}}

		api := api.WardenApi{
			HostApi:   hostApi,
			DockerApi: dockerApi,
		}

		api.Serve()

	},
}
