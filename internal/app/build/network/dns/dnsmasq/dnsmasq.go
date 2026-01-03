package dnsmasq

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/build/network/dns/dnsmasq/config"
	"github.com/leftbin/mactl/internal/installer/brew"
)

const (
	BrewPkg = "dnsmasq"
)

func Setup() error {
	log.Infof("installing %s", BrewPkg)
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Infof("installed %s", BrewPkg)
	log.Infof("configuring dnsmsq")
	if err := configure(); err != nil {
		return errors.Wrapf(err, "failed to configure")
	}
	log.Infof("configured dnsmasq")
	return nil
}

func configure() error {
	if err := config.Setup(); err != nil {
		return errors.Wrapf(err, "failed to configure dnsmasq")
	}
	return nil
}

func Upgrade() error {
	log.Infof("upgrading %s", BrewPkg)
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Infof("upgraded %s", BrewPkg)
	return nil
}
