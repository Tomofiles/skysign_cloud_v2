import React from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  ExpansionPanel,
  ExpansionPanelSummary,
  Grid
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
        <Grid container spacing={2} >
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Arm</Button>
          </Grid>
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Disarm</Button>
          </Grid>
          <Grid item xs={4} />
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Upload</Button>
          </Grid>
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Start</Button>
          </Grid>
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Pause</Button>
          </Grid>
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Take Off</Button>
          </Grid>
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Land</Button>
          </Grid>
          <Grid item xs={4} >
            <Button className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Return</Button>
          </Grid>
        </Grid>
      </ExpansionPanelActions>
    </ExpansionPanel>
  );
}

export default FleetControl;