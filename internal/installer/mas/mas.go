package mas

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os/exec"
)

func Install(appStoreId string) error {
	if err := shell.RunCmd(exec.Command("mas", "install", appStoreId)); err != nil {
		return errors.Wrapf(err, "failed to Install with %s app store id", appStoreId)
	}
	return nil
}
