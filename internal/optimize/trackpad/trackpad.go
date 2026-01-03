package trackpad

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/cli/cache"
	"github.com/leftbin/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

// Optimize trackpad will
// increase trackpad speed to maximum
// TODO: enable tap to click
// WARNING: this functionality is in a broken state
func Optimize() error {
	log.Infof("optimzing trackpad")
	log.Infof("increasing tracking speed")
	if err := setTrackingSpeed(); err != nil {
		return errors.Wrap(err, "failed to set tracking speed")
	}
	log.Infof("increased tracking speed")
	log.Infof("enabling tap to click")
	if err := enableTapToClick(); err != nil {
		return errors.Wrap(err, "failed to enable tap to click")
	}
	log.Infof("enabled tap to click")
	log.Infof("optimzed trackpad")
	return nil
}

// TODO: implement
func enableTapToClick() error {
	return nil
}

func setTrackingSpeed() error {
	cacheLoc, err := cache.GetLoc()
	if err != nil {
		return errors.Wrap(err, "failed to get cache loc")
	}
	scriptPath := filepath.Join(cacheLoc, "optimize", "trackpad", "inc-tracking-speed.scpt")
	if err := os.MkdirAll(filepath.Dir(scriptPath), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(scriptPath))
	}
	if err := os.WriteFile(scriptPath, []byte(AppleScriptToIncTrackingSpeed), os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to create script")
	}
	if err := shell.RunCmd(exec.Command("osascript", scriptPath)); err != nil {
		return errors.Wrap(err, "failed to run apple-script")
	}
	return nil
}
