package plugin

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkgSwiftProtobufPlugin = "swift-protobuf"
	BrewPkgSwiftGrpcPlugin     = "grpc-swift"
)

func Setup() error {
	log.Infof("ensuring swift-protobuf plugin")
	if err := brew.Install(BrewPkgSwiftProtobufPlugin); err != nil {
		return errors.Wrapf(err, "failed to ensure %s brew pkg", BrewPkgSwiftProtobufPlugin)
	}
	log.Infof("ensured swift-protobuf plugin")
	log.Infof("ensuring grpc-swift plugin")
	if err := brew.Install(BrewPkgSwiftGrpcPlugin); err != nil {
		return errors.Wrapf(err, "failed to ensure %s brew pkg", BrewPkgSwiftGrpcPlugin)
	}
	log.Infof("ensured grpc-swift plugin")
	return nil
}

func Upgrade() error {
	if err := brew.Upgrade(BrewPkgSwiftProtobufPlugin); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s brew pkg", BrewPkgSwiftProtobufPlugin)
	}
	if err := brew.Upgrade(BrewPkgSwiftGrpcPlugin); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s brew pkg", BrewPkgSwiftGrpcPlugin)
	}
	return nil
}
