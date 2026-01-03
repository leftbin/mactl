package krisp

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
)

const (
	BrewPkg    = "krisp"
	MacAppName = "Krisp.app"
)

func Setup() error {
	log.Info("installing krisp")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install krisp")
	}
	log.Info("installed krisp")
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open shush")
	}
	return nil
}
