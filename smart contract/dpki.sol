//SPDX-License-Identifier: NONE

pragma solidity ^0.8.0;

contract DomainRegistry {
    struct Domain {
        string hostname;
        string ip;
        string certificate;
        uint256 expiryDate;
        address owner;
    }

    modifier onlyDomainOwner(string memory hostname) {
        require(domains[hostname].owner == msg.sender, "Domain not owned by sender");
        _;
    }

    mapping(string => Domain) public domains;

    event DomainCreated(
        string hostname,
        string ip,
        string certificate,
        uint256 data,
        address indexed owner
    );

    function createDomain(
        string calldata hostname,
        string calldata ip,
        string memory certificate,
        uint256 expiryDate
    ) external {
        bytes memory hostnameBytes = bytes(hostname);
        uint256 dotCount = 0;
        for (uint256 i = 0; i < hostnameBytes.length; i++) {
            if (hostnameBytes[i] == '.') {
                dotCount++;
            }
        }
        require(dotCount > 0, "Invalid hostname");

        // In this case, we skip the root domain
        if (dotCount > 1) {
            // Find the parent domain
            string memory parentDomain = _getParentDomain(hostnameBytes);

            // Verify that the parent domain is registered and owned by the sender
            require(
                domains[parentDomain].owner == msg.sender,"Parent domain not registered or not owned by sender"
            );
        }

        /* 
            Check if hostname is already registered
            In the case of expired domain, delete the domain
            so it is possible to re-register the domain
        */ 
        if (bytes(domains[hostname].hostname).length > 0) {
            require(domains[hostname].expiryDate <= block.timestamp, "Hostname already registered and not expired");
            delete domains[hostname];
        }

        // Registrer the domain
        domains[hostname] = Domain({
            hostname: hostname,
            ip: ip,
            certificate: certificate,
            owner: msg.sender,

            // If the date is 0, the domain will never expire
            // otherwise, the domain will expire after the given number of seconds
            expiryDate: expiryDate == 0 ? type(uint256).max : block.timestamp + expiryDate
        });

        emit DomainCreated(hostname, ip, certificate, expiryDate, msg.sender);
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
        string certificate,
        uint256 expiryDate
    );

    function updateDomain(
        string calldata hostname,
        string calldata ip,
        string memory certificate,
        uint256 expiryDate
    ) external onlyDomainOwner(hostname) {
        require(bytes(domains[hostname].hostname).length > 0, "Domain not found");

        domains[hostname].ip = ip;
        domains[hostname].certificate = certificate;
        domains[hostname].expiryDate = expiryDate == 0 ? type(uint256).max : block.timestamp + expiryDate;

        emit DomainUpdated(hostname, ip, certificate, expiryDate);
    }

    event DomainDeleted(string hostname);
    function deleteDomain(string calldata hostname) external onlyDomainOwner(hostname) {
        require(bytes(domains[hostname].hostname).length > 0, "Domain not found");
        delete domains[hostname];

        emit DomainDeleted(hostname);
    }

    function getCertificate(string calldata hostname) external view returns (string memory, string memory) {
        Domain storage domain = domains[hostname];

        require(bytes(domain.hostname).length > 0, "Domain not found");
        require(domain.expiryDate > block.timestamp, "Certificate has expired");

        string memory certificate = domain.certificate;

        // If the certificate is odd, we need to add a 0 byte at the beginning
        if (bytes(certificate).length % 2 != 0) {
        certificate = string(abi.encodePacked("0", certificate));
    }

        return (domain.ip, certificate);
    }
}

    