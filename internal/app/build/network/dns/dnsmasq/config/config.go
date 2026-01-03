package config

import (
	"fmt"
	"github.com/pkg/errors"
	osint "github.com/leftbin/mactl/internal/lib/os"
	"github.com/leftbin/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"regexp"
)

const (
	// PlistFileIntelFormatterString todo: hardcoded version should be replaced
	PlistFileIntelFormatterString   = "/usr/local/Cellar/dnsmasq/%s/homebrew.mxcl.dnsmasq.plist"
	PlistFileSiliconFormatterString = "/opt/homebrew/Cellar/dnsmasq/%s/homebrew.mxcl.dnsmasq.plist"
	LaunchDaemonLoc                 = "/Library/LaunchDaemons/homebrew.mxcl.dnsmasq.plist"
)

func Setup() error {
	if err := setupLaunchDaemon(); err != nil {
		return errors.Wrapf(err, "failed to write config file")
	}
	return nil
}

func setupLaunchDaemon() error {
	log.Infof("ensuring launch daemon")
	installedVersion, err := getInstalledVersion()
	if err != nil {
		return errors.Wrapf(err, "failed to get installed dnsmasq version")
	}
	if err := shell.RunCmd(exec.Command("sudo", "cp", getPlistPath(installedVersion), LaunchDaemonLoc)); err != nil {
		return errors.Wrapf(err, "failed to copy launch daemon file from %s to %s", getPlistPath(""), LaunchDaemonLoc)
	}
	if err := shell.RunCmd(exec.Command("sudo", "launchctl", "load", LaunchDaemonLoc)); err != nil {
		return errors.Wrapf(err, "failed to copy launch daemon file from %s to %s", getPlistPath(installedVersion), LaunchDaemonLoc)
	}
	log.Infof("ensured launch daemon")
	return nil
}

func getInstalledVersion() (string, error) {
	out, err := exec.Command("sh", "-c", "brew info dnsmasq | grep 'dnsmasq:'").Output()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get installed dnsmasq version")
	}

	// Parse the output using regex
	re := regexp.MustCompile(`dnsmasq: stable (\S+)`)
	match := re.FindStringSubmatch(string(out))
	if len(match) <= 1 {
		return "", errors.Errorf("failed to parse dnsmasq version from output %s", string(out))
	}
	return match[1], nil
}

func getPlistPath(installedVersion string) string {
	if osint.GetArch() == osint.ARM64 {
		return fmt.Sprintf(PlistFileSiliconFormatterString, installedVersion)
	}
	return fmt.Sprintf(PlistFileIntelFormatterString, installedVersion)
}
