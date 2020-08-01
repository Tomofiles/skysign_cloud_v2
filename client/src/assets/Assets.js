import React from 'react';

import {
  Drawer,
  Grid,
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

import Vehicles from './vehicles/Vehicles'

const Assets = (props) => {
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
                Assets
              </Typography>
            </Grid>
            <Grid item xs={4} />
          </Grid>
        </Box>
      </Box>
      <Vehicles classes={props.classes} open={props.open} />
    </Drawer>
  );
}

export default Assets;