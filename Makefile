# Define the UI folder path
UI_PATH=./ui/frontend/go-blockchain

# Check if Node.js version is 14.0 or higher
NODE_VERSION=$(shell node -v | cut -d. -f1 | sed 's/v//')
MIN_NODE_VERSION=14

# Check for react-scripts availability
CHECK_REACT_SCRIPTS=$(shell npm list -g | grep react-scripts || true)

all: check_node_version check_react_scripts build

# Check if Node.js version is >= 14.0
check_node_version:
	@if [ $(NODE_VERSION) -lt $(MIN_NODE_VERSION) ]; then \
		echo "Node.js v14.0 or higher is required. Current version: $(shell node -v)"; \
		exit 1; \
	fi

# Check if react-scripts is installed globally
check_react_scripts:
	@if [ -z "$(CHECK_REACT_SCRIPTS)" ]; then \
		echo "react-scripts is not installed globally. Please install it using 'npm install -g react-scripts'"; \
		exit 1; \
	fi

# Switch to UI directory and run npm commands
ui-build:
	@echo "Switching to UI folder: $(UI_PATH)"
	cd $(UI_PATH) && npm install && npm run build && npm start
