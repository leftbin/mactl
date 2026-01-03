package code

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/code/delta"
	"github.com/leftbin/mactl/internal/bundle/build/code/editor"
	"github.com/leftbin/mactl/internal/bundle/build/code/lang"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring git tools")
	if err := delta.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure git-delta")
	}
	log.Info("ensured git tools")
	log.Info("ensuring code bundle")
	if err := lang.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure code bundle")
	}
	log.Info("ensured code bundle")
	log.Info("ensuring editor bundle")
	if err := editor.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure editor bundle")
	}
	log.Info("ensured editor bundle")
	return nil
}
