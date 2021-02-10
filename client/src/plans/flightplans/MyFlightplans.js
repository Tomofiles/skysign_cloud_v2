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

const FLIGHTPLAN_MODE = Object.freeze({"NEW":1, "EDIT":2, "DETAIL":3, "LIST":4});

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
        <FlightplansDetail classes={props.classes} openList={openList} openEdit={openEdit} id={selected} />
      }
      {mode === FLIGHTPLAN_MODE.LIST &&
        <FlightplansList classes={props.classes} openDetail={openDetail} openNew={openNew} id={selected} open={props.open} />
      }
    </ExpansionPanel>
  );
}

export default MyFlightplans;