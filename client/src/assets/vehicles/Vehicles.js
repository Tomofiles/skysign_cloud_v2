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

const Vehicles = (props) => {
  const [edit, setEdit] = useState(false);
  const [selected, setSelected] = useState({id: undefined});

  const openEdit = (id) => {
    setEdit(true);
    setSelected(id);
  }

  const closeEdit = () => {
    setEdit(false);
    setSelected(undefined);
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
        <Typography>My Vehicles</Typography>
      </ExpansionPanelSummary>
      {edit ?
        <VehiclesEdit classes={props.classes} closeEdit={closeEdit} id={selected} />
      :
        <VehiclesList classes={props.classes} openEdit={openEdit} />
      }
    </ExpansionPanel>
  );
}

export default Vehicles;