package kubectx

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "kubectx"
)

func Setup() error {
	log.Info("installing kubectx")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install kubectx")
	}
	log.Info("installed kubectx")
	return nil
}

func Upgrade() error {
	log.Info("upgrading kubectx")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade kubectx")
	}
	log.Info("upgraded kubectx")
	return nil
}
