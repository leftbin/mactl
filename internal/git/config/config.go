package config

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/leftbin/go-util/pkg/file"
	lbntable "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	sshcon "github.com/leftbin/mactl/internal/git/ssh/config"
	"os"
	"path/filepath"
	"text/template"
)

var ErrConfigNotFound = errors.New("config not found")
var ErrSshKeyPathNotFound = errors.New("ssh key path not found")

const (
	GitConfigFileName    = ".gitconfig"
	GlobalIgnoreFileName = ".gitignore_global"
)

type FilePath string

type Config struct {
	Host       string
	Username   string
	Email      string
	SshKeyPath FilePath
}

func Get(host, workspace string) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user home dir")
	}
	includeDir := filepath.Join(homeDir, "scm", host)
	if workspace != "" {
		includeDir = filepath.Join(homeDir, "scm", host, workspace)
	}
	gitConfigPath, err := getGitConfigForIncludeDir(includeDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get git config path for include dir %s", includeDir)
	}
	if !file.IsFileExists(gitConfigPath) {
		return nil, ErrConfigNotFound
	}
	gitConfig, err := parseConfig(gitConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read git config from %s file", gitConfigPath)
	}
	gitConfig.Host = host
	return gitConfig, nil
}

func CreGlobalIgnore() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	globalIgnoreFilePath := filepath.Join(homeDir, GlobalIgnoreFileName)
	if err := os.WriteFile(globalIgnoreFilePath, []byte(GlobalIgnore), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", globalIgnoreFilePath)
	}
	return nil
}

func Set(host, workspace, username, email, sshKeyPath string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home dir")
	}
	includesDir := filepath.Join(homeDir, "scm", host, workspace)
	gitConfigPath := filepath.Join(includesDir, GitConfigFileName)
	if err := creConfigFile(gitConfigPath, &Config{
		Host:       host,
		Username:   username,
		Email:      email,
		SshKeyPath: FilePath(sshKeyPath),
	}); err != nil {
		return errors.Wrapf(err, "failed to create git config file at %s", gitConfigPath)
	}
	if err := addIncludes(includesDir, gitConfigPath); err != nil {
		return errors.Wrap(err, "failed to update global config file")
	}
	log.Infof("ensuring %s host is in ssh config", host)
	if err := sshcon.AddHost(host); err != nil {
		return errors.Wrapf(err, "failed to ensure %s host is in ssh config", host)
	}
	log.Infof("ensured %s host is in ssh config", host)
	return nil
}

func Del(host, workspace string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home dir")
	}
	includesDir := filepath.Join(homeDir, "scm", host)
	if workspace != "" {
		includesDir = filepath.Join(homeDir, "scm", host, workspace)
	}
	if err := delIncludes(includesDir); err != nil {
		return errors.Wrap(err, "failed to delete includes")
	}
	return nil
}

func Use(host, username string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home dir")
	}
	includesDir := filepath.Join(homeDir, "scm", host)
	gitConfigPath := filepath.Join(includesDir, username, GitConfigFileName)
	if err := addIncludes(includesDir, gitConfigPath); err != nil {
		return errors.Wrap(err, "failed to add git includes")
	}
	return nil
}

func creConfigFile(gitConfigPath string, config *Config) error {
	tmpl := template.New("page")
	tmpl, err := tmpl.Parse(GitConfigTemplate)
	if err != nil {
		fmt.Println(err)
	}
	renderedBytes, err := file.RenderTmplt(config, GitConfigTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to render git config template")
	}
	if err := os.MkdirAll(filepath.Dir(gitConfigPath), 0744); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(gitConfigPath))
	}
	if err := os.WriteFile(gitConfigPath, renderedBytes, 0644); err != nil {
		return errors.Wrapf(err, "failed to write %s file", gitConfigPath)
	}
	return nil
}

func Print(c *Config) {
	rows := make([]table.Row, 0)
	rows = append(rows, table.Row{"host", c.Host})
	rows = append(rows, table.Row{"username", c.Username})
	rows = append(rows, table.Row{"email", c.Email})
	rows = append(rows, table.Row{"ssh-key", c.SshKeyPath})
	lbntable.PrintTable(nil, rows)
}

func getGlobalConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home dir")
	}
	gitConfigPath := filepath.Join(homeDir, GitConfigFileName)
	return gitConfigPath, err
}
