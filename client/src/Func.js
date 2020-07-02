import React, { useState } from 'react';

import Menu from './menu/Menu'
import Missions from './missions/Missions'
import Assets from './assets/Assets'
import Controls from './controls/Controls';

const Func = (props) => {
  const [controlsOpen, setControlsOpen] = useState(false);
  const [missionsOpen, setMissionsOpen] = useState(false);
  const [assetsOpen, setAssetsOpen] = useState(false);

  const toggleControls = () => {
    if (missionsOpen) {
      setMissionsOpen(false);
    }
    if (assetsOpen) {
      setAssetsOpen(false);
    }
    setControlsOpen(!controlsOpen);
  }

  const toggleMissions = () => {
    if (assetsOpen) {
      setAssetsOpen(false);
    }
    if (controlsOpen) {
      setControlsOpen(false);
    }
    setMissionsOpen(!missionsOpen);
  }

  const toggleAssets = () => {
    if (missionsOpen) {
      setMissionsOpen(false);
    }
    if (controlsOpen) {
      setControlsOpen(false);
    }
    setAssetsOpen(!assetsOpen);
  }

  return (
    <div className={props.classes.root}>
      <Menu
        classes={props.classes}
        controlsOpen={controlsOpen}
        missionsOpen={missionsOpen}
        assetsOpen={assetsOpen}
        toggleControls={toggleControls}
        toggleMissions={toggleMissions}
        toggleAssets={toggleAssets} />
      <Controls classes={props.classes} open={controlsOpen} />
      <Missions classes={props.classes} open={missionsOpen} />
      <Assets classes={props.classes} open={assetsOpen} />
    </div>
  );
}

export default Func;
