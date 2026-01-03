package kind

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "kind"
)

func Setup() error {
	log.Info("installing kind")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install stern")
	}
	log.Info("installed kind")
	return nil
}

func Upgrade() error {
	log.Info("upgrading kind")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade stern")
	}
	log.Info("upgraded kind")
	return nil
}
