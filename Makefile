GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=easywiki

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

install:
	sudo mkdir -p /etc/easywiki
	sudo cp -f easywiki.ini /etc/easywiki/
	sudo mkdir -p /var/easywiki/blogs
	sudo cp -rf template /var/easywiki/
	sudo cp -f $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

clean:
	$(GOCLEAN)
	sudo rm -f $(BINARY_NAME)
	sudo rm -rf /etc/easywiki
	sudo rm -rf /var/easywiki
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	sudo rm -rf /var/log/easywiki
