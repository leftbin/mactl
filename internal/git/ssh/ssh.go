package ssh

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/git/ssh/config"
	"os"
	"path/filepath"
)

func Init() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	sshDir := filepath.Join(homeDir, ".ssh")
	if err := os.MkdirAll(sshDir, 0744); err != nil {
		return errors.Wrapf(err, "failed to create %s dir", sshDir)
	}
	if err := config.Init(); err != nil {
		return errors.Wrapf(err, "failed to initilize config")
	}
	return nil
}
