package metrics

import (
	"encoding/json"
	"github.com/GaruGaru/Warden/agent"
	"github.com/sirupsen/logrus"
)

type MetricsReporter interface {
	Send(info agent.AgentInfo) error
}

type MetricsReporterLogger struct {
}

func (l MetricsReporterLogger) Send(info agent.AgentInfo) error {
	infoJson, err := json.Marshal(info)

	if err != nil {
		return err
	}

	logrus.Info(string(infoJson))

	return nil
}
