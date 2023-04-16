#!/bin/bash

# Function to detect package manager and install dependencies
install_dependencies() {
    if command -v apt-get &> /dev/null ; then
        INSTALL_CMD="sudo apt-get install -y"
    elif command -v pacman &> /dev/null ; then
        INSTALL_CMD="sudo pacman -S -y"
    elif command -v dnf &> /dev/null ; then
        INSTALL_CMD="sudo dnf install -y"
    else
        echo "Unsupported package manager. Exiting."
        exit 1
    fi

    echo "Installing dependencies..."
    $INSTALL_CMD dnsmasq libnss3-tools golang
}

# Install dependencies
install_dependencies

# Build the Go application
echo "Building the Go application..."

if [ ! -f "go.mod" ] || [ ! -f "go.sum" ]; then
    go mod init test
    if [ $? -ne 0 ]; then
        echo "Error initializing the project. Exiting."
        exit 1
    fi

    go mod tidy
    if [ $? -ne 0 ]; then
        echo "Error fetching dependencies. Exiting."
        exit 1
    fi
fi

go build -o DNSonChain main.go
if [ $? -ne 0 ]; then
    echo "Error building the application. Exiting."
    exit 1
fi

echo "Installation and build complete."

# Configure dnsmasq
echo "Configuring dnsmasq..."
echo "server=/chain/127.0.0.1#5354" | sudo tee /etc/dnsmasq.d/chain.conf

# Update /etc/resolv.conf
echo "Updating /etc/resolv.conf..."
echo "nameserver 127.0.0.1" | sudo tee -a /etc/resolv.conf

# Update NetworkManager.conf
echo "Updating NetworkManager.conf..."
sudo sed -i '/^\[main\]$/a dns=dnsmasq' /etc/NetworkManager/NetworkManager.conf

# Restart dnsmasq and NetworkManager
echo "Restarting dnsmasq..."
sudo systemctl restart dnsmasq

# Ask the user if they want to run the app
read -p "Do you want to run the app? (y/n) " answer

if [ "$answer" == "y" ] || [ "$answer" == "Y" ]; then
    echo "Running the app..."
    sudo ./DNSonChain
else
    echo "Exiting."
    exit 0
fi
