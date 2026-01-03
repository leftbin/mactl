package iac

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/iac/packer"
	"github.com/leftbin/mactl/internal/app/build/iac/pulumi"
	"github.com/leftbin/mactl/internal/app/build/iac/terraform"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring packer")
	if err := packer.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure packer")
	}
	log.Info("ensured packer")
	log.Info("ensuring terraform")
	if err := terraform.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure terraform")
	}
	log.Info("ensured terraform")
	log.Info("ensuring pulumi")
	if err := pulumi.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure pulumi")
	}
	log.Info("ensured pulumi")
	return nil
}

func Upgrade() error {
	log.Info("upgrading packer")
	if err := packer.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade packer")
	}
	log.Info("upgraded packer")

	log.Info("upgrading terraform")
	if err := terraform.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade terraform")
	}
	log.Info("upgraded terraform")

	log.Info("upgrading pulumi")
	if err := pulumi.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade pulumi")
	}
	log.Info("upgraded pulumi")
	return nil
}
