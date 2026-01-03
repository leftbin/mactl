package hotcorner

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/cli/cache"
	"github.com/leftbin/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

const appleScript = `
property theSavedValues : {"-", "-", "-", "Lock Screen"} -- for example

tell application "System Preferences"
    activate
    set current pane to pane id "com.apple.preference.expose"
    tell application "System Events"
        tell window "Desktop & Screen Saver" of process "System Preferences"
            click button "Hot Corners…"
            tell sheet 1
                tell group 1
                    set theCurrentValues to value of pop up buttons
                    if theCurrentValues is {"-", "-", "-", "-"} then
                        repeat with i from 1 to 4
                            set thisValue to item i of theSavedValues
                            tell pop up button i
                                click
                                click menu item thisValue of menu 1
                            end tell
                        end repeat
                    else
                        copy theCurrentValues to theSavedValues
                        repeat with i from 1 to 4
                            tell pop up button i
                                click
                                click last menu item of menu 1
                            end tell
                        end repeat
                    end if
                end tell
                click button "OK"
            end tell
        end tell
    end tell
    quit
end tell
`

// Setup locking macbook by swiping finger to bottom right corner
// WARNING: this does not work anymore https://apple.stackexchange.com/a/108635
func Setup() error {
	log.Infof("setting up hot corner")
	cacheLoc, err := cache.GetLoc()
	if err != nil {
		return errors.Wrap(err, "failed to get cache loc")
	}
	scriptPath := filepath.Join(cacheLoc, "hack", "hot-corner", "apple-script.scpt")
	if err := os.MkdirAll(filepath.Dir(scriptPath), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(scriptPath))
	}
	if err := os.WriteFile(scriptPath, []byte(appleScript), os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to create script")
	}
	if err := shell.RunCmd(exec.Command("osascript", scriptPath)); err != nil {
		return errors.Wrap(err, "failed to run apple-script")
	}
	log.Infof("hot corner set up completed")
	return nil
}
