package brew

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os/exec"
)

func Install(pkg string) error {
	if err := shell.RunCmd(exec.Command("brew", "install", "--force", pkg)); err != nil {
		return errors.Wrapf(err, "failed to install %s using brew", pkg)
	}
	return nil
}

func InstallCask(pkg string) error {
	if err := shell.RunCmd(exec.Command("brew", "install", "--cask", pkg)); err != nil {
		return errors.Wrapf(err, "failed to install %s cask using brew", pkg)
	}
	return nil
}

func Upgrade(pkg string) error {
	if err := shell.RunCmd(exec.Command("brew", "upgrade", pkg)); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s using brew", pkg)
	}
	return nil
}

func UpgradeCask(pkg string) error {
	if err := shell.RunCmd(exec.Command("brew", "upgrade", "--cask", pkg)); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s cask using brew", pkg)
	}
	return nil
}
