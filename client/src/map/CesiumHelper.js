import { Cartesian3, Transforms, Matrix3, Matrix4, Quaternion, HeadingPitchRoll, Math as CesiumMath } from "cesium";

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

  // ローカルクォータニオンをオイラー角に変換
  let hpr = HeadingPitchRoll.fromQuaternion(quatlocal);

  let description =   '<table class="cesium-infoBox-defaultTable"><tbody>' +
                      '<tr><th>機体ID</th><td>' + vehicleID + '</td></tr>' +
                      '<tr><th>飛行モード</th><td>' + telemetry.flightMode + '</td></tr>' +
                      '<tr><th>緯度(°)</th><td>' + dispFloor(telemetry.latitude, 10) + '</td></tr>' +
                      '<tr><th>経度(°)</th><td>' + dispFloor(telemetry.longitude, 10) + '</td></tr>' +
                      '<tr><th>海抜高度(m)</th><td>' + dispFloor(telemetry.altitude, 10) + '</td></tr>' +
                      '<tr><th>相対高度(m)</th><td>' + dispFloor(telemetry.relativeHeight, 10) + '</td></tr>' +
                      '<tr><th>対地速度(m/s)</th><td>' + dispFloor(telemetry.speed, 10) + '</td></tr>' +
                      '<tr><th>ヘディング(°)</th><td>' + dispFloor(CesiumMath.toDegrees(hpr.heading), 10) + '</td></tr>' +
                      '<tr><th>ピッチ(°)</th><td>' + dispFloor(CesiumMath.toDegrees(hpr.pitch), 10) + '</td></tr>' +
                      '<tr><th>ロール(°)</th><td>' + dispFloor(CesiumMath.toDegrees(hpr.roll), 10) + '</td></tr>' +
                      '</tbody></table>';

  let entityID = "drone_" + vehicleID;

  let data = {
    id: entityID,
    name: telemetry.name,
    description: description,
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

const dispFloor = (num, digit) => {
  return Math.floor(num * Math.pow(10, digit) ) / Math.pow(10, digit);
}