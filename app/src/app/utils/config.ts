class Paths {

  static readonly home='/'
  static readonly blockchain = '/blockchain-data'
  static readonly about='/about'
  static readonly overview='/overview'
  static readonly help='/help'
  static readonly writer='/create'
  static readonly reader='/read'
  static readonly orgWriter='/create-organisation'
  static readonly orgReader='/read-organisations'
}

class Blockchain {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'localhost'
  static readonly port = '8545'

  static readonly checkAccountInterval = 3000
}

class OrgsContract {

  static organisationsABI = [
    "event SetOrganisation(string _reference, string _name, string _type)",
    "function setOrganisation(string _reference, string _name, string _type)",
    "function getOrganisationExists(string _reference) constant returns (bool)",
    "function getNumOrganisations() constant returns (uint256)",
    "function getOrganisationReference(uint256 _index) constant returns (string)",
    "function getOrganisationName(string _reference) constant returns (string)",
    "function getOrganisationType(string _reference) constant returns (string)",
  ]

  /*
  organisations: 0x749f861de9e83807e0ebaadaedd88a2f645dc176
  */

  static organisationsAddress = "0x4a43ac2b58d0986bc528c7a13d52c07a53d9d9b0"
}

export { Paths, Blockchain, OrgsContract }
