package lang

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/code/lang/golang"
	"github.com/leftbin/mactl/internal/app/build/code/lang/java"
	"github.com/leftbin/mactl/internal/app/build/code/lang/javascript"
	"github.com/leftbin/mactl/internal/app/build/code/lang/javascript/nodejs"
	"github.com/leftbin/mactl/internal/app/build/code/lang/python"
	"github.com/leftbin/mactl/internal/app/build/code/lang/sql"
	"github.com/leftbin/mactl/internal/app/build/code/lang/swift"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring golang")
	if err := golang.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure golang")
	}
	log.Info("ensured golang")
	log.Info("ensuring javascript")
	if err := javascript.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure javascript")
	}
	log.Info("ensured javascript")
	log.Info("ensuring python")
	if err := python.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure python")
	}
	log.Info("ensured python")
	//log.Info("ensuring battenberg")
	//if err := battenberg.Setup(); err != nil {
	//	return errors.Wrap(err, "failed to ensure battenberg")
	//}
	//log.Info("ensured battenberg")
	//log.Info("ensuring cookiecutter")
	//if err := cookiecutter.Setup(); err != nil {
	//	return errors.Wrap(err, "failed to ensure cookiecutter")
	//}
	//log.Info("ensured cookiecutter")
	log.Info("ensuring java")
	if err := java.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure java")
	}
	log.Info("ensured java")
	log.Info("ensuring sql")
	if err := sql.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure sql")
	}
	log.Info("ensured sql")
	log.Info("ensuring swift")
	if err := swift.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure swift")
	}
	log.Info("ensured swift")
	return nil
}

func Upgrade() error {
	log.Info("upgrading golang")
	if err := golang.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade golang")
	}
	log.Info("upgraded golang")

	log.Info("upgrading nodejs")
	if err := nodejs.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade nodejs")
	}
	log.Info("upgraded nodejs")

	log.Info("upgrading python")
	if err := python.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade python")
	}
	log.Info("upgraded python")

	// java is not being upgraded. need to discuss and will come up with strategy to upgrade java version
	return nil
}
