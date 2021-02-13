import React, { useState } from 'react';

import FlightplansDetail from './FlightplansDetail';
import FlightplansList from './FlightplansList';
import FlightplansNew from './FlightplansNew';
import FlightplansEdit from './FlightplansEdit';
import AssignVehicleDetail from './AssignVehicleDetail';
import AssignVehicleEdit from './AssignVehicleEdit';
import { Box } from '@material-ui/core';
import AssignMissionEdit from './AssignMissionEdit';
import AssignMissionDetail from './AssignMissionDetail';

const FLIGHTPLAN_MODE = Object.freeze({
  "NEW":1,
  "EDIT":2,
  "DETAIL":3,
  "LIST":4,
  "ASSIGN_VEHICLE_EDIT":5,
  "ASSIGN_VEHICLE_DETAIL":6,
  "ASSIGN_MISSION_EDIT":7,
  "ASSIGN_MISSION_DETAIL":8
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

  const openAssignVehicleEdit = (id) => {
    setMode(FLIGHTPLAN_MODE.ASSIGN_VEHICLE_EDIT);
    setSelected(id);
  }

  const openAssignVehicleDetail = (id) => {
    setMode(FLIGHTPLAN_MODE.ASSIGN_VEHICLE_DETAIL);
    setSelected(id);
  }

  const openAssignMissionEdit = (id) => {
    setMode(FLIGHTPLAN_MODE.ASSIGN_MISSION_EDIT);
    setSelected(id);
  }

  const openAssignMissionDetail = (id) => {
    setMode(FLIGHTPLAN_MODE.ASSIGN_MISSION_DETAIL);
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
          openAssignVehicleDetail={openAssignVehicleDetail}
          openAssignMissionDetail={openAssignMissionDetail}
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
      {mode === FLIGHTPLAN_MODE.ASSIGN_VEHICLE_EDIT &&
        <AssignVehicleEdit
          classes={props.classes}
          openList={openList}
          openAssignVehicleDetail={openAssignVehicleDetail}
          id={selected}
          open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.ASSIGN_VEHICLE_DETAIL &&
        <AssignVehicleDetail
          classes={props.classes}
          openDetail={openDetail}
          openAssignVehicleEdit={openAssignVehicleEdit}
          id={selected}
          open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.ASSIGN_MISSION_EDIT &&
        <AssignMissionEdit
          classes={props.classes}
          openList={openList}
          openAssignMissionDetail={openAssignMissionDetail}
          id={selected}
          open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.ASSIGN_MISSION_DETAIL &&
        <AssignMissionDetail
          classes={props.classes}
          openDetail={openDetail}
          openAssignMissionEdit={openAssignMissionEdit}
          id={selected}
          open={props.open} />
      }
      </Box>
  );
}

export default YourFlightplans;