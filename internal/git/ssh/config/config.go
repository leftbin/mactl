package config

import (
	"github.com/kevinburke/ssh_config"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/git/ssh/config/host"
	"os"
	"path/filepath"
	"strings"
)

const (
	HostConfigTemplate = `
Host {{.HostName}}
  #mactl-scm-host
  HostName {{.HostName}}
  User git
  IdentityFile {{.ScmSshDir}}/{{.HostName}}
`
)

type HostConfigInput struct {
	HostName  string
	ScmSshDir string
}

func Init() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	sshConfigFile := filepath.Join(homeDir, ".ssh", "config")
	if file.IsFileExists(sshConfigFile) {
		return nil
	}
	defaultConfig, err := getDefaultConfig()
	if err != nil {
		return errors.Wrap(err, "failed to get default config")
	}
	if err := os.WriteFile(sshConfigFile, []byte(defaultConfig), 0744); err != nil {
		return errors.Wrapf(err, "failed to initialize %s file", sshConfigFile)
	}
	return nil
}

func AddHost(hostname string) error {
	hostExists, err := host.IsExists(hostname)
	if err != nil {
		return errors.Wrap(err, "failed to check if host already exists")
	}
	if hostExists {
		return nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	sshConfigFile := filepath.Join(homeDir, ".ssh", "config")
	var sshConfigBuilder strings.Builder
	cfgBytes, err := os.ReadFile(sshConfigFile)
	if err != nil {
		return errors.Wrapf(err, "failed to read %s file", sshConfigFile)
	}
	sshConfigBuilder.Write(cfgBytes)
	scmSshDir := filepath.Join(homeDir, ".ssh", "scm")
	hostConfig, err := file.RenderTmplt(&HostConfigInput{
		HostName:  hostname,
		ScmSshDir: scmSshDir,
	}, HostConfigTemplate)
	if err != nil {
		return errors.Wrapf(err, "failed to render host config template for %s host", hostname)
	}
	sshConfigBuilder.Write(hostConfig)
	if err := os.WriteFile(sshConfigFile, []byte(sshConfigBuilder.String()), 0744); err != nil {
		return errors.Wrapf(err, "failed to initialize %s file", sshConfigFile)
	}
	return nil
}

func DelHost(hostname string) error {
	hostExists, err := host.IsExists(hostname)
	if err != nil {
		return errors.Wrap(err, "failed to check if host already exists")
	}
	if !hostExists {
		return nil
	}
	var sshConfigBuilder strings.Builder
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	sshConfigFile := filepath.Join(homeDir, ".ssh", "config")
	f, err := os.Open(sshConfigFile)
	if err != nil {
		return errors.Wrapf(err, "failed to open %s file", sshConfigFile)
	}
	cfg, err := ssh_config.Decode(f)
	if err != nil {
		return errors.Wrapf(err, "failed to decode ssh config file %s", sshConfigFile)
	}
	for _, h := range cfg.Hosts {
		if host.Match(hostname, h) {
			continue
		}
		sshConfigBuilder.Write([]byte(h.String()))
	}
	if err := os.WriteFile(sshConfigFile, []byte(sshConfigBuilder.String()), 0744); err != nil {
		return errors.Wrapf(err, "failed to initialize %s file", sshConfigFile)
	}
	return nil
}

func getDefaultConfig() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get home dir")
	}
	scmSshDir := filepath.Join(homeDir, ".ssh", "scm")
	if err := os.MkdirAll(scmSshDir, 0744); err != nil {
		return "", errors.Wrapf(err, "failed to ensure %s dir", scmSshDir)
	}
	var sshConfigBuilder strings.Builder
	for _, h := range host.DefaultHosts() {
		hostConfig, err := file.RenderTmplt(&HostConfigInput{
			HostName:  h,
			ScmSshDir: scmSshDir,
		}, HostConfigTemplate)
		if err != nil {
			return "", errors.Wrapf(err, "failed to render host config template for %s host", h)
		}
		sshConfigBuilder.Write(hostConfig)
	}
	return sshConfigBuilder.String(), nil
}
