import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

import YourVehicles from './vehicles/YourVehicles'
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';

const Assets = (props) => {
  const { funcMode } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);

  useEffect(() => {
    switch (funcMode) {
      case FUNC_MODE.ASSETS:
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
            <Box p={4}>
              <Box style={{display: 'flex'}}>
                <Settings style={{ color: grey[50] }} fontSize="small" />
                <Typography align="left" component="div">
                  Assets
                </Typography>
              </Box>
            </Box>
            <YourVehicles classes={props.classes} open={open} />
          </div>
        </div>
      )}
    </>
  );
}

export default Assets;