package rectangle

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	"github.com/leftbin/mactl/internal/installer/macapp"
	"github.com/leftbin/mactl/internal/lib/plist"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg    = "rectangle"
	MacAppName = "Rectangle.app"
	AppGroupId = "com.knollsoft.Rectangle"
	PlistB64   = "YnBsaXN0MDDfEBQBAgMEBQYHCAkKCwwNDg8QERITFBUWGxwfICAiIyYnKi0wMyAgODk6XxAmTlNTdGF0dXNJdGVtIFByZWZlcnJlZCBQb3NpdGlvbiBJdGVtLTBXdG9wTGVmdF8QF1NVRW5hYmxlQXV0b21hdGljQ2hlY2tzWG1heGltaXplXxAhTlNOYXZQYW5lbEV4cGFuZGVkU2l6ZUZvck9wZW5Nb2RlXxATU1VIYXNMYXVuY2hlZEJlZm9yZV8QKE5TTmF2TGFzdFVzZXJTZXRIaWRlRXh0ZW5zaW9uQnV0dG9uU3RhdGVbbGFzdFZlcnNpb25YbGVmdEhhbGZfECVOU1dpbmRvdyBGcmFtZSBOU05hdlBhbmVsQXV0b3NhdmVOYW1lWXJpZ2h0SGFsZlpib3R0b21IYWxmWmJvdHRvbUxlZnRbYm90dG9tUmlnaHRXdG9wSGFsZl8QGWFsdGVybmF0ZURlZmF1bHRTaG9ydGN1dHNdbGF1bmNoT25Mb2dpbl8QF3N1YnNlcXVlbnRFeGVjdXRpb25Nb2RlXxAWTlNOYXZMYXN0Um9vdERpcmVjdG9yeVh0b3BSaWdodCJFulAA0hcYGRpXa2V5Q29kZV1tb2RpZmllckZsYWdzEBsSAB4AAAjSFx0eGl1tb2RpZmllckZsYWdzEANaezgwMCwgNDQ4fQkJUjU00hckJRpdbW9kaWZpZXJGbGFncxB7XxAdNjA0IDU1NSAzMjggMTk1IDAgMCAxNTM2IDkzNSDSFygpGl1tb2RpZmllckZsYWdzEHzSFyssGl1tb2RpZmllckZsYWdzEH3SFy4vGl1tb2RpZmllckZsYWdzECHSFzEyGl1tb2RpZmllckZsYWdzEB7SFzQ1Gl1tb2RpZmllckZsYWdzEH4JCRABW34vRG93bmxvYWRz0hc7PBpdbW9kaWZpZXJGbGFncxAYAAgAMwBcAGQAfgCHAKsAwQDsAPgBAQEpATMBPgFJAVUBXQF5AYcBoQG6AcMByAHNAdUB4wHlAeoB6wHwAf4CAAILAgwCDQIQAhUCIwIlAkUCSgJYAloCXwJtAm8CdAKCAoQCiQKXApkCngKsAq4CrwKwArICvgLDAtEAAAAAAAACAQAAAAAAAAA9AAAAAAAAAAAAAAAAAAAC0w=="
)

func Setup() error {
	log.Infof("installing rectangle")
	if err := install(); err != nil {
		return errors.Wrap(err, "failed to install")
	}
	log.Infof("installed rectangle")
	log.Infof("configuring rectangle")
	if err := configure(); err != nil {
		return errors.Wrap(err, "failed to install")
	}
	log.Infof("configured rectangle")
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open %s", MacAppName)
	}
	return nil
}

func install() error {
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	return nil
}

func Upgrade() error {
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	return nil
}

// configure will restore config using plistB64 file
func configure() error {
	log.Infof("importing rectlangle config from plist file")
	if err := plist.Import(AppGroupId, PlistB64); err != nil {
		return errors.Wrapf(err, "failed to import %s plist file", AppGroupId)
	}
	log.Infof("imported rectangle config from plist file")
	return nil
}
