import React, { useState, useEffect, useContext } from 'react';

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

import { COMMAND_TYPE, controlCommunication } from './FleetControlUtils'
import { AppContext } from '../../context/Context';

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
  const { stagingRows } = useContext(AppContext);

  const onClickControl = (type) => {
    return () => {
      stagingRows.filter(row => row.isControlled)
          .forEach(row => controlCommunication(type, row.id));
    }
  }

  useEffect(() => {
    if (stagingRows.filter(row => row.isControlled).length === 0) {
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
  }, [stagingRows]);

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
        <Typography>Fleet Control</Typography>
      </ExpansionPanelSummary>
      <ExpansionPanelDetails>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Grid container spacing={2} >
          <Grid item xs={4} >
            <Button
              disabled={buttonState.arm}
              className={props.classes.funcButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.ARM)}>
                Arm
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.disarm}
              className={props.classes.funcButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.DISARM)}>
                Disarm
              </Button>
          </Grid>
          <Grid item xs={4} />
          <Grid item xs={4} >
            <Button
              disabled={buttonState.upload}
              className={props.classes.funcButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.UPLOAD)}>
                Upload
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.start}
              className={props.classes.funcButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.START)}>
                Start
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.pause}
              className={props.classes.funcButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.PAUSE)}>
                Pause
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.takeoff}
              className={props.classes.funcButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.TAKEOFF)}>
                Take Off
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.land}
              className={props.classes.funcButton}
              fullWidth
              onClick={onClickControl(COMMAND_TYPE.LAND)}>
                Land
              </Button>
          </Grid>
          <Grid item xs={4} >
            <Button
              disabled={buttonState.return}
              className={props.classes.funcButton}
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