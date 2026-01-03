package python

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/build/code/lang/python/pycharm"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "python"
)

func Setup() error {
	log.Infof("ensuring python")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to ensure python")
	}
	log.Infof("ensured python")
	log.Infof("ensuring ide")
	if err := pycharm.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure ide")
	}
	log.Infof("ensured ide")
	return nil
}

func Upgrade() error {
	log.Infof("upgrading python")
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade python")
	}
	log.Infof("upgraded python")
	return nil
}
