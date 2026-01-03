package macapp

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os/exec"
)

type App struct {
	Name     string
	BrewPkg  string
	FileName string
}

func Open(appName string) error {
	log.Infof("opening %s", appName)
	if err := shell.RunCmd(exec.Command("open", GetPath(appName))); err != nil {
		return errors.Wrapf(err, "failed to open %s", appName)
	}
	log.Infof("opened %s", appName)
	return nil
}

func GetPath(appName string) string {
	return fmt.Sprintf("/Applications/%s", appName)
}
