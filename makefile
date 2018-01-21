GO_CMD=go
GO_OS=windows
GO_ARCH=amd64
GO_BUILD=GOOS=$(GO_OS) GOARCH=$(GO_ARCH) $(GO_CMD) build

OUT = autorun.exe

SRC = main.go

all:
	$(GO_BUILD) -o $(OUT) $(SRC)

clean:
	$(GO_CMD) clean
	rm -f $(OUT)
