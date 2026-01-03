package api

import (
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/app/build/api/grpc/client/grpcurl"
	"github.com/leftbin/mactl/internal/app/build/api/grpc/compiler/buf"
	"github.com/leftbin/mactl/internal/app/build/api/grpc/compiler/protobuf"
	"github.com/leftbin/mactl/internal/app/build/api/rest/client/postman"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Infof("ensuring rest-api client - postman")
	if err := postman.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure rest-api client - postman")
	}
	log.Infof("ensured rest-api client - postman")
	log.Infof("ensuring grpc api compilers")
	log.Infof("ensuring grpc api compiler - prtobuf")
	if err := protobuf.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api compiler - prtobuf")
	}
	log.Infof("ensured grpc api compiler - prtobuf")
	log.Infof("ensuring grpc api compiler - buf")
	if err := buf.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api compiler - buf")
	}
	log.Infof("ensured grpc api compiler - buf")
	log.Infof("ensured grpc api compilers")
	log.Infof("ensuring grpc api clients")
	log.Infof("ensured grpc api client - wombat")
	log.Infof("ensuring grpc api client - grpcurl")
	if err := grpcurl.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure grpc api client - grpcurl")
	}
	log.Infof("ensured grpc api client - grpcurl")
	log.Infof("ensured grpc api clients")
	return nil
}

func Upgrade() error {
	log.Infof("upgrading rest-api client - postman")
	if err := postman.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade rest-api client - postman")
	}
	log.Infof("upgraded rest-api client - postman")

	log.Infof("upgrading grpc api compilers")
	log.Infof("upgrading grpc api compiler - protoc")
	if err := protobuf.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api compiler - protoc")
	}
	log.Infof("upgraded grpc api compiler - protoc")

	log.Infof("upgrading grpc api compiler - buf")
	if err := buf.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api compiler - buf")
	}
	log.Infof("upgraded grpc api compiler - buf")
	log.Infof("upgraded grpc api compilers")

	log.Infof("upgrading grpc api clients")

	log.Infof("upgrading grpc api client - grpcurl")
	if err := grpcurl.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to upgrade grpc api client - grpcurl")
	}
	log.Infof("upgraded grpc api client - grpcurl")
	log.Infof("upgraded grpc api clients")
	return nil
}
