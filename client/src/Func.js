import React, { useState } from 'react';

import Menu from './menu/Menu'
import Plans from './plans/Plans'
import Assets from './assets/Assets'
import Controls from './controls/Controls';
import Actual from './actual/Actual';

const Func = (props) => {
  const [controlsOpen, setControlsOpen] = useState(false);
  const [plansOpen, setPlansOpen] = useState(false);
  const [actualOpen, setActualOpen] = useState(false);
  const [assetsOpen, setAssetsOpen] = useState(false);

  const toggleControls = () => {
    setControlsOpen(open => !open);
    setPlansOpen(false);
    setActualOpen(false);
    setAssetsOpen(false);
  }

  const togglePlans = () => {
    setControlsOpen(false);
    setPlansOpen(open => !open);
    setActualOpen(false);
    setAssetsOpen(false);
  }

  const toggleActual = () => {
    setControlsOpen(false);
    setPlansOpen(false);
    setActualOpen(open => !open);
    setAssetsOpen(false);
  }

  const toggleAssets = () => {
    setControlsOpen(false);
    setPlansOpen(false);
    setActualOpen(false);
    setAssetsOpen(open => !open);
  }

  return (
    <div className={props.classes.root}>
      <Menu
        classes={props.classes}
        controlsOpen={controlsOpen}
        plansOpen={plansOpen}
        actualOpen={actualOpen}
        assetsOpen={assetsOpen}
        toggleControls={toggleControls}
        togglePlans={togglePlans}
        toggleActual={toggleActual}
        toggleAssets={toggleAssets} />
      <Controls classes={props.classes} open={controlsOpen} />
      <Plans classes={props.classes} open={plansOpen} />
      <Actual classes={props.classes} open={actualOpen} />
      <Assets classes={props.classes} open={assetsOpen} />
    </div>
  );
}

export default Func;
