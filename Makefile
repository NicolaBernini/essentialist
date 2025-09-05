
# Makefile for essentialist and flashdown

# Go parameters
GOCMD=go

# Wayland support
WAYLAND ?= false
ifeq ($(WAYLAND), true)
	GOTAGS += wayland
endif

# Binaries
ESSENTIALIST_BIN=essentialist
FLASHDOWN_BIN=flashdown

# Install directories
PREFIX?=/usr/local
BINDIR?=$(PREFIX)/bin
DATADIR?=$(PREFIX)/share
APPDIR?=$(DATADIR)/applications
ICONSDIR?=$(DATADIR)/icons/hicolor/scalable/apps

all: $(ESSENTIALIST_BIN) $(FLASHDOWN_BIN)

$(ESSENTIALIST_BIN):
	$(GOCMD) build -tags="$(GOTAGS)" -o $(ESSENTIALIST_BIN) ./cmd/essentialist

$(FLASHDOWN_BIN):
	$(GOCMD) build -o $(FLASHDOWN_BIN) ./cmd/flashdown

clean:
	rm -f $(ESSENTIALIST_BIN) $(FLASHDOWN_BIN)

install: all
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 $(ESSENTIALIST_BIN) $(DESTDIR)$(BINDIR)
	install -m 755 $(FLASHDOWN_BIN) $(DESTDIR)$(BINDIR)
	install -d $(DESTDIR)$(APPDIR)
	install -m 644 cmd/essentialist/flatpak/io.github.essentialist_app.essentialist.desktop $(DESTDIR)$(APPDIR)/
	install -d $(DESTDIR)$(ICONSDIR)
	install -m 644 cmd/essentialist/flatpak/io.github.essentialist_app.essentialist.svg $(DESTDIR)$(ICONSDIR)/


.PHONY: all clean install
