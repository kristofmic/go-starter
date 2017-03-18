# Startup a godoc server on http port 8080. Browse to
# http://localhost:8080/pkg/github.com/YikYakApp/platform/golang/
# to peruse the documentation.
godoc-http:
	godoc -http=:8080

protoc:
	protoc -I=./proto/example/v1 --go_out=./proto/example/v1 ./proto/example/v1/example.proto

# Install the libs and executables for platform.
install:
	glide install ./...

# Run lint against the code, ignore any errors reported against *.pb.go files.
# lint checks are not run on vendored dependencies
golint:
	@echo "gometalinter"
	@gometalinter.v1 \
	--deadline=60s \
	--vendor \
	--disable-all \
	--enable=golint \
	--enable=goimports \
	--enable=vet \
	--enable=deadcode \
	--enable=gosimple \
	--exclude=.*\.pb\.go \
	--exclude=bulk_data_test/github.com/* \
	--exclude=interaction_service/migrations/migrations.go \
	./...

# Run all tests under this project. If the directory contains a Makefile, use that to run the tests.
test: lint
	go test -tags=integration -cover -parallel 16 ./...

	# Install the libs and executables for linux on arch amd64.
xinstall-linux:
	GOOS=linux GOARCH=amd64 go install ./...

# Applies goimports to every go file (excluding vendored files)
goimports-fix:
	goimports -w $$(find . -type f -name '*.go' -not -path "*/vendor/*" )

build:
	GOOS=linux GOARCH=amd64 glide install ./...
	./scripts/build.sh
