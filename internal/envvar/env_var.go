package envvar

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/lib/shell"
	"github.com/leftbin/mactl/internal/zshrc/override"
	"os/exec"
)

// List environment variables
func List() error {
	return shell.RunCmd(exec.Command("env"))
}

// Add a new environment variable
// if there is already an existing environment variable with same name, the value is overwritten with new value.
// environment variable is added by adding a new line to ~/.zshrc.override file
// duplicates are not checked, a new line is added with provided key value
func Add(key, val string) error {
	if key == "" {
		return errors.New("environment variable name can not be empty string")
	}
	if err := override.Append([]byte(fmt.Sprintf("\nexport %s=%s\n", key, val))); err != nil {
		return errors.Wrapf(err, "failed to write updated overwrite contents")
	}
	return nil
}
