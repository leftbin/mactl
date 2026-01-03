package snagit

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	log "github.com/sirupsen/logrus"
)

const (
	AppName        = "snagit"
	BrewPkg        = "snagit"
	MacAppFileName = "Snagit 2022.app"
)

func Setup() error {
	log.Info("installing snagit")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Info("installed snagit")
	return nil
}

func Upgrade() error {
	log.Info("upgrading snagit")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Info("upgraded snagit")
	return nil
}

func configure() error {
	log.Info("importing config")
	log.Warn("TODO: configuring snagit has not been implemented yet")
	log.Info("imported config")
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppFileName); err != nil {
		return errors.Wrapf(err, "failed to open %s app", MacAppFileName)
	}
	return nil
}

func GetPath() string {
	return macapp.GetPath(MacAppFileName)
}
