import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

import YourMissions from './missions/YourMissions'
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';
import { EDIT_MODE } from '../context/EditMode';

const Missions = (props) => {
  const { funcMode, editMode } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);
  const [ edit, setEdit ] = useState(false);

  useEffect(() => {
    switch (funcMode) {
      case FUNC_MODE.MISSIONS:
        setOpen(true);
        break;
      default:
        setOpen(false);
    }
  }, [ funcMode ])

  useEffect(() => {
    switch (editMode) {
      case EDIT_MODE.MISSION:
        setEdit(true);
        break;
      default:
        setEdit(false);
    }
  }, [ editMode ])

  return (
    <>
      {open && (
        <div className={props.classes.func + ` ${edit ? props.classes.funcEditable : ''}`} >
          <div className={props.classes.funcPaper} >
            <Box m={4}>
              <Box style={{display: 'flex'}}>
                <Settings style={{ color: grey[50] }} fontSize="small" />
                <Typography align="left" component="div">
                  Missions
                </Typography>
              </Box>
            </Box>
            <YourMissions classes={props.classes} open={open}/>
          </div>
        </div>
      )}
    </>
  );
}

export default Missions;