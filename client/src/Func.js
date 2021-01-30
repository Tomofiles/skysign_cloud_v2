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
    setControlsOpen(open => !open);
    setPlansOpen(false);
    setAssetsOpen(false);
  }

  const togglePlans = () => {
    setControlsOpen(false);
    setPlansOpen(open => !open);
    setAssetsOpen(false);
  }

  const toggleAssets = () => {
    setControlsOpen(false);
    setPlansOpen(false);
    setAssetsOpen(open => !open);
  }

  return (
    <div className={props.classes.root}>
      <Menu
        classes={props.classes}
        controlsOpen={controlsOpen}
        plansOpen={plansOpen}
        assetsOpen={assetsOpen}
        toggleControls={toggleControls}
        togglePlans={togglePlans}
        toggleAssets={toggleAssets} />
      <Controls classes={props.classes} open={controlsOpen} />
      <Plans classes={props.classes} open={plansOpen} />
      <Assets classes={props.classes} open={assetsOpen} />
    </div>
  );
}

export default Func;
