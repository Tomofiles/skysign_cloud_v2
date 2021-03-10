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
import { initialFlight, flightReducer } from './Flight';
import { initialAssignments, assignmentsReducer } from './Assignments';
import { initialFlightplan, flightplanReducer } from './Flightplan';
import { initialSteps, stepsReducer } from './Steps';

import BridgeAssignmentToMission from '../flights/flights/context_bridge/BridgeAssignmentToMission';
import BridgeAssignmentToVehicle from '../flights/flights/context_bridge/BridgeAssignmentToVehicle';
import BridgeFlightIDToFlightplanAndAssignment from '../flights/flights/context_bridge/BridgeFlightIDToFlightplanAndAssignment';
import BridgeVehicleToStep from '../flights/flights/context_bridge/BridgeVehicleToStep';
import BridgeVehicleToTelemetry from '../flights/flights/context_bridge/BridgeVehicleToTelemetry';

export const AppContext = createContext();

const AppContextProvider = ({children}) => {
  const [ flight, dispatchFlight ] = useReducer(flightReducer, initialFlight);
  const [ flightplan, dispatchFlightplan ] = useReducer(flightplanReducer, initialFlightplan);
  const [ assignments, dispatchAssignments ] = useReducer(assignmentsReducer, initialAssignments);
  const [ vehicles, dispatchVehicles ] = useReducer(vehiclesReducer, initialVehicles);
  const [ missions, dispatchMissions ] = useReducer(missionsReducer, initialMissions);
  const [ telemetries, dispatchTelemetries ] = useReducer(telemetriesReducer, initialTelemetries);
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
        flight,
        dispatchFlight,
        flightplan,
        dispatchFlightplan,
        assignments,
        dispatchAssignments,
        vehicles,
        dispatchVehicles,
        missions,
        dispatchMissions,
        telemetries,
        dispatchTelemetries,
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
      <BridgeAssignmentToMission />
      <BridgeAssignmentToVehicle />
      <BridgeFlightIDToFlightplanAndAssignment />
      <BridgeVehicleToStep />
      <BridgeVehicleToTelemetry />
    </AppContext.Provider>
  )
}

export default AppContextProvider;