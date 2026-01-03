package plugin

import (
	"fmt"
	"github.com/pkg/errors"
	golangenv "github.com/leftbin/mactl/internal/app/build/code/lang/golang/env"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os"
	"os/exec"
)

const (
	ProtocGenGoPluginVersion     = "v1.28"
	ProtocGenGoGrpcPluginVersion = "v1.1"
)

var (
	GoPlugin   = fmt.Sprintf("google.golang.org/protobuf/cmd/protoc-gen-go@%s", ProtocGenGoPluginVersion)
	GrpcPlugin = fmt.Sprintf("google.golang.org/grpc/cmd/protoc-gen-go-grpc@%s", ProtocGenGoGrpcPluginVersion)
)

func Setup() error {
	goPathLoc, err := golangenv.GetGoPathLoc()
	if err != nil {
		return errors.Wrapf(err, "failed to get gopath location")
	}
	//gopath environment variable should be set before running go install
	if err := os.Setenv(golangenv.GoPathEnvVarKey, goPathLoc); err != nil {
		return errors.Wrapf(err, "failed to set gopath environment variable")
	}
	if err := shell.RunCmd(exec.Command("go", "install", GoPlugin)); err != nil {
		return errors.Wrap(err, "failed to install go protoc plugin")
	}
	if err := shell.RunCmd(exec.Command("go", "install", GrpcPlugin)); err != nil {
		return errors.Wrap(err, "failed to install go grpc protoc plugin")
	}
	return nil
}
