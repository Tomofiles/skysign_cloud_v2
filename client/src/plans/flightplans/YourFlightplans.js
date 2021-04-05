import React, { useState } from 'react';

import FlightplansDetail from './FlightplansDetail';
import FlightplansList from './FlightplansList';
import FlightplansNew from './FlightplansNew';
import FlightplansEdit from './FlightplansEdit';
import AssignAssetsDetail from './AssignAssetsDetail';
import AssignAssetsEdit from './AssignAssetsEdit';
import { Box } from '@material-ui/core';
import ChangeNumberOfVehicles from './ChangeNumberOfVehicles';

const FLIGHTPLAN_MODE = Object.freeze({
  "NEW":1,
  "EDIT":2,
  "DETAIL":3,
  "LIST":4,
  "ASSIGN_EDIT":5,
  "ASSIGN_DETAIL":6,
  "CHANGE_NUMBER_OF_VEHICLES":7,
});

const YourFlightplans = (props) => {
  const [ mode, setMode ] = useState(FLIGHTPLAN_MODE.LIST);
  const [ selected, setSelected ] = useState(undefined);

  const openEdit = (id) => {
    setMode(FLIGHTPLAN_MODE.EDIT);
    setSelected(id);
  }

  const openNew = () => {
    setMode(FLIGHTPLAN_MODE.NEW);
    setSelected(undefined);
  }

  const openDetail = (id) => {
    setMode(FLIGHTPLAN_MODE.DETAIL);
    setSelected(id);
  }

  const openList = () => {
    setMode(FLIGHTPLAN_MODE.LIST);
    setSelected(undefined);
  }

  const openAssignEdit = (id) => {
    setMode(FLIGHTPLAN_MODE.ASSIGN_EDIT);
    setSelected(id);
  }

  const openAssignDetail = (id) => {
    setMode(FLIGHTPLAN_MODE.ASSIGN_DETAIL);
    setSelected(id);
  }

  const openChangeNumberOfVehicles = (id) => {
    setMode(FLIGHTPLAN_MODE.CHANGE_NUMBER_OF_VEHICLES);
    setSelected(id);
  }

  return (
    <Box px={4}>
      {mode === FLIGHTPLAN_MODE.EDIT &&
        <FlightplansEdit
          classes={props.classes}
          openList={openList}
          openDetail={openDetail}
          id={selected} />
      }
      {mode === FLIGHTPLAN_MODE.NEW &&
        <FlightplansNew
          classes={props.classes}
          openList={openList} />
      }
      {mode === FLIGHTPLAN_MODE.DETAIL &&
        <FlightplansDetail
          classes={props.classes}
          openList={openList}
          openEdit={openEdit}
          openAssignDetail={openAssignDetail}
          openChangeNumberOfVehicles={openChangeNumberOfVehicles}
          id={selected} />
      }
      {mode === FLIGHTPLAN_MODE.LIST &&
        <FlightplansList
          classes={props.classes}
          openDetail={openDetail}
          openNew={openNew}
          id={selected}
          open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.ASSIGN_EDIT &&
        <AssignAssetsEdit
          classes={props.classes}
          openList={openList}
          openAssignDetail={openAssignDetail}
          id={selected}
          open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.ASSIGN_DETAIL &&
        <AssignAssetsDetail
          classes={props.classes}
          openDetail={openDetail}
          openAssignEdit={openAssignEdit}
          id={selected}
          open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.CHANGE_NUMBER_OF_VEHICLES &&
        <ChangeNumberOfVehicles
          classes={props.classes}
          openList={openList}
          openDetail={openDetail}
          id={selected}
          open={props.open} />
      }
      </Box>
  );
}

export default YourFlightplans;