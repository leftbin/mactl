package linkerd

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "linkerd"
)

func Setup() error {
	log.Info("installing linkerd")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install linkerd")
	}
	log.Info("installed linkerd")
	return nil
}

func Upgrade() error {
	log.Info("upgrading linkerd")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade linkerd")
	}
	log.Info("upgraded linkerd")
	return nil
}
