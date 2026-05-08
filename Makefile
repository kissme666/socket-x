.PHONY: build build-all clean

APP    := socketx
BUILD  := build
LDFLAGS := -s -w

define compile
	CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) \
		go build -ldflags "$(LDFLAGS)" \
		-o $(BUILD)/$(APP)-$(1)-$(2)$(if $(filter windows,$(1)),.exe,) ./cmd/socketx
endef

build:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD)/$(APP) ./cmd/socketx

build-all:
	$(call compile,linux,amd64)
	$(call compile,linux,arm64)
	$(call compile,darwin,amd64)
	$(call compile,darwin,arm64)
	$(call compile,windows,amd64)

clean:
	rm -rf $(BUILD)/
