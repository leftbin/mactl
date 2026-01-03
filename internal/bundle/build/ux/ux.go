package ux

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/ux/figma"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring design tools")
	if err := figma.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure %s", figma.AppName)
	}
	log.Info("ensured design tools")
	return nil
}
