package postman

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
)

const (
	AppName        = "postman"
	BrewPkg        = "postman"
	MacAppFileName = "Postman.app"
)

func Setup() error {
	log.Infof("installing %s", AppName)
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s", AppName)
	}
	log.Infof("installed %s", AppName)
	return nil
}

func Upgrade() error {
	log.Infof("upgrading %s", AppName)
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s", AppName)
	}
	log.Infof("upgraded %s", AppName)
	return nil
}

func GetPath() string {
	return macapp.GetPath(MacAppFileName)
}
