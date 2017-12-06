.PHONY: default
default: all ;

# default target
all: clean lint vet test build

# install build dependencies, REQUIRED
build-deps:
	go get -u -v github.com/golang/lint/golint

# install development dependencies, OPTIONAL
dev-deps:
	go get -u -v github.com/golang/dep/cmd/dep
	go get -u -v github.com/jteeuwen/go-bindata/...
	go get -u -v github.com/murlokswarm/macpack

# call golint on all packages except vendor folder
lint:
	for p in $$(go list ./...); do \
		golint $$p ;\
	done

# call go vet on all packages except vendor folder
vet:
	go vet ./...

# build executables on default arch
build:
	go build -v github.com/jkusniar/vetbbedit/cmd/vetbbedit

# build & install executables on default arch
install:
	go install github.com/jkusniar/vetbbedit/cmd/vetbbedit

# run all tests
test:
	go test -v ./...

# clean build artifacts
clean:
	go clean
	rm -fr vetbbedit