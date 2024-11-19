PROJECT_NAME = jobbr
VERSION = 1.0
IDENTIFIER = com.github.linkinlog.jobbr
RELEASE_DIR = ./build/release
EXECUTABLE = ./jobbr
PLIST = ./build/jobbr.plist

build: pkg
	rm -rf $(RELEASE_DIR)
	rm -f $(EXECUTABLE)

	@echo "\n\nBuild completed"

$(EXECUTABLE):
	CGO_ENABLED=0 GOOS=darwin go build -o $(EXECUTABLE) -ldflags "-s -w" .

$(RELEASE_DIR)/usr/local/bin/$(PROJECT_NAME): $(EXECUTABLE)
	mkdir -p $(RELEASE_DIR)/usr/local/bin
	cp $(EXECUTABLE) $(RELEASE_DIR)/usr/local/bin/

$(RELEASE_DIR)/Library/LaunchDaemons/$(PROJECT_NAME).plist: $(PLIST)
	mkdir -p $(RELEASE_DIR)/Library/LaunchDaemons
	cp $(PLIST) $(RELEASE_DIR)/Library/LaunchDaemons/${IDENTIFIER}.plist

pkg: $(EXECUTABLE) $(RELEASE_DIR)/usr/local/bin/$(PROJECT_NAME) $(RELEASE_DIR)/Library/LaunchDaemons/$(PROJECT_NAME).plist
	pkgbuild --root $(RELEASE_DIR) \
		--identifier $(IDENTIFIER) \
		--version $(VERSION) \
		--install-location / \
		./$(PROJECT_NAME).pkg

uninstall:
	rm -f /Library/LaunchDaemons/$(IDENTIFIER).plist
	rm -f /usr/local/bin/$(PROJECT_NAME)

.PHONY: build pkg uninstall
