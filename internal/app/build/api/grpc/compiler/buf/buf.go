package buf

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	Name    = "buf"
	BrewPkg = "bufbuild/buf/buf"
)

func Setup() error {
	log.Infof("installing %s", Name)
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s", Name)
	}
	log.Infof("installed %s", Name)
	return nil
}

func Upgrade() error {
	log.Infof("upgrading %s", Name)
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s", Name)
	}
	log.Infof("upgraded %s", Name)
	return nil
}
