import { Mission } from "../plans/missions/MissionHelper";

export const initialEditMission = new Mission();

export const editMissionReducer = (state, action) => {
  let newMission = Object.assign(Object.create(Mission.prototype), state);
  switch (action.type) {
    case 'CHANGE_NAME':
      newMission.nameMission(action.name);
      return newMission;
    case 'CHANGE_RELATIVE_HEIGHT':
      const changeRelativeHeight = async () => {
        await newMission.changeRelativeHeight(action.index, action.height);
        return newMission;
      }
      return changeRelativeHeight();
    case 'CHANGE_SPEED':
      const changeSpeed = async () => {
        await newMission.changeSpeed(action.index, action.speed);
        return newMission;
      };
      return changeSpeed();
    case 'ADD_WAYPOINT':
      const addWaypoint = async () => {
        await newMission.addWaypoint(action.latitude, action.longitude);
        return newMission;
      };
      return addWaypoint();
    case 'REMOVE_WAYPOINT':
      newMission.removeWaypoint(action.index);
      return newMission;
    case 'CHANGE_POSITION':
      newMission.changePosition(action.index, action.latitude, action.longitude);
      return newMission;
    case 'CLEAR':
      return new Mission();
    default:
      return new Mission();
  }
}
  