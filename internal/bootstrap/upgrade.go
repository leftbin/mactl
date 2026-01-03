package bootstrap

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/browser/chrome"
	"github.com/leftbin/mactl/internal/app/build/terminal/iterm"
	"github.com/leftbin/mactl/internal/app/installer/mas"
	"github.com/leftbin/mactl/internal/app/keyboard/rectangle"
	"github.com/leftbin/mactl/internal/bundle/build"
	"github.com/leftbin/mactl/internal/bundle/hotkey"
	"github.com/leftbin/mactl/internal/bundle/tool"
	"github.com/leftbin/mactl/internal/optimize/dock"
	log "github.com/sirupsen/logrus"
)

const (
	UpgradeOptionalCommApps = true
)

func Upgrade() error {
	log.Info("bootstrap upgrade started")

	done := make(chan bool)
	fatalErrors := make(chan error)

	if err := upgrade(fatalErrors); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to ensure upgrade all bootstrap components")
	}

	if err := open(fatalErrors); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to open all apps installed as part of bootstrap upgrade")
	}

	close(done)
	select {
	case <-done:
		break
	case err := <-fatalErrors:
		close(fatalErrors)
		log.Warnf("Error encountered : %v", err.Error())
		return err
	}

	log.Info("bootstrap upgrade completed")
	return nil
}

func upgrade(fatalErrors chan error) error {
	log.Info("optimizing dock")
	if err := dock.Optimize(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to optimize dock")
	}
	log.Info("optimized dock")

	log.Info("upgrading installers")
	if err := mas.UpgradeMas(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to upgrade installers")
	}
	log.Info("upgraded installers")

	log.Info("upgrading hotkey")
	if err := hotkey.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to upgrade hotkey")
	}
	log.Info("upgraded hotkey")

	log.Info("upgrading window-management")
	if err := rectangle.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to upgrade window management using rectangle app")
	}
	log.Info("upgraded window-management")

	log.Info("upgrading browser")
	if err := chrome.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to upgrade browser")
	}
	log.Info("upgraded browser")

	log.Info("upgrading terminal")
	if err := iterm.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to upgrade scm")
	}
	log.Info("upgraded terminal")

	log.Info("upgrading tool bundle")
	if err := tool.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to upgrade tool bundle")
	}
	log.Info("upgraded tool bundle")

	log.Info("upgrading build bundle")
	if err := build.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to upgrade ides")
	}
	log.Info("upgraded build bundle")
	return nil
}
