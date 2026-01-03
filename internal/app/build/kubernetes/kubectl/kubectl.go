package kubectl

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/kubectl/plugin"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkgKubectl = "kubectl"
	BrewPkgKrew    = "krew"
)

func Setup() error {
	log.Info("installing kubectl")
	if err := brew.Install(BrewPkgKubectl); err != nil {
		return errors.Wrap(err, "failed to install kubectl")
	}
	log.Info("installed kubectl")
	log.Info("installing krew")
	if err := brew.Install(BrewPkgKrew); err != nil {
		return errors.Wrap(err, "failed to install krew")
	}
	log.Info("installed krew")
	log.Info("ensuring kubectl plugins")
	if err := plugin.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure kubectl plugins")
	}
	log.Info("ensured kubectl plugins")
	return nil
}

func Upgrade() error {
	log.Info("upgrading kubectl")
	if err := brew.Upgrade(BrewPkgKubectl); err != nil {
		return errors.Wrap(err, "failed to upgrade kubectl")
	}
	log.Info("upgraded kubectl")

	log.Info("upgrading krew")
	if err := brew.Upgrade(BrewPkgKrew); err != nil {
		return errors.Wrap(err, "failed to upgrade krew")
	}
	log.Info("upgraded krew")

	// upgrade kubectl plugins and handle error gracefully as they are run as go routines
	return nil
}
