import React, { useContext, useEffect, useState } from 'react';

import {
  Box,
  Typography,
  Grid
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';
import FleetControl from './fleetcontrol/FleetControl';
import Staging from './staging/Staging';
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';

const Controls = (props) => {
  const { funcMode } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);

  useEffect(() => {
    switch (funcMode) {
      case FUNC_MODE.CONTROLS:
        setOpen(true);
        break;
      default:
        setOpen(false);
    }
  }, [ funcMode ])

  return (
    <>
      {open && (
        <div className={props.classes.func} >
          <div className={props.classes.funcPaper} >
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
            <Staging classes={props.classes} open={open} />
            <FleetControl classes={props.classes} />
          </div>
        </div>
      )}
    </>
  );
}

export default Controls;