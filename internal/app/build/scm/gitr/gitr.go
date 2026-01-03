package gitr

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const (
	BrewPkg        = "leftbin/homebrew-tap/gitr"
	ConfigFileName = ".gitr.yaml"
)

type ConfigTemplateInput struct {
	UserHomeDir string
}

func Setup() error {
	log.Infof("installing gitr")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Infof("installed gitr")
	log.Infof("configuring gitr")
	if err := configure(); err != nil {
		return errors.Wrap(err, "failed to configure gitr")
	}
	log.Infof("configured gitr")
	return nil
}

func Upgrade() error {
	log.Infof("upgrading gitr")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Infof("upgraded gitr")

	log.Infof("configuring gitr")
	if err := configure(); err != nil {
		return errors.Wrap(err, "failed to configure gitr")
	}
	log.Infof("configured gitr")
	return nil
}

func configure() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	renderedBytes, err := file.RenderTmplt(&ConfigTemplateInput{UserHomeDir: homeDir}, ConfigTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to render gitr config")
	}
	configPath := filepath.Join(homeDir, ConfigFileName)
	if err := os.WriteFile(configPath, renderedBytes, 0644); err != nil {
		return errors.Wrapf(err, "failed to write %s file", configPath)
	}
	return nil
}
