
VERSION?=v1.1.0

DIR 		= dist
EXECUTABLE 	= ass-shifter
GOARCH		= amd64
GOOSWIN		= windows
GOOSX		= darwin
GOOSLINUX	= linux

WINBIN 		= $(DIR)/$(EXECUTABLE)-win-$(GOARCH).exe
OSXBIN 		= $(DIR)/$(EXECUTABLE)-darwin-$(GOARCH)
LINUXBIN 	= $(DIR)/$(EXECUTABLE)-linux-$(GOARCH)

CC 			= go build
LDFLAGS		= all=-w -s -X main.version=$(VERSION)

.PHONY: default all

default: all

all: darwin linux win64

.PHONY: darwin
darwin: $(OSXBIN)
	chmod +x $(OSXBIN)

.PHONY: $(OSXBIN)
$(OSXBIN):
	GOARCH=$(GOARCH) GOOS=$(GOOSX) $(CC) -o $(OSXBIN) -ldflags="$(LDFLAGS)"

.PHONY: linux
linux: $(LINUXBIN)
	chmod +x $(LINUXBIN)

.PHONY: $(LINUXBIN)
$(LINUXBIN):
	GOARCH=$(GOARCH) GOOS=$(GOOSLINUX) $(CC) -o $(LINUXBIN) -ldflags="$(LDFLAGS)"

.PHONY: win64
win64: $(WINBIN)

.PHONY: $(WINBIN)
$(WINBIN):
	GOARCH=$(GOARCH) GOOS=$(GOOSWIN) $(CC) -o $(WINBIN) -ldflags="$(LDFLAGS)"

.PHONY: clean
clean:
	rm ./$(DIR) -rf
