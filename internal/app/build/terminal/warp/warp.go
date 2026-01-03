// Package warp installs https://www.warp.dev/
package warp

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "warp"
)

func Setup() error {
	log.Info("installing warp")
	if err := brew.InstallCask(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install warp")
	}
	log.Info("installed warp")
	return nil
}

func Upgrade() error {
	log.Info("upgrading warp")
	if err := brew.UpgradeCask(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade warp")
	}
	log.Info("upgraded warp")
	return nil
}
