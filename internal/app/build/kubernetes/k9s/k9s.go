package k9s

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "k9s"
)

func Setup() error {
	log.Info("installing k9s")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install stern")
	}
	log.Info("installed k9s")
	return nil
}

func Upgrade() error {
	log.Info("upgrading k9s")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade stern")
	}
	log.Info("upgraded k9s")
	return nil
}
