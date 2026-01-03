package database

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/database/kcat"
	"github.com/leftbin/mactl/internal/app/build/database/postgresql"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring postgresql")
	if err := postgresql.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure postgresql")
	}
	log.Info("ensured postgresql")
	log.Info("ensuring kcat")
	if err := kcat.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure kcat")
	}
	log.Info("ensured kcat")
	return nil
}
