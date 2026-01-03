package scm

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/scm/github"
	"github.com/leftbin/mactl/internal/app/build/scm/gitr"
	gitcon "github.com/leftbin/mactl/internal/git/config"
	log "github.com/sirupsen/logrus"
)

type Provider string

const (
	ProviderGitlab                          Provider = "gitlab.com"
	ProviderGitHub                          Provider = "github.com"
	ProviderGitlabSshKeyConfigurePageUrl             = "https://gitlab.com/-/profile/keys"
	ProviderGitGithubSshKeyConfigurePageUrl          = "https://github.com/settings/ssh/new"
)

// Setup tools and config needed for
// cloning repos to correct dirs
// cloning public & private repos using correct ssh keys
// making commits with correct author
// pushing changes using correct ssh keys
func Setup() error {
	log.Infof("ensuring global ignore")
	if err := gitcon.CreGlobalIgnore(); err != nil {
		return errors.Wrap(err, "failed to ensure global ignore file")
	}
	log.Infof("ensured global ignore")
	log.Infof("ensuring gitr")
	if err := gitr.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure gitr")
	}
	log.Infof("ensured gitr")
	log.Infof("ensuring github-desktop")
	if err := github.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure github desktop")
	}
	log.Infof("ensured github-desktop")
	return nil
}

func Upgrade() error {
	log.Infof("upgrading gitr")
	if err := gitr.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade gitr")
	}
	log.Infof("upgraded gitr")

	log.Infof("upgrading github-desktop")
	if err := github.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade github desktop")
	}
	log.Infof("upgraded github-desktop")
	return nil
}
