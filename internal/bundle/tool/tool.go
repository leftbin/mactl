package tool

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/tool/flux"
	"github.com/leftbin/mactl/internal/app/tool/flycut"
	"github.com/leftbin/mactl/internal/app/tool/snagit"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func Setup() error {
	log.Info("ensuring ${HOME}/bin directory")
	if err := EnsureBinDir(); err != nil {
		return errors.Wrap(err, "failed to ensure ${HOME}/bin directory")
	}
	log.Info("ensured ${HOME}/bin directory")
	log.Info("ensuring flycut")
	if err := flycut.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure flycut")
	}
	log.Info("ensured flycut")
	log.Info("ensuring flux")
	if err := flux.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure flux")
	}
	log.Info("ensured flux")
	log.Info("ensuring snagit")
	if err := snagit.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure snagit")
	}
	log.Info("ensured snagit")
	return nil
}

func EnsureBinDir() error {
	binDir, err := GetBinDir()
	if err != nil {
		return errors.Wrapf(err, "failed to ensure bin dir")
	}
	if file.IsDirExists(binDir) {
		return nil
	}
	if err := os.Mkdir(binDir, os.FileMode(0755)); err != nil {
		return errors.Wrapf(err, "failed to create %s dir", binDir)
	}
	return nil
}

func GetBinDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home dir")
	}
	return filepath.Join(dir, "bin"), nil
}

func Upgrade() error {
	log.Info("upgrading flycut")
	if err := flycut.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade flycut")
	}
	log.Info("upgraded flycut")

	log.Info("upgrading flux")
	if err := flux.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade flux")
	}
	log.Info("upgraded flux")

	log.Info("upgrading snagit")
	if err := snagit.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade snagit")
	}
	log.Info("upgraded snagit")

	return nil
}
