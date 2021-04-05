import React, { useState } from 'react';

import {
  Box,
} from '@material-ui/core';

import FlightsList from './FlightsList'
import FlightOperation from './FlightOperation'

const FLIGHT_MODE = Object.freeze({"LIST":1, "OPERATION":2});

const YourFlights = (props) => {
  const [ mode, setMode ] = useState(FLIGHT_MODE.LIST);
  const [ selected, setSelected ] = useState(undefined);

  const openList = () => {
    setMode(FLIGHT_MODE.LIST);
    setSelected(undefined);
  }

  const openOperation = (id) => {
    setMode(FLIGHT_MODE.OPERATION);
    setSelected(id);
  }

  return (
    <Box px={4}>
      {mode === FLIGHT_MODE.OPERATION &&
        <FlightOperation
          classes={props.classes}
          openList={openList}
          id={selected} />
      }
      {mode === FLIGHT_MODE.LIST &&
        <FlightsList
          classes={props.classes}
          openOperation={openOperation}
          id={selected}
          open={props.open} />
      }
    </Box>
  );
}

export default YourFlights;