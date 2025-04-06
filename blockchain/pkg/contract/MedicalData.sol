// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MedicalDataRegistry {
    struct MedicalRecord {
        string dataHash;
        address owner;
        uint256 timestamp;
        address[] authorizedAddresses; // Массив адресов, которые имеют доступ
    }

    mapping(string => MedicalRecord) public records;

    event RecordAdded(address indexed owner, string dataHash, uint256 timestamp);
    event AddressAuthorized(string dataHash, address indexed authorizedAddress);

    // Добавление медицинской записи
    function addMedicalRecord(string memory _dataHash) public {
        require(bytes(records[_dataHash].dataHash).length == 0, "Record already exists");

        // Создаём массив с двумя заранее определёнными адресами
        address[] memory initialAddresses = new address[](3); // Declare the array with size 3
        initialAddresses[0] = 0x51F1c37A80c5FD05De5C20801145aA8DF27e70dD;
        initialAddresses[1] = 0x283DeF31EeA8409Cd8168AA9B609E3B1BdCfe595;
        initialAddresses[2] = msg.sender; // Add owner as authorized

        records[_dataHash] = MedicalRecord({
            dataHash: _dataHash,
            owner: msg.sender,
            timestamp: block.timestamp,
            authorizedAddresses: initialAddresses // Сразу добавляем адреса в список
        });

        emit RecordAdded(msg.sender, _dataHash, block.timestamp);
        emit AddressAuthorized(_dataHash, msg.sender);
    }

    // Получение записи
    function getRecord(string memory _dataHash) public view returns (address owner, uint256 timestamp, address[] memory authorizedAddresses) {
        require(bytes(records[_dataHash].dataHash).length > 0, "Record not found");

        MedicalRecord memory record = records[_dataHash];
        return (record.owner, record.timestamp, record.authorizedAddresses);
    }

    // Добавление адреса в список авторизованных
    function authorizeAddress(string memory _dataHash, address _address) public {
        require(msg.sender == records[_dataHash].owner, "Only the owner can authorize addresses");
        records[_dataHash].authorizedAddresses.push(_address);
        emit AddressAuthorized(_dataHash, _address);
    }

    // Удаление адреса из списка авторизованных
    function revokeAddress(string memory _dataHash, address _address) public {
        require(msg.sender == records[_dataHash].owner, "Only the owner can revoke addresses");

        address[] storage authorizedAddresses = records[_dataHash].authorizedAddresses;
        for (uint i = 0; i < authorizedAddresses.length; i++) {
            if (authorizedAddresses[i] == _address) {
                authorizedAddresses[i] = authorizedAddresses[authorizedAddresses.length - 1];
                authorizedAddresses.pop();
                break;
            }
        }
    }
}
