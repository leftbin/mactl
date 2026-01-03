package app

import (
	"github.com/jedib0t/go-pretty/v6/table"
	lbnutil "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/audio/krisp"
	"github.com/leftbin/mactl/internal/app/audio/shush"
	"github.com/leftbin/mactl/internal/app/browser/chrome"
	"github.com/leftbin/mactl/internal/app/build/cloud/gcloud"
	"github.com/leftbin/mactl/internal/app/build/code/lang/java"
	"github.com/leftbin/mactl/internal/app/build/network/dns/dnsmasq"
	"github.com/leftbin/mactl/internal/app/build/scm/battenberg"
	"github.com/leftbin/mactl/internal/app/build/scm/github"
	"github.com/leftbin/mactl/internal/app/build/scm/gitr"
	"github.com/leftbin/mactl/internal/app/build/terminal/iterm"
	"github.com/leftbin/mactl/internal/app/installer/mas"
	"github.com/leftbin/mactl/internal/app/keyboard/karabiner"
	"github.com/leftbin/mactl/internal/app/keyboard/rectangle"
	"github.com/leftbin/mactl/internal/app/network/telnet"
	"github.com/leftbin/mactl/internal/app/tidy/menubar/bartender"
	"github.com/leftbin/mactl/internal/app/tidy/menubar/toothfairy"
	"github.com/leftbin/mactl/internal/app/tool/flycut"
	"github.com/leftbin/mactl/internal/app/tool/mactl"
	"github.com/leftbin/mactl/internal/app/tool/sleep_control_center"
	"github.com/leftbin/mactl/internal/app/tool/snagit"
)

type App string

const (
	bartenderApp          App = "bartender"
	battenbergApp         App = "battenberg"
	chromeApp             App = "chrome"
	dnsmasqApp            App = "dnsmasq"
	flycutApp             App = "flycut"
	gcloudApp             App = "gcloud"
	githubApp             App = "github"
	gitrApp               App = "gitr"
	itermApp              App = "iterm"
	javaApp               App = "java"
	karabinerApp          App = "karabiner"
	krispApp              App = "krisp"
	mactlApp              App = "mactl"
	masApp                App = "mas"
	rectangleApp          App = "rectangle"
	shushApp              App = "shush"
	sleepControlCenterApp App = "sleep-control-center"
	snagitApp             App = "snagit"
	telnetApp             App = "telnet"
	toothFairyApp         App = "tooth-fairy"
	INVALID               App = "invalid"
)

func Get(app string) App {
	switch app {
	case
		"bartender",
		"battenberg",
		"chrome",
		"dnsmasq",
		"flycut",
		"gcloud",
		"github",
		"gitr",
		"iterm",
		"java",
		"karabiner",
		"keyboard-maestro",
		"krisp",
		"mactl",
		"mas",
		"rectangle",
		"shush",
		"sleep-control-center",
		"snagit",
		"telnet",
		"tooth-fairy":
		return App(app)
	}
	return INVALID
}

func (e App) IsInvalid() bool {
	return e == INVALID
}

func PrintList() {
	apps := []App{
		bartenderApp,
		battenbergApp,
		chromeApp,
		dnsmasqApp,
		flycutApp,
		gcloudApp,
		githubApp,
		gitrApp,
		itermApp,
		javaApp,
		karabinerApp,
		krispApp,
		mactlApp,
		masApp,
		rectangleApp,
		shushApp,
		sleepControlCenterApp,
		snagitApp,
		telnetApp,
		toothFairyApp,
	}
	header := table.Row{"#", "name"}
	rows := make([]table.Row, 0)
	for index, app := range apps {
		rows = append(rows, table.Row{index + 1, app})
	}
	lbnutil.PrintTable(header, rows)
}

func Add(app App) error {
	switch app {
	case bartenderApp:
		return bartender.Setup()
	case chromeApp:
		return chrome.Setup()
	case flycutApp:
		return flycut.Setup()
	case gcloudApp:
		return gcloud.Setup()
	case githubApp:
		return github.Setup()
	case gitrApp:
		return gitr.Setup()
	case itermApp:
		return iterm.Setup()
	case karabinerApp:
		if err := karabiner.Install(); err != nil {
			return errors.Wrapf(err, "failed to install %s", karabinerApp)
		}
		if err := karabiner.Configure(); err != nil {
			return errors.Wrapf(err, "failed to configure karabiner")
		}
		return nil
	case krispApp:
		return krisp.Setup()
	case mactlApp:
		return mactl.Setup()
	case masApp:
		return mas.Setup()
	case telnetApp:
		return telnet.Setup()
	case rectangleApp:
		return rectangle.Setup()
	case shushApp:
		return shush.Setup()
	case sleepControlCenterApp:
		return sleep_control_center.Setup()
	case snagitApp:
		return snagit.Setup()
	case toothFairyApp:
		return toothfairy.Setup()
	case battenbergApp:
		return battenberg.Setup()
	case dnsmasqApp:
		return dnsmasq.Setup()
	case javaApp:
		return java.Setup()
	}
	return nil
}
