package slack

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
)

const (
	AppName        = "slack"
	BrewPkg        = "slack"
	MacAppFileName = "Slack.app"
)

func Setup() error {
	log.Infof("installing %s", AppName)
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Infof("installed %s", AppName)
	return nil
}

func GetPath() string {
	return macapp.GetPath(MacAppFileName)
}
