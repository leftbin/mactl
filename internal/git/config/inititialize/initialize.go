package inititialize

import (
	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	gitconfig "github.com/leftbin/mactl/internal/git/config"
	sshcon "github.com/leftbin/mactl/internal/git/ssh/config"
	"github.com/leftbin/mactl/internal/git/ssh/key"
)

func Do(host, workspace, username, email string) error {
	log.Infof("initializing config for %s host and %s workspace", host, workspace)
	log.Info("ensuring ssh key")
	sshPvtPath, err := key.Cre(host, workspace)
	if err != nil {
		return errors.Wrap(err, "failed to ensure ssh key")
	}
	log.Info("ensured ssh key")
	if err := gitconfig.Set(host, workspace, username, email, sshPvtPath); err != nil {
		return errors.Wrap(err, "failed to set git config")
	}
	if err := key.Use(host, workspace); err != nil {
		return errors.Wrapf(err, "failed to set the ssh key as the %s key for %s", workspace, host)
	}
	pubKey, err := key.Get(host, workspace)
	if err != nil {
		return errors.Wrap(err, "failed to get pub key content")
	}
	if err := clipboard.WriteAll(pubKey); err != nil {
		log.Fatal("failed to copy pub key to clipboard")
	}
	log.Infof("ensuring %s host is in ssh config", host)
	if err := sshcon.AddHost(host); err != nil {
		return errors.Wrapf(err, "failed to ensure %s host is in ssh config", host)
	}
	log.Infof("ensured %s host is in ssh config", host)
	key.HandlePubKeyUpdOnScm(host)
	return nil
}
