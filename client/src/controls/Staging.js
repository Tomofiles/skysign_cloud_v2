import React, { useState } from 'react';

import {
  Typography,
  ExpansionPanel,
  ExpansionPanelSummary
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import StagingList from './StagingList';
import StagingNew from './StagingNew';

import { getVehicle } from '../assets/vehicles/VehicleUtils'

const STAGING_MODE = Object.freeze({"NEW":1, "LIST":2});

const getVehicleName = async (id) => {
  const vehicle = await getVehicle(id);
  return vehicle.name;
}

const Staging = (props) => {
  const [ mode, setMode ] = useState(STAGING_MODE.LIST);
  const [ rows, setRows ] = useState([]);

  const openNew = () => {
    setMode(STAGING_MODE.NEW);
  }

  const openList = () => {
    setMode(STAGING_MODE.LIST);
  }

  const addRow = async (data) => {
    setMode(STAGING_MODE.LIST);
    data.vehicleName = await getVehicleName(data.vehicle);
    setRows([...rows, data]);
  }

  const toggleSelectRow = (id) => {
    const newRows = [...rows];
    for (let row of newRows) {
      if (row.id === id) {
        row.selected = !row.selected;
      }
    }
    setRows(newRows);
  }

  const removeRows = () => {
    const newRows = rows
        .filter(row => !row.selected);
    setRows(newRows);
  }

  return (
    <ExpansionPanel
        className={props.classes.myVehicleRoot}
        defaultExpanded>
      <ExpansionPanelSummary
        expandIcon={<ExpandMoreIcon style={{ color: grey[50] }} />}
        aria-controls="panel1a-content"
        id="panel1a-header"
        className={props.classes.myVehicleSummary}
      >
        <Typography>Staging</Typography>
      </ExpansionPanelSummary>
      {mode === STAGING_MODE.NEW &&
        <StagingNew classes={props.classes} openList={openList} addRow={addRow} />
      }
      {mode === STAGING_MODE.LIST &&
        <StagingList
          classes={props.classes}
          openNew={openNew}
          toggleSelectRow={toggleSelectRow}
          removeRows={removeRows}
          rows={rows} />
      }
    </ExpansionPanel>
  );
}

export default Staging;