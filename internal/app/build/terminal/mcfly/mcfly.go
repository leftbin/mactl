// Package mcfly installs https://github.com/cantino/mcfly
package mcfly

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
)

const (
	BrewPkg = "cantino/mcfly/mcfly"
)

func Setup() error {
	if runtime.GOARCH == "arm64" {
		log.Info("installing rosetta required for mcfly software to run on apple silicon")
		if err := shell.RunCmd(exec.Command("softwareupdate", "--install-rosetta")); err != nil {
			return errors.Wrapf(err, "failed to install rosetta required by mcfly")
		}
	}
	log.Info("installing mcfly")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install mcfly")
	}
	log.Info("installed mcfly")
	return nil
}

func Upgrade() error {
	log.Info("upgrading mcfly")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade mcfly")
	}
	log.Info("upgraded mcfly")
	return nil
}
