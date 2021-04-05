import React from 'react';

import Waypoint from './Waypoint';
import Path from './Path';
import { getPathsForDisplayToMap, getWaypointsForDisplayToMap } from '../missions/missions/MissionHelper';

const Mission = (props) => {
  return (
    <>
      {getWaypointsForDisplayToMap(props.mission).map((waypoint, index) => (
        <Waypoint key={waypoint.id} index={index} waypoint={waypoint} />
      ))}
      {getPathsForDisplayToMap(props.mission).map(path => (
        <Path key={path.id} path={path} />
      ))}
    </>
  );
}
  
export default Mission;