import React, { useCallback, useContext, useEffect, useState } from 'react';

import { Cartesian3 } from 'cesium';
import { CameraFlyTo } from 'resium';
import { v4 as uuidv4 } from 'uuid';

import { AppContext } from '../context/Context';

const CameraCurrentPosition = () => {
  const { mapPosition, dispatchMapPosition } = useContext(AppContext);
  const [ position, setPosition ] = useState({ cartesian3: undefined });

  const getCurrentPosition = useCallback(() => {
    navigator.geolocation.getCurrentPosition(p => {
      const { latitude, longitude } = p.coords;
      dispatchMapPosition({
        type: 'CURRENT',
        longitude: longitude,
        latitude: latitude,
        height: 2000,
      })
    });
  }, [ dispatchMapPosition ]);

  useEffect(() => {
    getCurrentPosition();
  }, [ getCurrentPosition ])

  useEffect(() => {
    setPosition({ cartesian3: undefined });
    const pos = Cartesian3.fromDegrees(
      mapPosition.longitude,
      mapPosition.latitude,
      mapPosition.height);
    setPosition({ id: uuidv4(), cartesian3: pos });
  }, [ mapPosition ])

  return (
    <>
      {position.cartesian3 && (
        <CameraFlyTo key={position.id} duration={0} destination={position.cartesian3} once/>
      )}
    </>
  );
}
export default CameraCurrentPosition;