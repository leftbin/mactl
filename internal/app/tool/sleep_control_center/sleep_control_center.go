package sleep_control_center

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"github.com/leftbin/mactl/internal/installer/mas"
)

const (
	AppStoreId = "946798523"
	MacAppName = "Sleep Control Center.app"
)

func Setup() error {
	log.Info("installing sleep-control-center")
	if err := mas.Install(AppStoreId); err != nil {
		return errors.Wrap(err, "failed to install sleep-control-center")
	}
	log.Info("installed sleep-control-center")
	log.Info("opening sleep-control-center")
	if err := open(); err != nil {
		return errors.Wrap(err, "failed to open sleep-control-center")
	}
	log.Info("opened sleep-control-center")
	return nil
}

func open() error {
	err := macapp.Open(MacAppName)
	if err != nil {
		return errors.Wrap(err, "failed to open sleep-control-center")
	}
	return nil
}
