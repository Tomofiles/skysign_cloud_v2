import React, { useContext, useEffect, useState } from 'react';

import useInterval from 'use-interval';
import {
  HeadingPitchRoll,
  Math,
  Quaternion,
} from 'cesium';

import { getTelemetry } from '../map/MapUtils';
import { AppContext } from '../context/Context';
import { OPERATION_MODE } from '../context/OperationMode';

const BridgeVehicleToTelemetry = () => {
  const { vehicles, operationMode, dispatchTelemetries } = useContext(AppContext);
  const [ isRender, setIsRender ] = useState(false);

  useEffect(() => {
    if (operationMode === OPERATION_MODE.OPERATION) {
      setIsRender(true);
    } else {
      setIsRender(false);
      dispatchTelemetries({ type: 'NONE' });
    }
  }, [ vehicles, operationMode, dispatchTelemetries ])

  useInterval(() => {
    if (!isRender) {
      return;
    }
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
          let quatlocal = new Quaternion(
            telemetry.telemetry.orientationY,
            telemetry.telemetry.orientationX,
            -telemetry.telemetry.orientationZ,
            telemetry.telemetry.orientationW);
          let hpr = HeadingPitchRoll.fromQuaternion(quatlocal);
          telemetry.telemetry.heading = Math.toDegrees(hpr.heading);
          telemetry.telemetry.pitch = Math.toDegrees(hpr.pitch);
          telemetry.telemetry.roll = Math.toDegrees(hpr.roll);
        }
        dispatchTelemetries({ type: 'ROWS', rows: data });
      });
  },
  1000);

  return (<></>)
}

export default BridgeVehicleToTelemetry;