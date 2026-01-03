package pip

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os/exec"
)

func Install(pkg string) error {
	if err := shell.RunCmd(exec.Command("pip3", "install", pkg)); err != nil {
		return errors.Wrapf(err, "failed to install %s using pip3", pkg)
	}
	return nil
}

func UnInstall(pkg string) error {
	if err := shell.RunCmd(exec.Command("pip3", "uninstall", "-y", pkg)); err != nil {
		return errors.Wrapf(err, "failed to uninstall %s using pip3", pkg)
	}
	return nil
}

func UpgradePip() error {
	if err := shell.RunCmd(exec.Command("pip3", "install", "--upgrade", "pip")); err != nil {
		return errors.Wrap(err, "failed to upgrade pip using pip3")
	}
	return nil
}
