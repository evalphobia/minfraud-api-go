package minfraud

import (
	"github.com/evalphobia/minfraud-api-go/client"
	"github.com/evalphobia/minfraud-api-go/config"
	"github.com/evalphobia/minfraud-api-go/log"
)

// MinFraud is service struct for MinFraud API.
type MinFraud struct {
	client *client.Client
	logger log.Logger
}

// New creates MinFraud from Config data.
func New(conf config.Config) (*MinFraud, error) {
	cli, err := conf.Client()
	if err != nil {
		return nil, err
	}

	return &MinFraud{
		client: cli,
		logger: log.DefaultLogger,
	}, nil
}

// SetLogger sets internal API logger.
func (s *MinFraud) SetLogger(logger log.Logger) {
	s.logger = logger
}
