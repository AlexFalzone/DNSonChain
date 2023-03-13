//SPDX-License-Identifier: NONE

pragma solidity ^0.8.0;

interface dpki {
    
    struct Certificate {
        bytes certificateData;

        //salviamo l'hostname nella struct per semplificare la ricerca dei certificati
        string hostname;
    }

    function storeCertificate(bytes memory _certificateData) external returns (bool);

    //Prende in input l'hostname con il suffisso "speciale", per cercare il certificato corrispondente
    function getCertificateByHostname(string memory hostname) external view returns (bytes memory);

    //ritorna tutti i certificati revocati
    function getCRL() external pure returns (bytes[] memory);

    /*
        la lista dei certificati revocati prende in input 
        l'hostname con il suffisso "speciale" e restituisce 
        un array di bytes, dove ogni elemento è un certificato revocato
    */
    function revokeCertificate(string memory hostname) external  returns (bytes[] memory);
}
