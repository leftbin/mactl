version?=v0.0.17
name=mactl
name_local=mactl
pkg=github.com/leftbin/mactl
build_dir=build
LDFLAGS=-ldflags "-X ${pkg}/internal/version.Version=${version}"
build_cmd=go build -v ${LDFLAGS}

.PHONY: deps
deps:
	go mod download

.PHONY: build
build: ${build_dir}/${name}

${build_dir}/${name}: deps
	GOOS=darwin ${build_cmd} -o ${build_dir}/${name}-darwin .
	GOOS=darwin GOARCH=amd64 ${build_cmd} -o ${build_dir}/${name}-darwin-amd64 .
	openssl dgst -sha256 ${build_dir}/${name}-darwin-amd64
	GOOS=darwin GOARCH=arm64 ${build_cmd} -o ${build_dir}/${name}-darwin-arm64 .
	openssl dgst -sha256 ${build_dir}/${name}-darwin-arm64
.PHONY: test
test:
	go test -race -v -count=1 ./...

.PHONY: run
run: build
	${build_dir}/${name}

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf ${build_dir}

checksum:
	@openssl dgst -sha256 ${build_dir}/${name}-darwin

local: build
	sudo rm -f /usr/local/bin/${name_local}
	sudo cp ./${build_dir}/${name}-darwin /usr/local/bin/${name_local}
	sudo chmod +x /usr/local/bin/${name_local}

.PHONY: release-github
release-github:
	git tag ${version}
	git push origin ${version}
	gh release create ${version} \
		 --generate-notes \
         --title ${version} \
         build/mactl-darwin-amd64 \
         build/mactl-darwin-arm64

release: build release-github
