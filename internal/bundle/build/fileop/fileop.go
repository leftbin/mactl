package fileop

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	log "github.com/sirupsen/logrus"
)

var Apps = []macapp.App{
	{
		Name:    "tree",
		BrewPkg: "tree",
	}, {
		Name:    "gnu-sed",
		BrewPkg: "gnu-sed",
	}, {
		Name:    "jq",
		BrewPkg: "jq",
	}, {
		Name:    "yq",
		BrewPkg: "yq",
	}, {
		Name:    "bat", //https://github.com/sharkdp/bat
		BrewPkg: "bat",
	}, {
		Name:    "ripgrep", //https://github.com/BurntSushi/ripgrep
		BrewPkg: "ripgrep",
	},
}

func Setup() error {
	log.Info("installing file management apps")
	for _, app := range Apps {
		log.Infof("ensuring %s", app.Name)
		if err := brew.Install(app.BrewPkg); err != nil {
			return errors.Wrapf(err, "failed to install %s pkg using brew", app.BrewPkg)
		}
		log.Infof("ensured %s", app.Name)
	}
	log.Info("installed file management apps")
	return nil
}

func Upgrade() error {
	log.Info("upgrading file management apps")
	for _, app := range Apps {
		log.Infof("upgrading %s", app.Name)
		if err := brew.Upgrade(app.BrewPkg); err != nil {
			return errors.Wrapf(err, "failed to upgrade %s pkg using brew", app.BrewPkg)
		}
		log.Infof("upgraded %s", app.Name)
	}
	log.Info("upgraded file management apps")
	return nil
}
