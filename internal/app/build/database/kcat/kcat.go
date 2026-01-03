package kcat

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "kcat"
)

func Setup() error {
	log.Info("installing kcat")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install kcat")
	}
	log.Info("installed kcat")
	return nil
}

func Upgrade() error {
	log.Info("upgrading kcat")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade kcat")
	}
	log.Info("upgraded kcat")
	return nil
}
