package pulumi

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "pulumi"
)

func Setup() error {
	log.Info("installing pulumi")
	err := brew.Install(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to install pulumi")
	}
	log.Info("installed pulumi")
	return nil
}

func Upgrade() error {
	log.Info("upgrading pulumi")
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade pulumi")
	}
	log.Info("upgraded pulumi")
	return nil
}
