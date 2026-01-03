package karabiner

import (
	"github.com/pkg/errors"
	karabinerconfig "github.com/leftbin/mactl/internal/app/keyboard/karabiner/config"
    "github.com/leftbin/mactl/internal/installer/brew"
    "github.com/leftbin/mactl/internal/installer/macapp"
)

const (
	BrewPkg    = "karabiner-elements"
	MacAppName = "Karabiner-Elements.app"
)

func Install() error {
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s using brew", BrewPkg)
	}
	return nil
}

func Upgrade() error {
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s using brew", BrewPkg)
	}
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open %s", MacAppName)
	}
	return nil
}

func Configure() error {
	if err := karabinerconfig.Setup(); err != nil {
		return errors.Wrapf(err, "failed to configure karabiner")
	}
	return nil
}
