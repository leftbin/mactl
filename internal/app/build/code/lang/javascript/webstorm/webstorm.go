package webstorm

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/bundle/tool"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"os"
	"path/filepath"
)

const (
	AppName        = "webstorm"
	BrewPkg        = "webstorm"
	MacAppFileName = "WebStorm.app"
	LaunchAlias    = "www"
	LaunchScript   = `
#!/bin/sh
IDEA="/Applications/WebStorm.app"
open -a "$IDEA" .
`
)

func Setup() error {
	log.Infof("installing %s", AppName)
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Infof("installed %s", AppName)
	log.Infof("ensuring launch alias")
	if err := setupLaunchAlias(); err != nil {
		return errors.Wrapf(err, "failed to ensure launch alias")
	}
	log.Infof("ensured launch alias")
	return nil
}

func GetPath() string {
	return macapp.GetPath(MacAppFileName)
}

func setupLaunchAlias() error {
	binDir, err := tool.GetBinDir()
	if err != nil {
		return errors.Wrapf(err, "failed to get bin dir")
	}
	if err := os.WriteFile(filepath.Join(binDir, LaunchAlias), []byte(LaunchScript), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", filepath.Join(binDir, LaunchScript))
	}
	return nil
}
