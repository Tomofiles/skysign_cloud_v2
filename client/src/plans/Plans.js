import React, { useContext, useEffect, useState } from 'react';

import {
  Grid,
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';
import MyFlightplans from './flightplans/MyFlightplans';
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';

const Plans = (props) => {
  const { funcMode } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);

  useEffect(() => {
    switch (funcMode) {
      case FUNC_MODE.PLANS:
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
                      Plans
                    </Typography>
                  </Grid>
                  <Grid item xs={4} />
                </Grid>
              </Box>
            </Box>
            <MyFlightplans classes={props.classes} />
          </div>
        </div>
      )}
    </>
  );
}

export default Plans;