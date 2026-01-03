package bootstrap

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/browser/chrome"
	"github.com/leftbin/mactl/internal/app/build/terminal/iterm"
	"github.com/leftbin/mactl/internal/app/build/terminal/warp"
	"github.com/leftbin/mactl/internal/app/installer/mas"
	"github.com/leftbin/mactl/internal/app/keyboard/karabiner"
	"github.com/leftbin/mactl/internal/app/keyboard/rectangle"
	"github.com/leftbin/mactl/internal/app/tool/flycut"
	"github.com/leftbin/mactl/internal/app/tool/mactl"
	"github.com/leftbin/mactl/internal/bundle/build"
	"github.com/leftbin/mactl/internal/bundle/build/scm"
	"github.com/leftbin/mactl/internal/bundle/comm"
	"github.com/leftbin/mactl/internal/bundle/hotkey"
	"github.com/leftbin/mactl/internal/bundle/tool"
	"github.com/leftbin/mactl/internal/git/ssh"
	"github.com/leftbin/mactl/internal/optimize/dock"
	log "github.com/sirupsen/logrus"
)

func Checklist() {
	mandatory := []string{
		"optimize dock",
		"setup installers(brew & mas)",
		"setup hotkey",
		"setup window management",
		"setup browser profile",
		"setup git config & ssh-key",
		"setup iterm",
		"install comm apps",
		"install build apps",
		"setup env vars",
		"setup aliases",
	}
	optional := []string{
		"optimize trackpad",
		"setup hot corner to lock screen",
		"install app store apps",
	}
	fmt.Println("\nmandatory:")
	for index, item := range mandatory {
		fmt.Println(fmt.Sprintf("%d. %s", index+1, item))
	}
	fmt.Println("\noptional:")
	for index, item := range optional {
		fmt.Println(fmt.Sprintf("%d. %s", index+1, item))
	}
}

func Run() error {
	log.Info("bootstrap started")

	done := make(chan bool)
	fatalErrors := make(chan error)
	if err := ensure(fatalErrors); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to ensure all bootstrap components")
	}
	if err := open(fatalErrors); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to open all apps installed as part of bootstrap")
	}

	close(done)
	select {
	case <-done:
		break
	case err := <-fatalErrors:
		close(fatalErrors)
		log.Warnf("Error encountered : %v", err.Error())
		return err
	}

	log.Info("bootstrap completed")
	return nil
}

func ensure(fatalErrors chan error) error {
	log.Info("optimizing dock")
	if err := dock.Optimize(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to optimize dock")
	}
	log.Info("optimized dock")

	log.Info("ensuring installers")
	if err := mas.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup mas")
	}
	if err := mactl.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup mactl")
	}
	log.Info("ensured installers")

	log.Info("ensuring tool bundle")
	if err := tool.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup tool bundle")
	}
	log.Info("ensured tool bundle")

	log.Info("ensuring window-management")
	if err := rectangle.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup window management using rectangle app")
	}
	log.Info("ensured window-management")

	log.Info("ensuring browser")
	if err := chrome.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup browser")
	}
	log.Info("ensured browser")

	log.Info("ensuring scm")
	if err := scm.Setup(); err != nil {
		return errors.Wrap(err, "failed to setup scm")
	}
	log.Info("ensured scm")

	log.Info("ensuring comm apps")
	if err := comm.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup comm apps")
	}
	log.Info("ensured comm apps")
	log.Info("ensuring build bundle")
	if err := build.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup ides")
	}
	log.Info("ensured build bundle")
	log.Info("ensuring hotkey")
	if err := hotkey.Install(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup hotkey")
	}
	log.Info("ensured hotkey")
	log.Info("initializing ssh")
	if err := ssh.Init(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to initialize ssh")
	}
	log.Info("initialized ssh")
	log.Info("ensuring terminal")
	if err := iterm.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup iterm terminal")
	}
	if err := warp.Setup(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to setup warp terminal")
	}
	log.Info("ensured terminal")
	return nil
}

func open(fatalErrors chan error) error {
	log.Infof("opening karabiner-elements")
	if err := karabiner.Open(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to open karabiner-elements")
	}
	log.Infof("opened karabiner-elements")
	log.Infof("opening rectangle window manager app")
	if err := rectangle.Open(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to open rectangle window manager app")
	}
	log.Infof("opened rectangle window manager app")
	log.Infof("opening flycut")
	if err := flycut.Open(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to open flycut")
	}
	log.Infof("opened flycut")
	log.Infof("opening google-chrome")
	if err := chrome.Open(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to open google-chrome")
	}
	log.Infof("opened google-chrome")
	return nil
}
