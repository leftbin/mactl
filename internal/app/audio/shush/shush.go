package shush

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"github.com/leftbin/mactl/internal/installer/mas"
)

const (
	AppStoreId = "496437906"
	MacAppName = "Shush.app"
)

func Setup() error {
	log.Info("installing shush")
	if err := mas.Install(AppStoreId); err != nil {
		return errors.Wrap(err, "failed to install shush")
	}
	log.Info("installed shush")
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open shush")
	}
	return nil
}
