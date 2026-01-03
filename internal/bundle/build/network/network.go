package network

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/network/dns/dnsmasq"
	"github.com/leftbin/mactl/internal/app/build/network/stunnel"
	"github.com/leftbin/mactl/internal/app/build/network/wget"
	"github.com/leftbin/mactl/internal/app/network/step"
	"github.com/leftbin/mactl/internal/app/network/telnet"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring telnet")
	if err := telnet.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure telnet")
	}
	log.Info("ensured telnet")
	log.Info("ensuring step")
	if err := step.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure step")
	}
	log.Info("ensured step")
	log.Info("ensuring stunnel")
	if err := stunnel.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure stunnel")
	}
	log.Info("ensured stunnel")
	log.Info("ensuring wget")
	if err := wget.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure wget")
	}
	log.Info("ensured wget")
	log.Info("ensuring dnsmasq")
	if err := dnsmasq.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure dnsmasq")
	}
	log.Info("ensured dnsmasq")
	return nil
}

func Upgrade() error {
	log.Info("upgrading telnet")
	if err := telnet.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade telnet")
	}
	log.Info("upgraded telnet")

	log.Info("upgrading step")
	if err := step.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade step")
	}
	log.Info("upgraded step")
	log.Info("upgrading dnsmasq")
	if err := dnsmasq.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade dnsmasq")
	}
	log.Info("upgraded dnsmasq")
	return nil
}
