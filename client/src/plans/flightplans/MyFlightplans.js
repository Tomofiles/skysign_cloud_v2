import React, { useState } from 'react';

import {
  Typography,
  ExpansionPanel,
  ExpansionPanelSummary,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import FlightplansDetail from './FlightplansDetail';
import FlightplansList from './FlightplansList';
import FlightplansNew from './FlightplansNew';
import FlightplansEdit from './FlightplansEdit';
import AssignDetail from './AssignDetail';
import AssignEdit from './AssignEdit';

const FLIGHTPLAN_MODE = Object.freeze({"NEW":1, "EDIT":2, "DETAIL":3, "LIST":4, "ASSIGN_EDIT":5, "ASSIGN_DETAIL":6});

const MyFlightplans = (props) => {
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

  return (
    <ExpansionPanel
        className={props.classes.funcPanel}
        defaultExpanded>
      <ExpansionPanelSummary
        expandIcon={<ExpandMoreIcon style={{ color: grey[50] }} />}
        aria-controls="panel1a-content"
        id="panel1a-header"
        className={props.classes.funcPanelSummary}
      >
        <Typography>My Flightplans</Typography>
      </ExpansionPanelSummary>
      {mode === FLIGHTPLAN_MODE.EDIT &&
        <FlightplansEdit classes={props.classes} openList={openList} openDetail={openDetail} id={selected} />
      }
      {mode === FLIGHTPLAN_MODE.NEW &&
        <FlightplansNew classes={props.classes} openList={openList} />
      }
      {mode === FLIGHTPLAN_MODE.DETAIL &&
        <FlightplansDetail classes={props.classes} openList={openList} openEdit={openEdit} openAssignDetail={openAssignDetail} id={selected} />
      }
      {mode === FLIGHTPLAN_MODE.LIST &&
        <FlightplansList classes={props.classes} openDetail={openDetail} openNew={openNew} id={selected} open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.ASSIGN_EDIT &&
        <AssignEdit classes={props.classes} openList={openList} openAssignDetail={openAssignDetail} id={selected} open={props.open} />
      }
      {mode === FLIGHTPLAN_MODE.ASSIGN_DETAIL &&
        <AssignDetail classes={props.classes} openDetail={openDetail} openAssignEdit={openAssignEdit} id={selected} open={props.open} />
      }
    </ExpansionPanel>
  );
}

export default MyFlightplans;