package tilt

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "tilt"
)

func Setup() error {
	log.Info("installing tilt")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install tilt")
	}
	log.Info("installed tilt")
	return nil
}

func Upgrade() error {
	log.Info("upgrading tilt")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade tilt")
	}
	log.Info("upgraded tilt")
	return nil
}
