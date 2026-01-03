package audio

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/audio/krisp"
	"github.com/leftbin/mactl/internal/app/audio/shush"
)

func Setup() error {
	log.Info("ensuring shush")
	if err := shush.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure shush")
	}
	log.Info("ensured shush")
	log.Info("ensuring krisp")
	if err := krisp.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure krisp")
	}
	log.Info("ensured krisp")
	log.Infof("opening shush")
	if err := shush.Open(); err != nil {
		return errors.Wrap(err, "failed to open shush")
	}
	log.Infof("opened shush")
	log.Infof("opening krisp")
	if err := krisp.Open(); err != nil {
		return errors.Wrap(err, "failed to open krisp")
	}
	log.Infof("opened krisp")
	return nil
}
