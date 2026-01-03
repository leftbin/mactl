package stunnel

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "stunnel"
)

func Setup() error {
	log.Info("installing stunnel")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install stunnel")
	}
	log.Info("installed stunnel")
	return nil
}

func Upgrade() error {
	log.Info("upgrading stunnel")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade stunnel")
	}
	log.Info("upgraded stunnel")
	return nil
}
