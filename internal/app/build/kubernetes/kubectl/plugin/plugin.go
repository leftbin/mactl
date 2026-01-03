package plugin

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

var KubectlPlugins = []string{
	"node-shell",  //https://github.com/kvaps/kubectl-node-shell
	"capture",     //https://github.com/sysdiglabs/kubectl-capture
	"cost",        //https://github.com/kubecost/kubectl-cost
	"tree",        //https://github.com/ahmetb/kubectl-tree
	"view-secret", //https://github.com/elsesiy/kubectl-view-secret
}

func Setup() error {
	for _, plugin := range KubectlPlugins {
		log.Infof("ensuring %s kubectl plugin", plugin)
		if err := shell.RunCmd(exec.Command("kubectl", "krew", "install", plugin)); err != nil {
			return errors.Wrapf(err, "failed to install %s kubectl plugin", plugin)
		}
		log.Infof("ensured %s kubectl plugin", plugin)
	}
	return nil
}

func Upgrade() error {
	for _, plugin := range KubectlPlugins {
		log.Infof("upgrading %s kubectl plugin", plugin)
		if err := shell.RunCmd(exec.Command("kubectl", "krew", "upgrade", plugin)); err != nil {
			return errors.Wrapf(err, "failed to upgrade %s kubectl plugin", plugin)
		}
		log.Infof("upgraded %s kubectl plugin", plugin)
	}
	return nil
}
