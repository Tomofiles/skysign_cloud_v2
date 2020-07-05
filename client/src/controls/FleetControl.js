import React, { useState, useEffect, useGlobal } from 'reactn';

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

import { arm, disarm } from './FleetControlUtils'

const FleetControl = (props) => {
  const [ buttonState, setButtonState] = useState({
    arm: true,
    disarm: true
  });
  const [ rows ] = useGlobal("stagingRows");

  const onClickArm = () => {
    rows.filter(row => row.selected)
        .forEach(row => arm(row.vehicle));
  }

  const onClickDisarm = () => {
    rows.filter(row => row.selected)
        .forEach(row => disarm(row.vehicle));
  }

  useEffect(() => {
    if (rows.filter(row => row.selected).length === 0) {
      setButtonState({
        arm: true,
        disarm: true
      });
    } else {
      setButtonState({
        arm: false,
        disarm: false
      });
    }
  }, [rows]);

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
            <Button disabled={buttonState.arm} className={props.classes.myVehicleButton} fullWidth onClick={onClickArm}>Arm</Button>
          </Grid>
          <Grid item xs={4} >
            <Button disabled={buttonState.disarm} className={props.classes.myVehicleButton} fullWidth onClick={onClickDisarm}>Disarm</Button>
          </Grid>
          <Grid item xs={4} />
          <Grid item xs={4} >
            <Button disabled className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Upload</Button>
          </Grid>
          <Grid item xs={4} >
            <Button disabled className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Start</Button>
          </Grid>
          <Grid item xs={4} >
            <Button disabled className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Pause</Button>
          </Grid>
          <Grid item xs={4} >
            <Button disabled className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Take Off</Button>
          </Grid>
          <Grid item xs={4} >
            <Button disabled className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Land</Button>
          </Grid>
          <Grid item xs={4} >
            <Button disabled className={props.classes.myVehicleButton} fullWidth onClick={() => {}}>Return</Button>
          </Grid>
        </Grid>
      </ExpansionPanelActions>
    </ExpansionPanel>
  );
}

export default FleetControl;