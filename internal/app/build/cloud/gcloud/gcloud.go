package gcloud

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
)

const (
	BrewPkg                 = "google-cloud-sdk"
	GkeGcloudAuthPluginName = "gke-gcloud-auth-plugin"
)

func Setup() error {
	log.Info("installing gcloud")
	err := brew.Install(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to install gcloud")
	}
	log.Info("installed gcloud")

	log.Infof("installing %s", GkeGcloudAuthPluginName)
	if err := installGkeGcloudAuthPlugin(); err != nil {
		return errors.Wrapf(err, "failed to install %s", GkeGcloudAuthPluginName)
	}
	log.Infof("installed %s", GkeGcloudAuthPluginName)
	return nil
}

func Upgrade() error {
	log.Info("upgrading gcloud")
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade gcloud")
	}
	log.Info("upgraded gcloud")
	return nil
}

// installGkeGcloudAuthPlugin installs kubectl auth plugin to work with gke cluster
// https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
func installGkeGcloudAuthPlugin() error {
	if err := shell.RunCmd(exec.Command(getGcloudBinaryLoc(), "components", "install", "--quiet", GkeGcloudAuthPluginName)); err != nil {
		return errors.Wrapf(err, "failed to install ")
	}
	return nil
}

func getGcloudBinaryLoc() string {
	if runtime.GOARCH == "arm64" {
		return "/opt/homebrew/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/bin/gcloud"
	}
	return "/usr/local/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/bin/gcloud"
}
