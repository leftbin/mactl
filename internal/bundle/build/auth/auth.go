package auth

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	log "github.com/sirupsen/logrus"
)

var Apps = []macapp.App{
	{
		Name:    "jwt-cli",
		BrewPkg: "mike-engel/jwt-cli/jwt-cli",
	},
}

func Setup() error {
	log.Info("installing auth apps")
	for _, app := range Apps {
		log.Infof("ensuring %s", app.Name)
		if err := brew.Install(app.BrewPkg); err != nil {
			return errors.Wrapf(err, "failed to install %s pkg using brew", app.BrewPkg)
		}
		log.Infof("ensured %s", app.Name)
	}
	log.Info("installed auth apps")
	return nil
}

func Upgrade() error {
	log.Info("upgrading auth apps")
	for _, app := range Apps {
		log.Infof("upgrading %s", app.Name)
		if err := brew.Upgrade(app.BrewPkg); err != nil {
			return errors.Wrapf(err, "failed to upgrade %s pkg using brew", app.BrewPkg)
		}
		log.Infof("upgraded %s", app.Name)
	}
	log.Info("upgraded auth apps")
	return nil
}
