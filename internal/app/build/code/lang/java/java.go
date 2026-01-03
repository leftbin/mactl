package java

import (
	"fmt"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/leftbin/mactl/internal/app/build/code/lang/java/intellij"
	"github.com/leftbin/mactl/internal/app/build/code/lang/java/protoc/plugin"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Version string

const (
	BrewPkgJenv                = "jenv"
	BrewPkgJavaPrefix          = "openjdk"
	BrewPkgMaven               = "maven"
	BrewPkgGradle              = "gradle"
	Version11          Version = "11"
	Version17          Version = "17"
	DefaultJavaVersion         = "17.0"
)

func Setup() error {
	log.Infof("ensure jenv")
	if err := ensureJenv(); err != nil {
		return errors.Wrap(err, "failed to ensure jenv")
	}
	log.Infof("ensured jenv")
	log.Infof("ensuring java versios")
	if err := ensureJavaVersions(); err != nil {
		return errors.Wrap(err, "failed to ensure java versions")
	}
	log.Infof("ensured java versions")
	log.Infof("ensuring default java version")
	if err := ensureDefaultJavaVersion(); err != nil {
		return errors.Wrap(err, "failed to ensure default java version")
	}
	log.Infof("ensured default java version")
	log.Infof("ensuring java dependency managers")
	if err := ensureDependencyManagers(); err != nil {
		return errors.Wrap(err, "failed to ensure java dependency managers")
	}
	log.Infof("ensured java dependency managers")
	log.Infof("ensuring java ides")
	if err := ensureIdes(); err != nil {
		return errors.Wrap(err, "failed to ensure ides")
	}
	log.Infof("ensured java ides")
	log.Infof("ensuring protoc plugins")
	if err := plugin.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure protoc plugins")
	}
	log.Infof("ensured protoc plugins")
	return nil
}

func ensureJenv() error {
	err := brew.Install(BrewPkgJenv)
	if err != nil {
		return errors.Wrap(err, "failed to ensure jenv")
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	versionsDir := filepath.Join(homeDir, ".jenv", "versions")
	if !file.IsDirExists(versionsDir) {
		if err := os.MkdirAll(versionsDir, os.ModePerm); err != nil {
			return errors.Wrapf(err, "failed to ensure %s dir", versionsDir)
		}
	}
	return nil
}

func ensureDependencyManagers() error {
	log.Infof("installing maven")
	if err := brew.Install(BrewPkgMaven); err != nil {
		return errors.Wrap(err, "failed to install maven")
	}
	log.Infof("installed maven")
	log.Infof("installing gradle")
	if err := brew.Install(BrewPkgGradle); err != nil {
		return errors.Wrap(err, "failed to install gradle")
	}
	log.Infof("installed gradle")
	return nil
}

func ensureIdes() error {
	log.Infof("ensuring intellij")
	if err := intellij.Setup(); err != nil {
		return errors.Wrapf(err, "failed to ensure intellij")
	}
	log.Infof("ensured intellij")

	return nil
}

func ensureJavaVersions() error {
	log.Infof("ensuring java version %s", Version11)
	if err := ensureJavaVersion(Version11); err != nil {
		return errors.Wrapf(err, "failed to ensure java version %s", Version11)
	}
	log.Infof("ensured java version %s", Version11)
	log.Infof("ensuring java version %s", Version17)
	if err := ensureJavaVersion(Version17); err != nil {
		return errors.Wrapf(err, "failed to ensure java version %s", Version17)
	}
	log.Infof("ensured java version %s", Version17)
	return nil
}

func ensureJavaVersion(version Version) error {
	log.Infof("installing java %s version", version)
	brewPkg := fmt.Sprintf("%s@%s", BrewPkgJavaPrefix, version)
	if err := brew.Install(brewPkg); err != nil {
		return errors.Wrapf(err, "failed to install java brew pkg %s", brewPkg)
	}
	log.Infof("installed java %s version", version)
	openJdkMacOsLoc := fmt.Sprintf("/usr/local/opt/openjdk@%s/libexec/openjdk.jdk/Contents/Home", version)
	if runtime.GOARCH == "arm64" {
		openJdkMacOsLoc = fmt.Sprintf("/opt/homebrew/opt/openjdk@%v/libexec/openjdk.jdk/Contents/Home", version)
	}
	log.Infof("adding java %s version to jenv", version)
	if err := shell.RunCmd(exec.Command("jenv", "add", openJdkMacOsLoc)); err != nil {
		return errors.Wrapf(err, "failed to add java %v version to jenv", version)
	}
	log.Infof("added java %s version to jenv", version)
	return nil
}

func ensureDefaultJavaVersion() error {
	if err := shell.RunCmd(exec.Command("jenv", "global", string(DefaultJavaVersion))); err != nil {
		return errors.Wrapf(err, "failed to set %s as the default java version", DefaultJavaVersion)
	}
	log.Infof("configured java %s as the default java version", DefaultJavaVersion)
	return nil
}
