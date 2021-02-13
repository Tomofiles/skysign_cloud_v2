import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';
import YourFlightplans from './flightplans/YourFlightplans';
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
        <div className={props.classes.func}>
          <div className={props.classes.funcPaper} >
            <Box m={4}>
              <Box style={{display: 'flex'}}>
                <Settings style={{ color: grey[50] }} fontSize="small" />
                <Typography align="left" component="div">
                  Plans
                </Typography>
              </Box>
            </Box>
            <YourFlightplans classes={props.classes} open={open}/>
          </div>
        </div>
      )}
    </>
  );
}

export default Plans;