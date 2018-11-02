pragma solidity ^0.4.24;

// IATI Organisations
// Steve Huckle

import "Organisations.sol";
import "Strings.sol";

contract IATIOrganisations is Organisations {

  struct Organisation {
    string name;
    string orgType;
  }

  string[] orgReferences;
  mapping(string => Organisation) private organisations;

  event OrganisationSet(string _reference, string _name, string _type);

  constructor() public {
  }

  function setOrganisation(string _reference, string _name, string _type) public {
    organisations[_reference].name = _name;
    organisations[_reference].orgType = _type;
    orgReferences.push(_reference);
    emit OrganisationSet(_reference, _name, _type);
  }

  function getOrganisationExists(string _reference) public constant returns (bool) {
    uint256 index = Strings.getIndex(_reference, orgReferences);
    return index != orgReferences.length;
  }

  function getNumOrganisations() public constant returns (uint256) {
    return orgReferences.length;
  }

  function getOrganisationReference(uint256 _index) public constant returns (string) {
    return orgReferences[_index];
  }

  function getOrganisationName(string _reference) public constant returns (string) {
    return organisations[_reference].name;
  }

  function getOrganisationType(string _reference) public constant returns (string) {
    return organisations[_reference].orgType;
  }
}
