package sql

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/code/lang/sql/datagrip"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Infof("ensuring ide")
	if err := datagrip.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure ide")
	}
	log.Infof("ensured ide")
	return nil
}
