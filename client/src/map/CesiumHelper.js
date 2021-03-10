import { Cartesian3, Transforms, Matrix3, Matrix4, Quaternion, Math as CesiumMath } from "cesium";

export const convertDroneData = (vehicleID, telemetry) => {
  // 地球固定座標での回転を計算
  let pos = Cartesian3.fromDegrees(
    telemetry.longitude,
    telemetry.latitude,
    telemetry.altitude);
  let mtx4 = Transforms.eastNorthUpToFixedFrame(pos);
  let mtx3 = Matrix4.getMatrix3(mtx4, new Matrix3());
  let base = Quaternion.fromRotationMatrix(mtx3);
  // ローカル座標での回転を計算（NED→ENU）
  let quatlocal = new Quaternion(
    telemetry.orientationY,
    telemetry.orientationX,
    -telemetry.orientationZ,
    telemetry.orientationW);
  let quat90 = Quaternion.fromAxisAngle(
    new Cartesian3(0, 0, 1),
    CesiumMath.toRadians(90)
  );
  let quatlocalaft = Quaternion.multiply(quatlocal, quat90, new Quaternion());
  // 回転を掛け合わせる
  let quat = Quaternion.multiply(base, quatlocalaft, new Quaternion());

  let entityID = "drone_" + vehicleID;

  let data = {
    id: entityID,
    position: Cartesian3.fromDegrees(
      telemetry.longitude,
      telemetry.latitude,
      telemetry.altitude
    ),
    orientation: new Quaternion(
      quat.x,
      quat.y,
      quat.z,
      quat.w
    ),
    armed: telemetry.armed,
    properties: {
      vehicleID: vehicleID
    }
  };

  return data;
}
