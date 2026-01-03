package cloud

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/cloud/awscli"
	"github.com/leftbin/mactl/internal/app/build/cloud/gcloud"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("installing cloud apps")
	log.Info("ensuring gcloud")
	if err := gcloud.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure gcloud")
	}
	log.Info("ensured gcloud")
	log.Info("ensuring awscli")
	if err := awscli.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure awscli")
	}
	log.Info("ensured awscli")
	log.Info("installed cloud apps")
	return nil
}

func Upgrade() error {
	log.Info("upgrading cloud apps")
	log.Info("upgrading gcloud")
	if err := gcloud.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade gcloud")
	}
	log.Info("upgraded gcloud")

	log.Info("upgrading awscli")
	if err := awscli.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade awscli")
	}
	log.Info("upgraded awscli")
	log.Info("upgraded cloud apps")
	return nil
}
