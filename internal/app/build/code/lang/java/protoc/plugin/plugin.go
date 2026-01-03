package plugin

import (
	"fmt"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/leftbin/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	ProtocGenGrpcJavaPluginVersion = "1.65.0"
	BinaryName                     = "protoc-gen-grpc-java"
)

var (
	DownloadUrl = fmt.Sprintf("https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/%s/protoc-gen-grpc-java-%s-osx-x86_64.exe", ProtocGenGrpcJavaPluginVersion, ProtocGenGrpcJavaPluginVersion)
)

func Setup() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	downloadPath := filepath.Join(homeDir, "bin", BinaryName)
	if err := file.Download(downloadPath, DownloadUrl); err != nil {
		return errors.Wrapf(err, "failed to download file from %s url", DownloadUrl)
	}
	if err := shell.RunCmd(exec.Command("chmod", "+x", downloadPath)); err != nil {
		return errors.Wrapf(err, "failed to add executable permission to %s", downloadPath)
	}
	return nil
}
