import React, { useGlobal } from 'reactn';
import { Scene } from 'resium';

const SceneMode = (props) => {
  const [ mode ] = useGlobal("mapMode");

  return (
    <Scene
      mode={mode}
      morphTime={0.0}
      morphDuration={0.0} />
  );
}

export default SceneMode;