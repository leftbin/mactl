package stern

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "stern"
)

func Setup() error {
	log.Info("installing stern")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install stern")
	}
	log.Info("installed stern")
	return nil
}

func Upgrade() error {
	log.Info("upgrading stern")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade stern")
	}
	log.Info("upgraded stern")
	return nil
}
