package build

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/bundle/build/api"
	"github.com/leftbin/mactl/internal/bundle/build/auth"
	"github.com/leftbin/mactl/internal/bundle/build/cloud"
	"github.com/leftbin/mactl/internal/bundle/build/code"
	"github.com/leftbin/mactl/internal/bundle/build/code/lang"
	"github.com/leftbin/mactl/internal/bundle/build/database"
	"github.com/leftbin/mactl/internal/bundle/build/fileop"
	"github.com/leftbin/mactl/internal/bundle/build/iac"
	"github.com/leftbin/mactl/internal/bundle/build/kubernetes"
	"github.com/leftbin/mactl/internal/bundle/build/network"
	"github.com/leftbin/mactl/internal/bundle/build/scm"
	"github.com/leftbin/mactl/internal/bundle/build/ux"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring file-operations bundle")
	if err := fileop.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure file bundle")
	}
	log.Info("ensured file-operations bundle")
	log.Info("ensuring ux bundle")
	if err := ux.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure ux bundle")
	}
	log.Info("ensured ux bundle")
	log.Info("ensuring auth bundle")
	if err := auth.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure auth bundle")
	}
	log.Info("ensured auth bundle")
	log.Info("ensuring api-client bundle")
	if err := api.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure api-client bundle")
	}
	log.Info("ensured api-client bundle")
	log.Info("ensuring cloud bundle")
	if err := cloud.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure cloud bundle")
	}
	log.Info("ensured cloud bundle")
	log.Info("ensuring iac bundle")
	if err := iac.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure iac bundle")
	}
	log.Info("ensured iac bundle")
	log.Info("ensuring kubernetes bundle")
	if err := kubernetes.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure k8s bundle")
	}
	log.Info("ensured kubernetes bundle")
	log.Info("ensuring code bundle")
	if err := code.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure code bundle")
	}
	log.Info("ensured code bundle")
	log.Info("ensuring database bundle")
	if err := database.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure database bundle")
	}
	log.Info("ensured database bundle")
	log.Info("ensuring network bundle")
	if err := network.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure network bundle")
	}
	log.Info("ensured network bundle")
	return nil
}

func Upgrade() error {
	log.Info("upgrading scm")
	if err := scm.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade scm")
	}
	log.Info("upgraded scm")

	log.Info("upgrading file-operations bundle")
	if err := fileop.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade file bundle")
	}
	log.Info("upgraded file-operations bundle")

	log.Info("upgrading auth bundle")
	if err := auth.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade auth bundle")
	}
	log.Info("upgraded auth bundle")

	log.Info("upgrading api-client bundle")
	if err := api.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade api-client bundle")
	}
	log.Info("upgraded api-client bundle")

	log.Info("upgrading cloud bundle")
	if err := cloud.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade cloud bundle")
	}
	log.Info("upgraded cloud bundle")

	log.Info("upgrading iac bundle")
	if err := iac.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade iac bundle")
	}
	log.Info("upgraded iac bundle")

	log.Info("upgrading k8s bundle")
	if err := kubernetes.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade k8s bundle")
	}
	log.Info("upgraded k8s bundle")

	log.Info("upgrading network bundle")
	if err := network.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade network bundle")
	}
	log.Info("upgraded network bundle")

	log.Info("upgrading code bundle")
	if err := lang.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade code bundle")
	}
	log.Info("upgraded code bundle")

	log.Info("upgrading editor bundle")
	if err := lang.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade editor bundle")
	}
	log.Info("upgraded editor bundle")
	return nil
}
