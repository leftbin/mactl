package swift

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/code/lang/swift/protoc/plugin"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "swift"
)

func Setup() error {
	log.Infof("ensuring swift compiler")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to ensure %s brew pkg", BrewPkg)
	}
	log.Infof("ensured swift compiler")
	log.Infof("ensuring swift protoc plugins")
	if err := plugin.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure swift protoc plugins")
	}
	log.Infof("ensured swift protoc plugins")
	return nil
}

func Upgrade() error {
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade %s brew pkg", BrewPkg)
	}
	return nil
}
