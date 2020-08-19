import React from 'react';

import {
  Drawer,
  Box,
  Typography,
  Grid
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';
import FleetControl from './fleetcontrol/FleetControl';
import Staging from './staging/Staging';

const Controls = (props) => {
  return (
    <Drawer
      className={props.classes.assets}
      anchor='right'
      variant="persistent"
      classes={{
        paper: props.classes.assetsPaper,
      }}
      open={props.open} >
      <Box m={2} alignContent="center">
        <Box >
          <Grid container>
            <Grid item xs={4} />
            <Grid item xs={1}>
              <Settings style={{ color: grey[50] }} fontSize="small" />
            </Grid>
            <Grid item xs={3}>
              <Typography align="left" component="div">
                Controls
              </Typography>
            </Grid>
            <Grid item xs={4} />
          </Grid>
        </Box>
      </Box>
      <Staging classes={props.classes} open={props.open} />
      <FleetControl classes={props.classes} />
    </Drawer>
  );
}

export default Controls;