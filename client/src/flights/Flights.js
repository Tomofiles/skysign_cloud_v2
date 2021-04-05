import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

import YourFlights from './flights/YourFlights'
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';
import { OPERATION_MODE } from '../context/OperationMode';
import FlightOperationSlider from './flights/FlightOperationSlider';

const Flights = (props) => {
  const { funcMode, operationMode } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);
  const [ operation, setOperation ] = useState(false);

  useEffect(() => {
    switch (funcMode) {
      case FUNC_MODE.FLIGHTS:
        setOpen(true);
        break;
      default:
        setOpen(false);
    }
  }, [ funcMode ])

  useEffect(() => {
    switch (operationMode) {
      case OPERATION_MODE.OPERATION:
        setOperation(true);
        break;
      default:
        setOperation(false);
    }
  }, [ operationMode ])

  return (
    <>
      {open && (
        <>
          <div className={props.classes.func + ` ${operation ? props.classes.funcEditable : ''}`} >
            <div className={props.classes.funcPaper} >
              <Box m={4}>
                <Box style={{display: 'flex'}}>
                  <Settings style={{ color: grey[50] }} fontSize="small" />
                  <Typography align="left" component="div">
                    Flights
                  </Typography>
                </Box>
              </Box>
              <YourFlights classes={props.classes} open={open}/>
            </div>
          </div>
          <FlightOperationSlider classes={props.classes} />
        </>
      )}
    </>
  );
}

export default Flights;