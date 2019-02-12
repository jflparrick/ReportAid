pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationRegionBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationRegionBudgets is OrganisationRegionBudgets {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private budgetRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => RegionBudget))) private budgets;

  event SetRegionBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, RegionBudget _budget);

  function setRegionBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, RegionBudget memory _budget) public {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _budgetRef[0] != 0 &&
             _budget.regionRef > 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

     budgets[_organisationsRef][_organisationRef][_budgetRef] = _budget;

     if(!Strings.getExists(_budgetRef, budgetRefs[_organisationsRef][_organisationRef])) {
       budgetRefs[_organisationsRef][_organisationRef].push(_budgetRef);
     }

     emit SetRegionBudget(_organisationsRef, _organisationRef, _budgetRef, _budget);
  }

  function getNumRegionBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return budgetRefs[_organisationsRef][_organisationRef].length;
  }

  function getRegionBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _index < budgetRefs[_organisationsRef][_organisationRef].length);

    return budgetRefs[_organisationsRef][_organisationRef][_index];
  }

  function getRegionsBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (RegionBudget memory) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef];
  }

  function getRegionsBudgetRegion(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].regionRef;
  }

  function getRegionsBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].budgetLine;
  }

  function getRegionsBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.value;
  }

  function getRegionsBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint8) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.status;
  }

  function getRegionsBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.start;
  }

  function getRegionsBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.end;
  }
}
