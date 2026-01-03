package plist

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/cli/cache"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
)

func Import(appGroupId, plistB64 string) error {
	cacheLoc, err := cache.GetLoc()
	if err != nil {
		return errors.Wrap(err, "failed to get cache loc")
	}
	plistFile := filepath.Join(cacheLoc, "plist", fmt.Sprintf("%s.plist", appGroupId))
	if err := os.MkdirAll(filepath.Dir(plistFile), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(plistFile))
	}
	plist, err := b64.StdEncoding.DecodeString(plistB64)
	if err != nil {
		return errors.Wrap(err, "failed to decode plist file")
	}
	if err := os.WriteFile(plistFile, plist, os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to create plist file")
	}
	if err := shell.RunCmd(exec.Command("defaults", "import", appGroupId, plistFile)); err != nil {
		return errors.Wrap(err, "failed to import plist file")
	}
	return nil
}
