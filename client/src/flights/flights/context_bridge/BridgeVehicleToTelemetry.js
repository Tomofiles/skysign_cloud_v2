import React, { useContext } from 'react';

import useInterval from 'use-interval';
import { Cartesian3, HeadingPitchRoll, Math, Matrix3, Matrix4, Quaternion, Transforms } from 'cesium';

import { getTelemetry } from '../../../map/MapUtils';
import { AppContext } from '../../../context/Context';

const BridgeVehicleToTelemetry = () => {
  const { vehicles, dispatchTelemetries } = useContext(AppContext);

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

  return (<></>)
}

export default BridgeVehicleToTelemetry;