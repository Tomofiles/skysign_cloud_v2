import React, { useState, useEffect } from 'react';

import { getMission } from '../plans/missions/MissionUtils'
import { Mission as MissionModel } from '../plans/missions/MissionHelper'
import Waypoint from './Waypoint';
import Path from './Path';

const Mission = (props) => {
  const [ mission, setMission ] = useState(new MissionModel());

  useEffect(() => {
    getMission(props.data.missionId)
      .then(mission => {
        let newMission = Object.assign(Object.create(MissionModel.prototype), mission);
        setMission(newMission);
      });
  }, [ props.data.missionId, setMission ]);

  return (
    <>
      {mission.getWaypointsForDisplayToMap().map((waypoint, index) => (
        <Waypoint key={waypoint.id} index={index} waypoint={waypoint} />
      ))}
      {mission.getPathsForDisplayToMap().map(path => (
        <Path key={path.id} path={path} />
      ))}
    </>
  );
}
  
export default Mission;