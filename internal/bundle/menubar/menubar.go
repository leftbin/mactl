package menubar

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/tidy/menubar/bartender"
	"github.com/leftbin/mactl/internal/app/tidy/menubar/toothfairy"
)

func Setup() error {
	log.Info("ensuring tooth-fairy")
	if err := toothfairy.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure tooth-fairy")
	}
	log.Info("ensured tooth-fairy")
	log.Info("ensuring bartender")
	if err := bartender.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure bartender")
	}
	log.Info("ensured bartender")
	log.Infof("opening tooth-fairy")
	if err := toothfairy.Open(); err != nil {
		return errors.Wrapf(err, "failed to open tooth-fairy")
	}
	log.Infof("opened tooth-fairy")
	log.Infof("opening bartender")
	if err := bartender.Open(); err != nil {
		return errors.Wrapf(err, "failed to open bartender")
	}
	log.Infof("opened bartender")
	return nil
}
