import React, { useContext, useEffect } from 'react';

import {
  Drawer,
  Grid,
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';
import PlanCalendar from './calendar/PlanCalendar';
import MyFlightplans from './flightplans/MyFlightplans';
import { AppContext } from '../context/Context';

const Plans = (props) => {
  const { dispatchPlannerMode } = useContext(AppContext);

  useEffect(() => {
    if (props.open) {
      dispatchPlannerMode({ type: 'PLANNING' });
    } else {
      dispatchPlannerMode({ type: 'NONE' });
    }
  }, [ props.open, dispatchPlannerMode ])

  return (
    <Drawer
        className={props.classes.func}
        anchor='right'
        variant="persistent"
        classes={{
          paper: props.classes.funcPaper,
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
      <PlanCalendar classes={props.classes} />
      <MyFlightplans classes={props.classes} />
    </Drawer>
  );
}

export default Plans;