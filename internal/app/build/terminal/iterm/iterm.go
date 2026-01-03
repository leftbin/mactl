package iterm

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/terminal/mcfly"
	"github.com/leftbin/mactl/internal/cli/cache"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"github.com/leftbin/mactl/internal/lib/plist"
	"github.com/leftbin/mactl/internal/lib/shell"
	"github.com/leftbin/mactl/internal/zshrc/default"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	AppName                   = "iterm2"
	MacAppFileName            = "iTerm.app"
	AppGroupId                = "com.googlecode.iterm2"
	BrewPkg                   = "iterm2"
	OhMyZshInstallScriptUrl   = "https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh"
	ColorSchemesGitCloneUrl   = "https://github.com/mbadolato/iTerm2-Color-Schemes.git"
	PowerlineFontsGitCloneUrl = "https://github.com/powerline/fonts.git"
	// BrewPkgKubeContextDisplayZshPlugin https://github.com/jonmosco/kube-ps1
	BrewPkgKubeContextDisplayZshPlugin = "kube-ps1"
)

var (
	ZshPlugins = map[string]string{
		"zsh-completions":         "https://github.com/zsh-users/zsh-completions.git",
		"zsh-autosuggestions":     "https://github.com/zsh-users/zsh-autosuggestions.git",
		"zsh-syntax-highlighting": "https://github.com/zsh-users/zsh-syntax-highlighting.git",
	}
	ZshThemes = map[string]string{
		"powerlevel10k": "https://github.com/romkatv/powerlevel10k.git",
	}
)

func Setup() error {
	log.Info("setting up iterm")
	log.Info("ensuring iterm installed")
	if err := Install(); err != nil {
		return errors.Wrap(err, "failed to ensure iterm installed")
	}
	log.Info("ensured iterm installed")
	log.Info("ensuring fonts")
	if err := setupFonts(); err != nil {
		return errors.Wrap(err, "failed to ensure fonts")
	}
	log.Info("ensured fonts")
	log.Info("ensuring iterm configured")
	if err := Configure(); err != nil {
		return errors.Wrap(err, "failed to ensure iterm configured")
	}
	log.Info("ensured iterm configured")
	log.Info("ensuring zsh")
	if err := installZsh(); err != nil {
		return errors.Wrap(err, "failed to ensure zsh")
	}
	log.Info("ensured zsh")
	log.Info("ensuring zsh themes")
	if err := installZshThemes(); err != nil {
		return errors.Wrap(err, "failed to ensure zsh plugins")
	}
	log.Info("ensured zsh themes")
	log.Info("ensuring zsh plugins")
	if err := installZshPlugins(); err != nil {
		return errors.Wrap(err, "failed to ensure zsh plugins")
	}
	log.Info("ensured zsh plugins")
	log.Info("ensuring mcfly")
	if err := mcfly.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure mcfly")
	}
	log.Info("ensured mcfly")
	log.Info("ensuring zshrc file")
	if err := _default.Create(); err != nil {
		return errors.Wrap(err, "failed to ensure zshrc file")
	}
	log.Info("ensured zshrc file")
	log.Info("ensure powerlevel10k config file")
	if err := ensurePowerlevel10kConfig(); err != nil {
		return errors.Wrap(err, "failed to ensure powerlevel10k config file")
	}
	log.Info("ensure powerlevel10k config file")
	log.Info("setup complete for iterm")
	return nil
}

func Install() error {
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	return nil
}

func Upgrade() error {
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	return nil
}

func Configure() error {
	log.Infof("importing iterm config from plist file")
	if err := plist.Import(AppGroupId, PlistB64); err != nil {
		return errors.Wrapf(err, "failed to import %s plist file", AppGroupId)
	}
	log.Infof("imported iterm config from plist file")
	return nil
}

func ensurePowerlevel10kConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	configScriptLoc := filepath.Join(homeDir, ".p10k.zsh")
	p10ConfigScript, err := b64.StdEncoding.DecodeString(P10kConfigScriptB64)
	if err != nil {
		return errors.Wrap(err, "failed to decode base64 encoded config script file")
	}
	if err := os.WriteFile(configScriptLoc, p10ConfigScript, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", configScriptLoc)
	}
	return nil
}

func installZsh() error {
	cacheLoc, err := cache.GetLoc()
	if err != nil {
		return errors.Wrap(err, "failed to get cache loc")
	}
	scriptPath := fmt.Sprintf("%s/app/iterm/download/script/install-oh-my-zsh.sh", cacheLoc)
	if err := os.MkdirAll(filepath.Dir(scriptPath), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(scriptPath))
	}
	if err := file.Download(scriptPath, OhMyZshInstallScriptUrl); err != nil {
		return errors.Wrapf(err, "failed to download file %s", OhMyZshInstallScriptUrl)
	}
	if err := os.Setenv("CHSH", "no"); err != nil {
		return errors.Wrap(err, "failed to setup CHSH env var")
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home dir")
	}
	ohMyZshRootDir := filepath.Join(homeDir, ".oh-my-zsh")
	if err := os.RemoveAll(ohMyZshRootDir); err != nil {
		return errors.Wrapf(err, "failed to clean up %s dir", ohMyZshRootDir)
	}
	if err = shell.RunScript(scriptPath, cacheLoc, append(os.Environ(), []string{"CHSH=no", "RUNZSH=no", "KEEP_ZSHRC=yes"}...)); err != nil {
		return errors.Wrap(err, "failed to run script")
	}
	return nil
}

func installZshThemes() error {
	for name, gitUrl := range ZshThemes {
		if err := installZshTheme(name, gitUrl); err != nil {
			return errors.Wrapf(err, "failed to install %s zsh theme", name)
		}
	}
	return nil
}

func installZshTheme(name, gitUrl string) error {
	log.Infof("installing zsh theme %s", name)
	err := shell.RunCmd(exec.Command("gitr", "clone", gitUrl))
	if err != nil {
		return errors.Wrapf(err, "failed to clone zsh theme %s using %s url", name, gitUrl)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	themeCloneLoc := fmt.Sprintf("%s/scm/github.com/romkatv/%s", homeDir, name) //TODO: this can be retrieved using gitr
	zshThemeRoot := fmt.Sprintf("%s/.oh-my-zsh/custom/themes", homeDir)
	zshThemeLoc := fmt.Sprintf("%s/%s", zshThemeRoot, name)
	if err := shell.RunCmd(exec.Command("mv", themeCloneLoc, zshThemeLoc)); err != nil {
		return errors.Wrapf(err, "failed to move zsh theme %s from %s to %s", name, themeCloneLoc, zshThemeLoc)
	}
	log.Infof("installed zsh theme %s", name)
	return nil
}

func installZshPlugins() error {
	for name, gitUrl := range ZshPlugins {
		if err := installZshPlugin(name, gitUrl); err != nil {
			return errors.Wrapf(err, "failed to install %s zsh plugin", name)
		}
	}
	if err := brew.Install(BrewPkgKubeContextDisplayZshPlugin); err != nil {
		return errors.Wrapf(err, "failed to install zsh-plugin %s", BrewPkgKubeContextDisplayZshPlugin)
	}
	return nil
}

func installZshPlugin(name, gitUrl string) error {
	log.Infof("installing zsh plugin %s", name)
	err := shell.RunCmd(exec.Command("gitr", "clone", gitUrl))
	if err != nil {
		return errors.Wrapf(err, "failed to clone zsh plugin %s using %s url", name, gitUrl)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	pluginCloneLoc := fmt.Sprintf("%s/scm/github.com/zsh-users/%s", homeDir, name) //TODO: this can be retrieved using gitr
	zshPluginRoot := fmt.Sprintf("%s/.oh-my-zsh/custom/plugins", homeDir)
	zshPluginLoc := fmt.Sprintf("%s/%s", zshPluginRoot, name)
	if err := shell.RunCmd(exec.Command("mv", pluginCloneLoc, zshPluginLoc)); err != nil {
		return errors.Wrapf(err, "failed to move zsh plugin %s from %s to %s", name, pluginCloneLoc, zshPluginLoc)
	}
	log.Infof("installed zsh plugin %s", name)
	return nil
}

func cloneColorSchemes() error {
	err := shell.RunCmd(exec.Command("gitr", "clone", ColorSchemesGitCloneUrl))
	if err != nil {
		return errors.Wrapf(err, "error while cloning color schemes using %s url", ColorSchemesGitCloneUrl)
	}
	return nil
}

func setupFonts() error {
	log.Infof("cloning powerline fonts")
	if err := shell.RunCmd(exec.Command("gitr", "clone", PowerlineFontsGitCloneUrl)); err != nil {
		return errors.Wrap(err, "failed to clone fonts repo")
	}
	log.Infof("cloned powerline fonts")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	fontsRepoPath := fmt.Sprintf("%s/scm/github.com/powerline/fonts", homeDir)
	log.Info("installing powerline fonts")
	err = shell.RunScript(fmt.Sprintf("%s/install.sh", fontsRepoPath), fontsRepoPath, []string{})
	if err != nil {
		return errors.Wrap(err, "failed to install powerline fonts")
	}
	log.Info("installed powerline fonts")
	return nil
}

func GetPath() string {
	return macapp.GetPath(MacAppFileName)
}
