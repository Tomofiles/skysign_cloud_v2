import React from 'react';

import Menu from './menu/Menu'
import Plans from './plans/Plans'
import Assets from './assets/Assets'
import Missions from './missions/Missions';
import Flights from './flights/Flights';
import Reports from './reports/Reports';
import MessageNotify from './MessageNotify';

const Func = (props) => {
  return (
    <div className={props.classes.root}>
      <MessageNotify />
      <Menu classes={props.classes}/>
      <Reports classes={props.classes} />
      <Flights classes={props.classes} />
      <Plans classes={props.classes} />
      <Missions classes={props.classes} />
      <Assets classes={props.classes} />
    </div>
  );
}

export default Func;
