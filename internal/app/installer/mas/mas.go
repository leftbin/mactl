package mas

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "mas"
)

func Setup() error {
	log.Infof("installing mas")
	if err := installMas(); err != nil {
		return errors.Wrap(err, "failed to install mas")
	}
	log.Infof("installed mas")
	return nil
}

//installMas installs mas cli tool to install apps from app store
func installMas() error {
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s package using brew", BrewPkg)
	}
	return nil
}

func UpgradeMas() error {
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s package using brew", BrewPkg)
	}
	return nil
}
