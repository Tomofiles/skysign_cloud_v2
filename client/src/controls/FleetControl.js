import React from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  ExpansionPanel,
  ExpansionPanelSummary
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

const FleetControl = (props) => {
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
        <Typography>Fleet Control</Typography>
      </ExpansionPanelSummary>
      <ExpansionPanelDetails>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Arm</Button>
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Disarm</Button>
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Upload</Button>
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Start</Button>
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Land</Button>
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Return</Button>
      </ExpansionPanelActions>
    </ExpansionPanel>
  );
}

export default FleetControl;