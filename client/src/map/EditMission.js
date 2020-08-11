import React, { useGlobal } from 'reactn';

import EditWaypoint from './EditWaypoint';
import EditPath from './EditPath';

const EditMission = () => {
  const mission = useGlobal("editMission")[0];

  return (
    <>
      {mission.getWaypointsForDisplayToMap().map((waypoint, index) => (
        <EditWaypoint key={waypoint.id} index={index} waypoint={waypoint} />
      ))}
      {mission.getPathsForDisplayToMap().map(path => (
        <EditPath key={path.id} path={path} />
      ))}
    </>
  );
}

export default EditMission;