import React, { createContext, useEffect, useReducer } from 'react';
import { assignmentsReducer, initialAssignments } from './Assignments';

import { initialEditMission, editMissionReducer } from './EditMission';
import { initialEditMode, editModeReducer } from './EditMode';
import { funcModeReducer, initialFuncMode } from './FuncMode';
import { initialMapMode, mapModeReducer } from './MapMode';
import { initialMapPosition, mapPositionReducer } from './MapPosition';
import { initialMissions, missionsReducer } from './Missions';
import { initialOperationMode, operationModeReducer } from './OperationMode';
import { initialStagingRows, stagingRowsReducer } from './StagingRows';

import { getMission } from '../missions/missions/MissionUtils'

export const AppContext = createContext();

const AppContextProvider = ({children}) => {
  const [ assignments, dispatchAssignments ] = useReducer(assignmentsReducer, initialAssignments);
  const [ missions, dispatchMissions ] = useReducer(missionsReducer, initialMissions);
  const [ stagingRows, dispatchStagingRows ] = useReducer(stagingRowsReducer, initialStagingRows);
  const [ editMission, dispatchEditMission ] = useReducer(editMissionReducer, initialEditMission);
  const [ editMode, dispatchEditMode ] = useReducer(editModeReducer, initialEditMode);
  const [ operationMode, dispatchOperationMode ] = useReducer(operationModeReducer, initialOperationMode);
  const [ mapMode, dispatchMapMode ] = useReducer(mapModeReducer, initialMapMode);
  const [ funcMode, dispatchFuncMode ] = useReducer(funcModeReducer, initialFuncMode);
  const [ mapPosition, dispatchMapPosition ] = useReducer(mapPositionReducer, initialMapPosition);

  useEffect(() => {
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

  return (
    <AppContext.Provider
      value={{
        assignments,
        dispatchAssignments,
        missions,
        dispatchMissions,
        stagingRows,
        dispatchStagingRows,
        editMission,
        dispatchEditMission,
        editMode,
        dispatchEditMode,
        operationMode,
        dispatchOperationMode,
        mapMode,
        dispatchMapMode,
        funcMode,
        dispatchFuncMode,
        mapPosition,
        dispatchMapPosition,
        }}>
      {children}
    </AppContext.Provider>
  )
}

export default AppContextProvider;