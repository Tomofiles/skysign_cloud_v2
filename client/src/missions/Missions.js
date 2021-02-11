import React, { useContext, useEffect, useState } from 'react';

import {
  Grid,
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

import MyMissions from './missions/MyMissions'
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';

const Missions = (props) => {
  const { funcMode } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);

  useEffect(() => {
    switch (funcMode) {
      case FUNC_MODE.MISSIONS:
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
                      Missions
                    </Typography>
                  </Grid>
                  <Grid item xs={4} />
                </Grid>
              </Box>
            </Box>
            <MyMissions classes={props.classes} open={open}/>
          </div>
        </div>
      )}
    </>
  );
}

export default Missions;