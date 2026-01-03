package hotkey

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/keyboard/karabiner"
)

func Install() error {
	log.Info("installing karabiner")
	if err := karabiner.Install(); err != nil {
		return errors.Wrap(err, "failed to install karabiner")
	}
	log.Info("installed karabiner")
	return nil
}

func Upgrade() error {
	log.Info("upgrading karabiner")
	if err := karabiner.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade karabiner")
	}
	log.Info("upgraded karabiner")
	return nil
}

func Configure() error {
	log.Info("configuring karabiner")
	if err := karabiner.Configure(); err != nil {
		return errors.Wrap(err, "failed to configure karabiner")
	}
	log.Info("configured karabiner")
	return nil
}
