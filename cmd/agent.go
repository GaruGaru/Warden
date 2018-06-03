package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	"strings"
	"github.com/GaruGaru/Warden/metrics"
	"fmt"
	"github.com/GaruGaru/Warden/agent"
)

var (
	reporterTypes = []string{"logger", "statsd"}
)

func init() {
	agentCmd.Flags().String("reporter", reporterTypes[0], strings.Join(reporterTypes, ","))

	agentCmd.Flags().String("statsd_host", "localhost:8125", "statsd host:port")
	agentCmd.Flags().String("statsd_prefix", "warden", "statsd metrix prefix")

	viper.BindPFlags(agentCmd.Flags())
	rootCmd.AddCommand(agentCmd)
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start agent for local host monitoring",
	Run: func(cmd *cobra.Command, args []string) {

		log.WithField("reporter", viper.Get("reporter")).
			Info("Warden agent started")

		reporter, err := createReporter()

		if err != nil{
			log.Error("error creating reporter: ", err)
			return
		}

		hostInfoFetcher := agent.DefaultHostInfoFetcher{}

		info, err := hostInfoFetcher.Fetch()

		if err != nil {
			log.Error("error fetching host info: ", err)
			return
		}

		reporter.Send(info)


	},
}

func createReporter() (metrics.MetricsReporter, error) {

	reporter := viper.GetString("reporter")

	if reporter == reporterTypes[0] {
		return metrics.MetricsReporterLogger{}, nil
	} else if reporter == reporterTypes[1] {
		return metrics.NewStatsdMetricsReporter(viper.GetString("statsd_host"), viper.GetString("statsd_prefix"))
	} else {
		return nil, fmt.Errorf("reporter of type %s not supported", reporter)
	}

}

