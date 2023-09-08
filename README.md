# Table of Contents

- [Table of Contents](#table-of-contents)
- [DNS-On-Chain](#dns-on-chain)
  - [Requirements](#requirements)
  - [Installation](#installation)
    - [Option 1: Manual Installation](#option-1-manual-installation)
  - [To-Do/Future Directions](#to-dofuture-directions)
  - [Usage](#usage)
  - [Contributing](#contributing)
  - [Contact](#contact)

<a name="DNSonChain"></a>
# DNS-On-Chain

DNSonChain is an implementation of decentralized public key infrastructure using Ethereum blockchain for storing and retrieving certificates. 
Users can create certificates starting from a website with a special suffix like ".chain". 
The project leverages Ethereum to store and retrieve certificates and includes a mini DNS service that listens for requests. 
When the DNS service receives a type A request with a ".chain" hostname, it queries the Ethereum blockchain for the certificate, 
injects it into the certificate database, and communicates the information to the browser.

The main components of the DNSonChain project are:
- A custom DNS server that handles requests for ".chain" domains and forwards other requests to an external DNS server (e.g., Cloudflare's 1.1.1.1).
- A utility for managing certificates, such as generating, revoking, and renewing certificates.
- Functions to interact with the Ethereum blockchain using Infura or a local Ethereum node for certificate storage and retrieval.

**Note**: the smart contract is included here for completeness only

<a name="requirements"></a>
## Requirements

- [Go](https://golang.org/doc/install) programming language (version 1.16 or later)
- libnss3-tools (on linux, very often it is already installed. On Windows, you can download it from here: [certutil](https://dist.torproject.org/torbrowser/)
- An Execution client should be sufficient for a light node. Otherwise you will also need a consensus client. (If your client uses port 8545, it is already supported)

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
3. Modify related files
- Create a 'chain.conf' file in the '/etc/dnsmasq.d' directory and add the following line:
  ```
  server=/chain/127.0.0.1#5354
  ```
- Modify the '/etc/NetworkManager/NetworkManager.conf' file to add the following line under the '[main]' section:
  ```
  dns=dnsmasq
  ```
- Modify the '/etc/resolv.conf' file to add the following line:
  *(You can delete the other lines if you want)*
  ```
  nameserver 127.0.0.1
- Restart the dnsmasq service:
  ```
  sudo service dnsmasq restart
  ```
  
4. Initialize the Go module and install dependencies
```
go mod init test
go mod tidy
```

5. Build the project
```
go build -o DNSonChain main.go
```

6. Run the application
```
sudo ./DNSonChain
```

<a name="script-installation"></a>
### Option 2: Automated Installation using the Script

 The script will automatically detect your package manager, install dependencies, modify related files, build the project, and prompt you to run the application.

1. Clone the repository
   ```
   git clone https://github.com/yourusername/DNSonChain.git
   cd DNSonChain
   ```

2. Give the script execution permissions:
   ```
   chmod +x scripLinux.sh
   ```
3. Run the script:
   ```
   ./scripLinux.sh
   ```
<a name="todo"></a>
## To-Do/Future Directions

- ~Full support to IPv6~
- A detailed description of the installation, for Windows
- A setup script for Windows
- In the future, DNS-On-Chain will not only handle ".chain" domains but also any domain. It will search the corresponding blockchain for the specific domain, and if it doesn't exist, it will revert the DNS request to an external DNS server. Additionally, a caching mechanism will be introduced for domain names. This way, if a domain name is present in the cache, the system will confidently make a request to the blockchain. If the domain name is not in the cache, the system will not even attempt a request.
- A browser extension will also be released in the future.


<a name="usage"></a>
## Usage

There's a utility included that displays available entries, just run the program.

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

## Contact

If you have any questions or need more information, please open an issue on GitHub or contact me.
