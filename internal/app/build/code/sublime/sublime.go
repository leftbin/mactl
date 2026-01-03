package vscode

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
)

const (
	AppName        = "vs-code"
	BrewPkg        = "visual-studio-code"
	MacAppFileName = "Visual Studio Code.app"
)

func Setup() error {
	log.Infof("installing %s", BrewPkg)
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Infof("installed %s", BrewPkg)
	return nil
}

func Upgrade() error {
	log.Infof("upgrading %s", BrewPkg)
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Infof("upgraded %s", BrewPkg)
	return nil
}

func GetPath() string {
	return macapp.GetPath(MacAppFileName)
}
