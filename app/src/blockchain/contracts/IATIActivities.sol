pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activities.sol";
import "./Strings.sol";

contract IATIActivities is Activities {

  bytes32[] activitiesRefs;
  mapping(bytes32 =>  OrgActivities) private activities;

  event SetActivities(OrgActivities _activities);

  function setActivities(OrgActivities memory _activities) public {
    require (_activities.activitiesRef[0] != 0  &&
             _activities.version[0] != 0 &&
             _activities.generatedTime[0] != 0);

    activities[_activities.activitiesRef] = _activities;

    if (!Strings.getExists(_activities.activitiesRef, activitiesRefs)) {
      activitiesRefs.push(_activities.activitiesRef);
    }

    emit SetActivities(_activities);
  }

  function getNumActivities() public view returns (uint256) {
    return activitiesRefs.length;
  }

  function getActivitiesReference(uint256 _index) public view returns (bytes32) {
    require (_index < activitiesRefs.length);

    return activitiesRefs[_index];
  }

  function getActivities(bytes32 _activitiesRef) public view returns (OrgActivities memory) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef];
  }

  function getVersion(bytes32 _activitiesRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].version;
  }

  function getGeneratedTime(bytes32 _activitiesRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].generatedTime;
  }

  function getLinkedData(bytes32 _activitiesRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].linkedData;
  }
}
