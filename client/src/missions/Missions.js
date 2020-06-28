import React from 'react';

import {
  Drawer
} from '@material-ui/core';
const list = (classes) => (
  <div>
    <span>Missions</span>
  </div>
);

const Missions = (props) => {
  return (
    <Drawer
      className={props.classes.assets}
      anchor='right'
      variant="persistent"
      classes={{
        paper: props.classes.assetsPaper,
      }}
      open={props.open} >
      {list(props.classes)}
    </Drawer>
  );
}

export default Missions;