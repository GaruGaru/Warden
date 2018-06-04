package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Name    = "Warden"
	Build   = "wip"
	Version = "0"
)

func init() {
	viper.AutomaticEnv()
}

var rootCmd = &cobra.Command{
	Use:   "warden",
	Short: "Warden is a fast and lightweight host monitor",
	Long:  `Warden is a fast and lightweight host monitor built for containers`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Warden is a fast and lightweight host monitor built for containers")
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
