import React, { useGlobal } from 'reactn';

import Waypoint from './Waypoint';
import Path from './Path';

const EditMission = () => {
  const mission = useGlobal("editMission")[0];

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

export default EditMission;