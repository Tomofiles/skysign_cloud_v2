import React from 'react';

import {
  Drawer,
  Grid,
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

import Missions from './missions/Missions'

const Plans = (props) => {
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
                Plans
              </Typography>
            </Grid>
            <Grid item xs={4} />
          </Grid>
        </Box>
      </Box>
      <Missions classes={props.classes} open={props.open} />
    </Drawer>
  );
}

export default Plans;