import React, { useContext, useEffect } from 'react';

import { AppContext } from '../context/Context';
import { getMission } from '../missions/missions/MissionUtils';

const BridgeAssignmentToMission = () => {
  const { assignments, dispatchMissions } = useContext(AppContext);

  useEffect(() => {
    if (assignments.length === 0) {
      dispatchMissions({ type: 'NONE' });
      return;
    }

    let missions = [];
    assignments
      .forEach(assignment => {
        missions.push(getMission(assignment.mission_id));
      });

    Promise
      .all(missions)
      .then(data => {
        dispatchMissions({ type: 'ROWS', rows: data });
      });
  }, [ assignments, dispatchMissions ])

  return (<></>)
}

export default BridgeAssignmentToMission;