package service

const (
	getRecord = `[  {
    "inputs": [
      {
        "internalType": "string",
        "name": "_dataHash",
        "type": "string"
      }
    ],
    "name": "getRecord",
    "outputs": [
      {
        "internalType": "address",
        "name": "owner",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "timestamp",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "authorizedAddresses",
        "type": "address[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }]`

	addMedicalRecord = `[{
    "inputs": [
      {
        "internalType": "string",
        "name": "_dataHash",
        "type": "string"
      }
    ],
    "name": "addMedicalRecord",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }]`

	grantAccess = `[{
    "inputs": [
      {
        "internalType": "string",
        "name": "_dataHash",
        "type": "string"
      },
      {
        "internalType": "address",
        "name": "_address",
        "type": "address"
      }
    ],
    "name": "authorizeAddress",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }]`

	revokeAccess = `[
{
    "inputs": [
      {
        "internalType": "string",
        "name": "_dataHash",
        "type": "string"
      },
      {
        "internalType": "address",
        "name": "_address",
        "type": "address"
      }
    ],
    "name": "revokeAddress",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`
)
