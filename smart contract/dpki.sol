//SPDX-License-Identifier: NONE

pragma solidity ^0.8.0;

contract DomainRegistry {
    struct Domain {
        string hostname;
        string ip;
        bytes certificate;
        uint256 data;
        address owner;
    }

    mapping(string => Domain) public domains;

    event DomainCreated(
        string hostname,
        string ip,
        bytes certificate,
        uint256 data,
        address indexed owner
    );

    function createDomain(
        string calldata hostname,
        string calldata ip,
        bytes calldata certificate,
        uint256 data
    ) external {
        bytes memory hostnameBytes = bytes(hostname);
        uint256 dotCount = 0;
        for (uint256 i = 0; i < hostnameBytes.length; i++) {
            if (hostnameBytes[i] == '.') {
                dotCount++;
            }
        }
        require(dotCount > 0, "Invalid hostname");

        // Trova il dominio padre
        string memory parentDomain = _getParentDomain(hostnameBytes);

        // Verifica che il dominio padre sia registrato e che l'owner sia lo stesso
        require(
            domains[parentDomain].owner == msg.sender,"Parent domain not registered or not owned by sender"
        );

        // Registra il dominio
        domains[hostname] = Domain({
            hostname: hostname,
            ip: ip,
            certificate: certificate,
            owner: msg.sender,

            //se la data è 0, allora la data è il massimo valore di uint256
            //altrimenti la data è quella passata come parametro
            data: data == 0 ? type(uint256).max : block.timestamp + data
        });

        emit DomainCreated(hostname, ip, certificate, data, msg.sender);
    }

    function _getParentDomain(bytes memory hostnameBytes) internal pure returns (string memory) {
        uint256 lastDot = 0;
        for (uint256 i = 0; i < hostnameBytes.length; i++) {
            if (hostnameBytes[i] == '.') {
                lastDot = i;
                break;
            }
        }
        bytes memory parentDomainBytes = new bytes(hostnameBytes.length - lastDot - 1);
        for (uint256 i = 0; i < parentDomainBytes.length; i++) {
            parentDomainBytes[i] = hostnameBytes[lastDot + 1 + i];
        }
        return string(parentDomainBytes);
    }

    event DomainUpdated(
        string hostname,
        string ip,
        bytes certificate,
        uint256 data
    );

    function updateDomain(
        string calldata hostname,
        string calldata ip,
        bytes calldata certificate,
        uint256 data
    ) external {
        require(domains[hostname].owner == msg.sender, "Domain not owned by sender");

        domains[hostname].ip = ip;
        domains[hostname].certificate = certificate;
        domains[hostname].data = data == 0 ? type(uint256).max : block.timestamp + data;

        emit DomainUpdated(hostname, ip, certificate, data);
    }

    event DomainDeleted(string hostname);
    function deleteDomain(string calldata hostname) external {
        require(domains[hostname].owner == msg.sender, "Domain not owned by sender");

        delete domains[hostname];

        emit DomainDeleted(hostname);
    }

    function getCertificate(string calldata hostname) external view returns (bytes memory, string memory) {
        Domain storage domain = domains[hostname];

        require(bytes(domain.hostname).length > 0, "Domain not found");
        require(domain.data > block.timestamp, "Certificate has expired");

        return (domain.certificate, domain.ip);
    }
}

    