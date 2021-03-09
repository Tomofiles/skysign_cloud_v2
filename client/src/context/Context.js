import React, { createContext, useEffect, useReducer } from 'react';

import { initialEditMission, editMissionReducer } from './EditMission';
import { initialEditMode, editModeReducer } from './EditMode';
import { initialFuncMode, funcModeReducer } from './FuncMode';
import { initialMapMode, mapModeReducer } from './MapMode';
import { initialMapPosition, mapPositionReducer } from './MapPosition';
import { initialMissions, missionsReducer } from './Missions';
import { initialOperationMode, operationModeReducer } from './OperationMode';
import { initialStagingRows, stagingRowsReducer } from './StagingRows';
import { initialVehicles, vehiclesReducer } from './Vehicles';
import { initialTelemetries, telemetriesReducer } from './Telemetries';
import { initialFlight, flightReducer } from './Flight';
import { initialAssignments, assignmentsReducer } from './Assignments';
import { initialFlightplan, flightplanReducer } from './Flightplan';
import { initialSteps, stepsReducer } from './Steps';

import { Cartesian3, HeadingPitchRoll, Quaternion, Math, Transforms, Matrix4, Matrix3 } from 'cesium';
import useInterval from 'use-interval';

import { getMission } from '../missions/missions/MissionUtils'
import { getVehicle } from '../assets/vehicles/VehicleUtils';
import { getTelemetry } from '../map/MapUtils';
import { getAssignments, getFlightplan } from '../plans/flightplans/FlightplansUtils';
import { getFlight } from '../flights/flights/FlightUtils';

export const AppContext = createContext();

const AppContextProvider = ({children}) => {
  const [ flight, dispatchFlight ] = useReducer(flightReducer, initialFlight);
  const [ flightplan, dispatchFlightplan ] = useReducer(flightplanReducer, initialFlightplan);
  const [ assignments, dispatchAssignments ] = useReducer(assignmentsReducer, initialAssignments);
  const [ vehicles, dispatchVehicles ] = useReducer(vehiclesReducer, initialVehicles);
  const [ missions, dispatchMissions ] = useReducer(missionsReducer, initialMissions);
  const [ telemetries, dispatchTelemetries ] = useReducer(telemetriesReducer, initialTelemetries);
  const [ steps, dispatchSteps ] = useReducer(stepsReducer, initialSteps);
  const [ stagingRows, dispatchStagingRows ] = useReducer(stagingRowsReducer, initialStagingRows);
  const [ editMission, dispatchEditMission ] = useReducer(editMissionReducer, initialEditMission);
  const [ editMode, dispatchEditMode ] = useReducer(editModeReducer, initialEditMode);
  const [ operationMode, dispatchOperationMode ] = useReducer(operationModeReducer, initialOperationMode);
  const [ mapMode, dispatchMapMode ] = useReducer(mapModeReducer, initialMapMode);
  const [ funcMode, dispatchFuncMode ] = useReducer(funcModeReducer, initialFuncMode);
  const [ mapPosition, dispatchMapPosition ] = useReducer(mapPositionReducer, initialMapPosition);

  useEffect(() => {
    if (flight) {
      getFlight(flight)
      .then(data => {
        getFlightplan(data.flightplan_id)
          .then(data => {
            dispatchFlightplan({ type: 'DATA', data: data });
          })
        getAssignments(data.flightplan_id)
          .then(data => {
            dispatchAssignments({ type: 'ROWS', rows: data.assignments });
          })
      })
    } else {
      dispatchFlightplan({ type: 'NONE' });
      dispatchAssignments({ type: 'NONE' });
    }
  }, [ flight, dispatchFlightplan, dispatchAssignments ])

  useEffect(() => {
    if (assignments.length === 0) {
      dispatchVehicles({ type: 'NONE' });
      return;
    }

    let vehicles = [];
    assignments
      .forEach(assignment => {
        vehicles.push(getVehicle(assignment.vehicle_id));
      });

    Promise
      .all(vehicles)
      .then(data => {
        dispatchVehicles({ type: 'ROWS', rows: data });
      });
  }, [ assignments, dispatchVehicles ])

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

  useInterval(() => {
    if (vehicles.length === 0) {
      dispatchTelemetries({ type: 'NONE' });
      return;
    }

    let telemetries = [];
    vehicles
      .forEach(vehicle => {
        telemetries.push(getTelemetry(vehicle.commId));
      });

    Promise
      .all(telemetries)
      .then(data => {
        for (let telemetry of data) {
          let pos = Cartesian3.fromDegrees(
            telemetry.telemetry.longitude,
            telemetry.telemetry.latitude,
            telemetry.telemetry.altitude);
          let mtx4 = Transforms.eastNorthUpToFixedFrame(pos);
          let mtx3 = Matrix4.getMatrix3(mtx4, new Matrix3());
          let base = Quaternion.fromRotationMatrix(mtx3);
          let quatlocal = new Quaternion(
            telemetry.telemetry.orientationY,
            telemetry.telemetry.orientationX,
            -telemetry.telemetry.orientationZ,
            telemetry.telemetry.orientationW);
          let quat90 = Quaternion.fromAxisAngle(
            new Cartesian3(0, 0, 1),
            Math.toRadians(90)
          );
          let quatlocalaft = Quaternion.multiply(quatlocal, quat90, new Quaternion());
          let quat = Quaternion.multiply(base, quatlocalaft, new Quaternion());
          let hpr = HeadingPitchRoll.fromQuaternion(quat);
          telemetry.telemetry.heading = hpr.heading;
          telemetry.telemetry.pitch = hpr.pitch;
          telemetry.telemetry.roll = hpr.roll;
        }
        dispatchTelemetries({ type: 'ROWS', rows: data });
      });
  },
  1000);

  useEffect(() => {
    if (vehicles.length === 0) {
      dispatchSteps({ type: 'NONE' });
      return;
    }

    let communicationIds = [];
    vehicles
      .forEach(vehicle => {
        communicationIds.push(vehicle.commId);
      });

    dispatchSteps({ type: 'INIT', ids: communicationIds });
  }, [ vehicles, dispatchSteps ])

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