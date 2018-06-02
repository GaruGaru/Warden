package metrics

import (
	"github.com/GaruGaru/Warden/agent"
	"github.com/sirupsen/logrus"
	"encoding/json"
)

type MetricsSender interface {
	Send(info agent.AgentInfo) error
}

type MetricsLogger struct {
}

func (l MetricsLogger) Send(info agent.AgentInfo) error {
	infoJson, err := json.Marshal(info)

	if err != nil {
		return err
	}

	logrus.Info(string(infoJson))

	return nil
}
