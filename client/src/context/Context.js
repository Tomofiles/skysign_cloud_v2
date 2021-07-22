import React, { createContext, useReducer } from 'react';

import { initialEditMission, editMissionReducer } from './EditMission';
import { initialEditMode, editModeReducer } from './EditMode';
import { initialFuncMode, funcModeReducer } from './FuncMode';
import { initialMapMode, mapModeReducer } from './MapMode';
import { initialMapPosition, mapPositionReducer } from './MapPosition';
import { initialMissions, missionsReducer } from './Missions';
import { initialOperationMode, operationModeReducer } from './OperationMode';
import { initialVehicles, vehiclesReducer } from './Vehicles';
import { initialTelemetries, telemetriesReducer } from './Telemetries';
import { initialFleet, fleetReducer } from './Fleet';
import { initialAssignments, assignmentsReducer } from './Assignments';
import { initialSteps, stepsReducer } from './Steps';
import { initialTrajectories, trajectoriesReducer } from './Trajectories';

import BridgeFleetToAssignment from '../context_bridge/BridgeFleetToAssignment';
import BridgeAssignmentToMission from '../context_bridge/BridgeAssignmentToMission';
import BridgeAssignmentToVehicle from '../context_bridge/BridgeAssignmentToVehicle';
import BridgeVehicleToStep from '../context_bridge/BridgeVehicleToStep';
import BridgeVehicleToTelemetry from '../context_bridge/BridgeVehicleToTelemetry';
import BridgeVehicleToReportsTrajectory from '../context_bridge/BridgeVehicleToReportsTrajectory';
import BridgeVehicleToOperationTrajectory from '../context_bridge/BridgeVehicleToOperationTrajectory';

export const AppContext = createContext();

const AppContextProvider = ({children}) => {
  const [ fleet, dispatchFleet ] = useReducer(fleetReducer, initialFleet);
  const [ assignments, dispatchAssignments ] = useReducer(assignmentsReducer, initialAssignments);
  const [ vehicles, dispatchVehicles ] = useReducer(vehiclesReducer, initialVehicles);
  const [ missions, dispatchMissions ] = useReducer(missionsReducer, initialMissions);
  const [ telemetries, dispatchTelemetries ] = useReducer(telemetriesReducer, initialTelemetries);
  const [ trajectories, dispatchTrajectories ] = useReducer(trajectoriesReducer, initialTrajectories);
  const [ steps, dispatchSteps ] = useReducer(stepsReducer, initialSteps);
  const [ editMission, dispatchEditMission ] = useReducer(editMissionReducer, initialEditMission);
  const [ editMode, dispatchEditMode ] = useReducer(editModeReducer, initialEditMode);
  const [ operationMode, dispatchOperationMode ] = useReducer(operationModeReducer, initialOperationMode);
  const [ mapMode, dispatchMapMode ] = useReducer(mapModeReducer, initialMapMode);
  const [ funcMode, dispatchFuncMode ] = useReducer(funcModeReducer, initialFuncMode);
  const [ mapPosition, dispatchMapPosition ] = useReducer(mapPositionReducer, initialMapPosition);

  return (
    <AppContext.Provider
      value={{
        fleet,
        dispatchFleet,
        assignments,
        dispatchAssignments,
        vehicles,
        dispatchVehicles,
        missions,
        dispatchMissions,
        telemetries,
        dispatchTelemetries,
        trajectories,
        dispatchTrajectories,
        steps,
        dispatchSteps,
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
      <BridgeFleetToAssignment />
      <BridgeAssignmentToMission />
      <BridgeAssignmentToVehicle />
      <BridgeVehicleToStep />
      <BridgeVehicleToTelemetry />
      <BridgeVehicleToOperationTrajectory />
      <BridgeVehicleToReportsTrajectory />
    </AppContext.Provider>
  )
}

export default AppContextProvider;