import React, { useState } from 'react';

import Menu from './menu/Menu'
import Plans from './plans/Plans'
import Assets from './assets/Assets'
import Controls from './controls/Controls';
import Missions from './missions/Missions';

const Func = (props) => {
  const [controlsOpen, setControlsOpen] = useState(false);
  const [plansOpen, setPlansOpen] = useState(false);
  const [missionsOpen, setMissionsOpen] = useState(false);
  const [assetsOpen, setAssetsOpen] = useState(false);

  const toggleControls = () => {
    setControlsOpen(open => !open);
    setPlansOpen(false);
    setMissionsOpen(false);
    setAssetsOpen(false);
  }

  const togglePlans = () => {
    setControlsOpen(false);
    setPlansOpen(open => !open);
    setMissionsOpen(false);
    setAssetsOpen(false);
  }

  const toggleMissions = () => {
    setControlsOpen(false);
    setPlansOpen(false);
    setMissionsOpen(open => !open);
    setAssetsOpen(false);
  }

  const toggleAssets = () => {
    setControlsOpen(false);
    setPlansOpen(false);
    setMissionsOpen(false);
    setAssetsOpen(open => !open);
  }

  return (
    <div className={props.classes.root}>
      <Menu
        classes={props.classes}
        controlsOpen={controlsOpen}
        plansOpen={plansOpen}
        missionsOpen={missionsOpen}
        assetsOpen={assetsOpen}
        toggleControls={toggleControls}
        togglePlans={togglePlans}
        toggleMissions={toggleMissions}
        toggleAssets={toggleAssets} />
      <Controls classes={props.classes} open={controlsOpen} />
      <Plans classes={props.classes} open={plansOpen} />
      <Missions classes={props.classes} open={missionsOpen} />
      <Assets classes={props.classes} open={assetsOpen} />
    </div>
  );
}

export default Func;
