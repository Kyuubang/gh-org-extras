install:
	@echo "Installing the project..."
	$(MAKE) build
	gh extension install .


build:
	@echo "Building the project..."
	go build -o ./gh-org-extras .