package flycut

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"github.com/leftbin/mactl/internal/lib/plist"
)

const (
	BrewPkg    = "flycut"
	MacAppName = "Flycut.app"
	AppGroupId = "com.generalarcade.flycut"
	PlistB64   = "YnBsaXN0MDDZAQIDBAUGBwgJCgsMDQwMDBUWW3JlbWVtYmVyTnVtXxAmTlNTdGF0dXNJdGVtIFByZWZlcnJlZCBQb3NpdGlvbiBJdGVtLTBfEBpzdXBwcmVzc0FjY2Vzc2liaWxpdHlBbGVydF8QG1Nob3J0Y3V0UmVjb3JkZXIgbWFpbkhvdGtleV8QD3Bhc3RlTW92ZXNUb1RvcF1sb2FkT25TdGFydHVwXxAQcmVtb3ZlRHVwbGljYXRlc18QHnByZXZpb3VzU3luY0NsaXBwaW5nc1ZpYUlDbG91ZFVzdG9yZSNAWMAAAAAAACJFrbgACdIODxARV2tleUNvZGVdbW9kaWZpZXJGbGFncxAyEgAEAAAJCQkI1xcYGRobHB0eHyAhICIjVmpjTGlzdFtyZW1lbWJlck51bVpkaXNwbGF5TGVuV3ZlcnNpb25fEBRmYXZvcml0ZXNSZW1lbWJlck51bV1mYXZvcml0ZXNMaXN0WmRpc3BsYXlOdW2gEGMQKFMwLjegEAoACAAbACcAUABtAIsAnQCrAL4A3wDlAO4A8wD0APkBAQEPAREBFgEXARgBGQEaASkBMAE8AUcBTwFmAXQBfwGAAYIBhAGIAYkAAAAAAAACAQAAAAAAAAAkAAAAAAAAAAAAAAAAAAABiw=="
)

func Setup() error {
	log.Info("installing flycut")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Info("installed flycut")
	log.Info("configuring flycut")
	if err := configure(); err != nil {
		return errors.Wrap(err, "failed to configure flycut")
	}
	log.Info("configured flycut")

	return nil
}

func Upgrade() error {
	log.Info("upgrading flycut")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Info("upgraded flycut")

	return nil
}

func configure() error {
	log.Info("importing config")
	if err := plist.Import(AppGroupId, PlistB64); err != nil {
		return errors.Wrap(err, "failed to import config")
	}
	log.Info("imported config")
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open %s app", MacAppName)
	}
	return nil
}
