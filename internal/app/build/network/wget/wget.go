package wget

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "wget"
)

func Setup() error {
	log.Info("installing wget")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install wget")
	}
	log.Info("installed wget")
	return nil
}

func Upgrade() error {
	log.Info("upgrading wget")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade wget")
	}
	log.Info("upgraded wget")
	return nil
}
