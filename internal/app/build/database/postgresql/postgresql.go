package postgresql

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg        = "postgresql"
	BrewPkgPgAdmin = "pgadmin4"
)

func Setup() error {
	log.Info("installing postgresql")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install postgresql")
	}
	log.Info("installed postgresql")
	log.Info("installing pgadmin")
	if err := brew.Install(BrewPkgPgAdmin); err != nil {
		return errors.Wrap(err, "failed to install pgadmin")
	}
	log.Info("installed pgadmin")
	return nil
}

func Upgrade() error {
	log.Info("upgrading postgresql")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade postgresql")
	}
	log.Info("upgraded postgresql")
	return nil
}
