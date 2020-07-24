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

import { COMMAND_TYPE, controlVehicle } from './FleetControlUtils'

const FleetControl = (props) => {
  const [ buttonState, setButtonState] = useState({
    arm: true,
    disarm: true,
    upload: true,
    start: true,
    pause: true,
    takeoff: true,
    land: true,
    return: true,
  });
  const [ rows ] = useGlobal("stagingRows");

  const onClickControl = (type) => {
    return () => {
      rows.filter(row => row.selected)
          .forEach(row => controlVehicle(type, row.vehicle));
    }
  }

  useEffect(() => {
    if (rows.filter(row => row.selected).length === 0) {
      setButtonState({
        arm: true,
        disarm: true,
        upload: true,
        start: true,
        pause: true,
        takeoff: true,
        land: true,
        return: true,
      });
    } else {
      setButtonState({
        arm: false,
        disarm: false,
        upload: false,
        start: false,
        pause: false,
        takeoff: false,
        land: false,
        return: false,
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
            <Button
              disabled={buttonState.arm}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.ARM)}>
                Arm
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.disarm}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.DISARM)}>
                Disarm
              </Button>
          </Grid>
          <Grid item xs={4} />
          <Grid item xs={4} >
            <Button
              disabled={buttonState.upload}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.UPLOAD)}>
                Upload
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.start}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.START)}>
                Start
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.pause}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.PAUSE)}>
                Pause
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.takeoff}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.TAKEOFF)}>
                Take Off
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.land}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.LAND)}>
                Land
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.return}
              className={props.classes.myVehicleButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.RETURN)}>
                Return
              </Button>
          </Grid>
        </Grid>
      </ExpansionPanelActions>
    </ExpansionPanel>
  );
}

export default FleetControl;