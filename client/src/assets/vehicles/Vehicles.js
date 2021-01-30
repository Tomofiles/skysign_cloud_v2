import React, { useState } from 'react';

import {
  Typography,
  ExpansionPanel,
  ExpansionPanelSummary,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import VehiclesList from './VehiclesList'
import VehiclesEdit from './VehiclesEdit'
import VehiclesNew from './VehiclesNew'
import VehiclesDetail from './VehiclesDetail'

const VEHICLE_MODE = Object.freeze({"NEW":1, "EDIT":2, "DETAIL":3, "LIST":4});

const Vehicles = (props) => {
  const [mode, setMode] = useState(VEHICLE_MODE.LIST);
  const [selected, setSelected] = useState(undefined);

  const openEdit = (id) => {
    setMode(VEHICLE_MODE.EDIT);
    setSelected(id);
  }

  const openNew = () => {
    setMode(VEHICLE_MODE.NEW);
    setSelected(undefined);
  }

  const openDetail = (id) => {
    setMode(VEHICLE_MODE.DETAIL);
    setSelected(id);
  }

  const openList = () => {
    setMode(VEHICLE_MODE.LIST);
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
        <Typography>My Vehicles</Typography>
      </ExpansionPanelSummary>
      {mode === VEHICLE_MODE.EDIT &&
        <VehiclesEdit classes={props.classes} openList={openList} openDetail={openDetail} id={selected} />
      }
      {mode === VEHICLE_MODE.NEW &&
        <VehiclesNew classes={props.classes} openList={openList} />
      }
      {mode === VEHICLE_MODE.DETAIL &&
        <VehiclesDetail classes={props.classes} openList={openList} openEdit={openEdit} id={selected} />
      }
      {mode === VEHICLE_MODE.LIST &&
        <VehiclesList classes={props.classes} openDetail={openDetail} openNew={openNew} id={selected} open={props.open} />
      }
    </ExpansionPanel>
  );
}

export default Vehicles;