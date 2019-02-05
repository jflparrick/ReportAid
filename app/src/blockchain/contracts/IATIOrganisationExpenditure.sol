pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationExpenditure.sol";
import "./Strings.sol";

contract IATIOrganisationExpenditure is OrganisationExpenditure {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private expenditureReferences;
  mapping(bytes32 => mapping(bytes32 => Expenditure)) private expenditures;

  event SetExpenditure(Expenditure _expenditure);

  function setExpenditure(Expenditure memory _expenditure) public {
    require (_expenditure.report.reportRef[0] != 0 &&
             _expenditure.report.orgRef[0] != 0 &&
             _expenditure.expenditureRef[0] != 0 &&
             _expenditure.expenditureLine[0] != 0 &&
             _expenditure.finance.status > 0 &&
             _expenditure.finance.start[0] != 0 &&
             _expenditure.finance.end[0] != 0 );

    expenditures[_expenditure.report.reportRef][_expenditure.expenditureRef] = _expenditure;

    if (!getReportExists(_expenditure.report.reportRef)) {
      reportReferences.push(_expenditure.report.reportRef);
    }

    if (!getExpenditureExists(_expenditure.report.reportRef, _expenditure.expenditureRef)) {
    expenditureReferences[_expenditure.report.reportRef].push(_expenditure.expenditureRef);
    }

    emit SetExpenditure(_expenditure);
  }

  function getReportExists(bytes32 _reportRef) public view returns (bool) {
    require (_reportRef[0] != 0);

    bool exists = false;
    if ( !(reportReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_reportRef, reportReferences);
      exists = (index != reportReferences.length);
    }
    return exists;
  }

  function getExpenditureExists(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    bool exists = false;
    if ( !(expenditureReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_expenditureRef, expenditureReferences[_reportRef]);
      exists = (index != expenditureReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumExpenditures(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return expenditureReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getExpenditureReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < expenditureReferences[_reportRef].length);

    return expenditureReferences[_reportRef][_index];
  }

  function getExpenditure(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (Expenditure memory) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef];
  }

  function getExpenditureReportingOrg(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].report.orgRef;
  }

  function getExpenditureLine(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].expenditureLine;
  }

  function getExpenditureValue(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.value;
  }

  function getExpenditureStatus(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint8) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.status;
  }

  function getExpenditureStart(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.start;
  }

  function getExpenditureEnd(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.end;
  }
}