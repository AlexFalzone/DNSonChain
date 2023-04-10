# Table of Contents

- [Table of Contents](#table-of-contents)
- [DNSonChain](#DNSonChain)
  - [Requirements](#requirements)
  - [Installation](#installation)
    - [Option 1: Manual Installation](#option-1-manual-installation)
    - [Option 2: Automated Installation using the Script](#option-2-automated-installation-using-the-script)
  - [Usage](#usage)
  - [Contributing](#contributing)
  - [License](#license)
  - [Contact](#contact)

<a name="DNSonChain"></a>
# DNSonChain

DNSonChain is an implementation of decentralized public key infrastructure using Ethereum blockchain for storing and retrieving certificates. 
Users can create certificates starting from a website with a special suffix like ".chain". 
The project leverages Ethereum to store and retrieve certificates and includes a mini DNS service that listens for requests. 
When the DNS service receives a type A request with a ".somet" hostname, it queries the Ethereum blockchain for the certificate, 
injects it into the certificate database, and communicates the information to the browser.

The main components of the DNSonChain project are:
- A custom DNS server that handles requests for ".chain" domains and forwards other requests to an external DNS server (e.g., Cloudflare's 1.1.1.1).
- A utility for managing certificates, such as generating, revoking, and renewing certificates.
- Functions to interact with the Ethereum blockchain using Infura or a local Ethereum node for certificate storage and retrieval.

<a name="requirements"></a>
## Requirements

- [Go](https://golang.org/doc/install) programming language (version 1.16 or later)
- libnss3-tools (on linux, very often it is already installed)
- An Execution layer client for the Ethereum Blockchain (currently supported: [Geth](https://geth.ethereum.org/downloads/))

<a name="installation"></a>
## Installation

<a name="manual-installation"></a>
### Option 1: Manual Installation

1. Clone the repository
```
git clone https://github.com/yourusername/DNSonChain.git
cd DNSonChain
```

2. Install dependencies
- For Debian-based systems (Ubuntu, Debian):
  ```
  sudo apt-get install -y dnsmasq libnss3-tools golang
  ```
- For Arch-based systems (Manjaro, Arch Linux):
  ```
  sudo pacman -S -y dnsmasq libnss3-tools golang
  ```
- For Fedora-based systems:
  ```
  sudo dnf install -y dnsmasq libnss3-tools golang
  ```

3. Initialize the Go module and install dependencies
```
go mod init test
go mod tidy
```

4. Build the project
```
go build -o DNSonChain main.go
```

5. Run the application
```
sudo ./DNSonChain
```

<a name="script-installation"></a>
### Option 2: Automated Installation using the Script

1. Save the provided script as `install.sh` in the project directory.
2. Give the script execution permissions:
   ```
   chmod +x install.sh
   ```
3. Run the script:
   ```
   ./install.sh
   ```
   The script will automatically detect your package manager, install dependencies, build the project, and prompt you to run the application.

<a name="usage"></a>
## Usage

There's a utility included that displays available entries, just run the program.

**Note**: Currently, to run the program, you need to stop the `systemd-resolved` service (or similar) and modify the `resolv.conf` configuration file (or similar). In the future, this won't be necessary as `dnsmasq` will be used.

1. Stop the `systemd-resolved` service:
   ```
   sudo service systemd-resolved stop
   ```

2. Modify the `resolv.conf` configuration file:
   - Open the file with a text editor (e.g., `nano` or `vim`):
     ```
     sudo nano /etc/resolv.conf
     ```
   - Replace the existing nameserver(s) with the IP address of your custom DNS server (e.g., `127.0.0.1`).
   - Save the changes and close the file.

3. Run the DPKI application:
   ```
   sudo ./dpki
   ```

4. To restore your original DNS settings, restart the `systemd-resolved` service with the following command:
   ```
   sudo service systemd-resolved start
   ```
   The `resolv.conf` file will be automatically restored to its original state during the restart of the `systemd-resolved` service.

Remember that these steps are temporary and will be changed once `dnsmasq` is implemented.

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

<a name="license"></a>
## License

This project is released under the GPL License.

## Contact

If you have any questions or need more information, please open an issue on GitHub or contact the project owner.
