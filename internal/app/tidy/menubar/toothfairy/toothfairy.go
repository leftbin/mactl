package toothfairy

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"github.com/leftbin/mactl/internal/installer/mas"
	log "github.com/sirupsen/logrus"
)

const (
	AppStoreId = "1191449274"
	MacAppName = "ToothFairy.app"
)

func Setup() error {
	log.Info("installing tooth fairy")
	if err := mas.Install(AppStoreId); err != nil {
		return errors.Wrap(err, "failed to install tooth fairy")
	}
	log.Info("installed tooth fairy")
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open tooth fairy")
	}
	return nil
}
