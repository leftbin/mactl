package istioctl

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "istioctl"
)

func Setup() error {
	log.Info("installing istioctl")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install helm")
	}
	log.Info("installed istioctl")
	return nil
}

func Upgrade() error {
	log.Info("upgrading istioctl")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade helm")
	}
	log.Info("upgraded istioctl")
	return nil
}
