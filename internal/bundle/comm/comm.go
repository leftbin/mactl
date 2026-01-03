package comm

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/comm/slack"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring slack")
	if err := slack.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure slack")
	}
	log.Info("ensured slack")
	return nil
}
