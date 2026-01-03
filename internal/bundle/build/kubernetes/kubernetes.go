package kubernetes

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/helm"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/istioctl"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/k9s"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/kind"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/kubectl"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/kubectx"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/kustomize"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/linkerd"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/stern"
	"github.com/leftbin/mactl/internal/app/build/kubernetes/tilt"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring kubectl")
	if err := kubectl.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure kubectl")
	}
	log.Info("ensured kubectl")
	log.Info("ensuring kubectx")
	if err := kubectx.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure kubectx")
	}
	log.Info("ensured kubectx")
	log.Info("ensuring helm")
	if err := helm.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure helm")
	}
	log.Info("ensured helm")
	log.Info("ensuring kustomize")
	if err := kustomize.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure kustomize")
	}
	log.Info("ensured kustomize")
	log.Info("ensuring stern")
	if err := stern.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure stern")
	}
	log.Info("ensured stern")
	log.Info("ensuring kind")
	if err := kind.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure kind")
	}
	log.Info("ensured kind")
	log.Info("ensuring k9s")
	if err := k9s.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure k9s")
	}
	log.Info("ensured k9s")
	log.Info("ensuring tilt")
	if err := tilt.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure tilt")
	}
	log.Info("ensured tilt")
	log.Info("ensuring linkerd")
	if err := linkerd.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure linkerd")
	}
	log.Info("ensured linkerd")
	log.Info("ensuring istioctl")
	if err := istioctl.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure istioctl")
	}
	log.Info("ensured istioctl")
	return nil
}

func Upgrade() error {
	log.Info("upgrading kubectl")
	if err := kubectl.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade kubectl")
	}
	log.Info("upgraded kubectl")

	log.Info("upgrading kubectx")
	if err := kubectx.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade kubectx")
	}
	log.Info("upgraded kubectx")

	log.Info("upgrading helm")
	if err := helm.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade helm")
	}
	log.Info("upgraded helm")

	log.Info("upgrading stern")
	if err := stern.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade stern")
	}
	log.Info("upgraded stern")

	log.Info("upgrading tilt")
	if err := tilt.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade tilt")
	}
	log.Info("upgraded tilt")

	log.Info("upgrading linkerd")
	if err := linkerd.Upgrade(); err != nil {
		return errors.Wrapf(err, "failed to upgrade linkerd")
	}
	log.Info("upgraded linkerd")
	return nil
}
