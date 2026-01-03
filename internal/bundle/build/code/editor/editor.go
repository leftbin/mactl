package editor

import (
	"github.com/pkg/errors" // Ensure you have added this dependency using 'go get github.com/pkg/errors'
	"github.com/leftbin/mactl/internal/app/build/code/vscode"
	"github.com/leftbin/mactl/internal/app/build/code/zed"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring editors")
	if err := vscode.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure vscode")
	}
	if err := zed.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure zed")
	}
	log.Info("ensured editors")
	return nil
}
