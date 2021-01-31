import React, { useState } from 'react';

import Menu from './menu/Menu'
import Plans from './plans/Plans'
import Assets from './assets/Assets'
import Controls from './controls/Controls';

const Func = (props) => {
  const [controlsOpen, setControlsOpen] = useState(false);
  const [plansOpen, setPlansOpen] = useState(false);
  const [assetsOpen, setAssetsOpen] = useState(false);

  const toggleControls = () => {
    if (plansOpen) {
      setPlansOpen(false);
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
    setPlansOpen(!plansOpen);
  }

  const toggleAssets = () => {
    if (plansOpen) {
      setPlansOpen(false);
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
        plansOpen={plansOpen}
        assetsOpen={assetsOpen}
        toggleControls={toggleControls}
        toggleMissions={toggleMissions}
        toggleAssets={toggleAssets} />
      <Controls classes={props.classes} open={controlsOpen} />
      <Plans classes={props.classes} open={plansOpen} />
      <Assets classes={props.classes} open={assetsOpen} />
    </div>
  );
}

export default Func;
