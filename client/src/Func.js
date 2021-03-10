import React from 'react';

import Menu from './menu/Menu'
import Plans from './plans/Plans'
import Assets from './assets/Assets'
import Missions from './missions/Missions';
import Flights from './flights/Flights';

const Func = (props) => {
  return (
    <div className={props.classes.root}>
      <Menu classes={props.classes}/>
      <Flights classes={props.classes} />
      <Plans classes={props.classes} />
      <Missions classes={props.classes} />
      <Assets classes={props.classes} />
    </div>
  );
}

export default Func;
