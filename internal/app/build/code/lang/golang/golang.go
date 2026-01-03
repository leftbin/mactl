package golang

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/code/lang/golang/goland"
	"github.com/leftbin/mactl/internal/app/build/code/lang/golang/protoc/plugin"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "go"
)

func Setup() error {
	log.Infof("ensuring golang compiler")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to ensure %s brew pkg", BrewPkg)
	}
	log.Infof("ensured golang compiler")
	log.Infof("ensuring goland ide")
	if err := goland.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure ide")
	}
	log.Infof("ensured goland ide")
	log.Infof("ensuring protoc plugins")
	if err := plugin.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure protoc plugins")
	}
	log.Infof("ensured protoc plugins")
	return nil
}

func Upgrade() error {
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade %s brew pkg", BrewPkg)
	}
	return nil
}
