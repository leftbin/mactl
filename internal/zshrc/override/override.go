package override

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const FileName = ".zshrc.override"

var ErrNotFound = errors.New("override file not found")

func Get() ([]byte, error) {
	f, err := GetPath()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get file path")
	}
	if !file.IsFileExists(f) {
		return []byte{}, nil
	}
	fileContent, err := os.ReadFile(f)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file contents")
	}
	return fileContent, nil
}

func Show() error {
	f, err := GetPath()
	if err != nil {
		return errors.Wrapf(err, "failed to get file path")
	}
	if !file.IsFileExists(f) {
		return ErrNotFound
	}
	if err := shell.RunCmd(exec.Command("cat", f)); err != nil {
		return errors.Wrapf(err, "failed to run command to open %s in vs code. is vs code installed?", f)
	}
	return nil
}

func Edit() error {
	f, err := GetPath()
	if err != nil {
		return errors.Wrapf(err, "failed to get file path")
	}
	if err := shell.RunCmd(exec.Command("code", f)); err != nil {
		return errors.Wrapf(err, "failed to run command to open %s in vs code. is vs code installed?", f)
	}
	return nil
}

func Append(content []byte) error {
	fileContent, err := Get()
	if err != nil {
		return errors.Wrapf(err, "failed to get %s file contents", FileName)
	}
	var overrideFileContentsBuilder strings.Builder
	overrideFileContentsBuilder.Write(fileContent)
	overrideFileContentsBuilder.Write(content)
	f, err := GetPath()
	if err != nil {
		return errors.Wrapf(err, "failed to get file path")
	}
	if err := os.WriteFile(f, []byte(overrideFileContentsBuilder.String()), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write updated overwrite contents")
	}
	return nil
}

func GetPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get home dir")
	}
	return filepath.Join(homeDir, FileName), nil
}
