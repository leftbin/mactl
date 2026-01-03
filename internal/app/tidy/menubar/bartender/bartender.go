package bartender

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/cli/cache"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	DmgDownloadUrl      = "https://www.macbartender.com/B2/updates/B4Latest/Bartender%204.dmg"
	MacAppName          = "Bartender 4.app"
	MacApplicationsPath = "/Applications"
)

func Setup() error {
	log.Info("installing bartender")
	if err := install(); err != nil {
		return errors.Wrap(err, "failed to install bartender")
	}
	log.Info("installed bartender")
	return nil
}

func install() error {
	cacheLoc, err := cache.GetLoc()
	if err != nil {
		return errors.Wrap(err, "failed to get cache loc")
	}
	dmgDownloadPath := filepath.Join(cacheLoc, "app", "menubar", "bartender.dmg")
	if err := os.RemoveAll(filepath.Dir(dmgDownloadPath)); err != nil {
		return errors.Wrapf(err, "failed to cleanup %s dir", filepath.Dir(dmgDownloadPath))
	}
	if err := os.MkdirAll(filepath.Dir(dmgDownloadPath), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(dmgDownloadPath))
	}
	if err := file.Download(dmgDownloadPath, DmgDownloadUrl); err != nil {
		return errors.Wrapf(err, "failed to download bartender using %s url", DmgDownloadUrl)
	}
	log.Infof("mounting dmg file")
	if err := shell.RunCmd(exec.Command("hdiutil", "attach", dmgDownloadPath)); err != nil {
		return errors.Wrapf(err, "failed to mount %s file", dmgDownloadPath)
	}
	log.Infof("mounted dmg file")
	dmgMountPath := filepath.Join("/Volumes", "Bartender 4")
	src := filepath.Join(dmgMountPath, MacAppName)
	dest := filepath.Join(MacApplicationsPath, MacAppName)
	log.Infof("copying app to %s", dest)
	if err := shell.RunCmd(exec.Command("sudo", "cp", "-R", src, dest)); err != nil {
		return errors.Wrapf(err, "failed to move %s to %s", src, dest)
	}
	log.Infof("copied app to %s", dest)
	log.Infof("unmounting dmg file")
	if err := shell.RunCmd(exec.Command("hdiutil", "detach", dmgMountPath)); err != nil {
		return errors.Wrapf(err, "failed to unmount %s", dmgMountPath)
	}
	log.Infof("unmounted dmg file")
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrap(err, "failed to open bartender")
	}
	return nil
}
