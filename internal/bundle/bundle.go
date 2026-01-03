package bundle

import (
	"github.com/jedib0t/go-pretty/v6/table"
	lbnutil "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/bundle/audio"
	"github.com/leftbin/mactl/internal/bundle/build"
	"github.com/leftbin/mactl/internal/bundle/build/network"
	"github.com/leftbin/mactl/internal/bundle/comm"
	"github.com/leftbin/mactl/internal/bundle/hotkey"
	"github.com/leftbin/mactl/internal/bundle/menubar"
	"github.com/leftbin/mactl/internal/bundle/tool"
)

type Bundle string

const (
	audioBundle   Bundle = "audio"
	buildBundle   Bundle = "build"
	commBundle    Bundle = "comm"
	hotkeyBundle  Bundle = "hotkey"
	menubarBundle Bundle = "menubar"
	networkBundle Bundle = "network"
	toolBundle    Bundle = "tool"
	INVALID       Bundle = "invalid"
)

func Get(bundle string) Bundle {
	switch bundle {
	case "audio", "build", "comm", "hotkey", "menubar", "network", "tool":
		return Bundle(bundle)
	}
	return INVALID
}

func (e Bundle) IsInvalid() bool {
	return e == INVALID
}

func PrintList() {
	bundles := []Bundle{audioBundle, buildBundle, commBundle, hotkeyBundle, menubarBundle, networkBundle, toolBundle}
	header := table.Row{"#", "name"}
	rows := make([]table.Row, 0)
	for index, bundle := range bundles {
		rows = append(rows, table.Row{index + 1, bundle})
	}
	lbnutil.PrintTable(header, rows)
}

func Add(bundle Bundle) error {
	switch bundle {
	case audioBundle:
		return audio.Setup()
	case buildBundle:
		return build.Setup()
	case commBundle:
		return comm.Setup()
	case hotkeyBundle:
		if err := hotkey.Install(); err != nil {
			return errors.Wrapf(err, "failed to install %s", hotkeyBundle)
		}
		return hotkey.Configure()
	case menubarBundle:
		return menubar.Setup()
	case networkBundle:
		return network.Setup()
	case toolBundle:
		return tool.Setup()
	}
	return nil
}
