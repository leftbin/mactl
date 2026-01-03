package javascript

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/build/code/lang/javascript/nodejs"
	"github.com/leftbin/mactl/internal/app/build/code/lang/javascript/webstorm"
)

func Setup() error {
	log.Infof("ensuring nodejs")
	if err := nodejs.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure ide")
	}
	log.Infof("ensured nodejs")
	log.Infof("ensuring ide")
	if err := webstorm.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure ide")
	}
	log.Infof("ensured ide")
	return nil
}
