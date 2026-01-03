package intellij

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/bundle/tool"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const (
	AppName        = "intellij"
	BrewPkg        = "intellij-idea"
	MacAppFileName = "IntelliJ IDEA.app"
	LaunchAlias    = "jjj"
	LaunchScript   = `
#!/bin/sh

IDEA="/Applications/IntelliJ IDEA.app"

# were we given a file?
if [ -f "$1" ]; then
  open -a "$IDEA" "$1"
  exit 0
fi

if [ -f build.gradle ]; then
  open -a "$IDEA" "build.gradle"
  exit 0
fi

if [ -f pom.xml ]; then
  open -a "$IDEA" "pom.xml"
  exit 0
fi
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
