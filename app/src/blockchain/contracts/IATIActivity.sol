pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activity.sol";
import "./Strings.sol";

contract IATIActivity is Activity {

  bytes32[] activitiesRefs;
  mapping(bytes32 => bytes32[]) private activityRefs;
  mapping(bytes32 =>  mapping(bytes32 => OrgActivity)) private activities;

  event SetActivity(bytes32 _activitiesRef, OrgActivity _activity);

  function setActivity(bytes32 _activitiesRef, OrgActivity memory _activity) public {
    require (_activitiesRef[0] != 0 &&
             _activity.activityRef[0] != 0 &&
             bytes(_activity.identifier).length > 0 &&
             _activity.reportingOrg.orgRef[0] != 0 &&
             _activity.reportingOrg.orgType > 0 &&
             bytes(_activity.title).length > 0 &&
             bytes(_activity.description).length > 0 &&
             _activity.lastUpdated[0] != 0 &&
             _activity.lang[0]  != 0 &&
             _activity.currency[0] != 0 &&
             _activity.hierarchy > uint8(Hierarchy.NONE) &&
             _activity.hierarchy < uint8(Hierarchy.MAX) &&
             _activity.budgetNotProvided >= uint8(BudgetNotProvided.NONE) &&
             _activity.budgetNotProvided < uint8(BudgetNotProvided.MAX) &&
             _activity.status > uint8(Status.NONE) &&
             _activity.status < uint8(Status.MAX) &&
             _activity.date[0] != 0 &&
             _activity.scope >= uint8(Scope.NONE) &&
             _activity.scope < uint8(Scope.MAX));

    activities[_activitiesRef][_activity.activityRef] = _activity;

    if (!Strings.getExists(_activitiesRef, activitiesRefs)) {
      activitiesRefs.push(_activitiesRef);
    }

    if(!Strings.getExists(_activity.activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(_activity.activityRef);
    }

    emit SetActivity(_activitiesRef, _activity);
  }

  function getNumActivities(bytes32 _activitiesRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0);

    return activityRefs[_activitiesRef].length;
  }

  function getActivityReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _index < activityRefs[_activitiesRef].length);

    return activityRefs[_activitiesRef][_index];
  }

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (OrgActivity memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

    return activities[_activitiesRef][_activityRef];
  }

  function getActivityIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].identifier;
  }

  function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (ReportingOrg memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].reportingOrg;
  }

  function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].title;
  }

  function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].description;
  }

  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].lastUpdated;
  }

  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].lang;
  }

  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].currency;
  }

  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bool) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].humanitarian;
  }

  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].hierarchy;
  }

  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].linkedData;
  }

  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].budgetNotProvided;
  }

  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].status;
  }

  function getDate(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].date;
  }

  function getScope(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].scope;
  }
}