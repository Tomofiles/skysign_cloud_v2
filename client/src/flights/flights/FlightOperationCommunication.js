import React, { useCallback, useContext, useEffect, useState } from 'react';

import {
  Box,
  Typography,
  Grid,
  Stepper,
  Step,
  StepButton,
  Divider,
  IconButton,
} from '@material-ui/core';
import Visibility from '@material-ui/icons/Visibility';
import { grey } from '@material-ui/core/colors';

import { AppContext } from '../../context/Context';
import NumberValue from './communication_value/NumberValue';
import ArmValue from './communication_value/ArmValue';

const initialTelemetry = {
    latitude: "-",
    longitude: "-",
    altitude: "-",
    relative_altitude: "-",
    speed: "-",
    armed: "-",
    flight_mode: "-",
    heading: "-",
    pitch: "-",
    roll: "-",
};

const allSteps = ['UPLOAD', 'TAKEOFF', 'START', 'PAUSE', 'LAND', 'RETURN'];

const FlightOperationCommunication = props => {
  const { telemetries, steps, dispatchSteps, dispatchMapPosition } = useContext(AppContext);
  const [ telemetry, setTelemetry ] = useState(initialTelemetry);
  const [ activeStep, setActiveStep ] = React.useState(0);

  useEffect(() => {
    if (props.communicationId) {
      telemetries
        .filter(telemetry => telemetry.id === props.communicationId)
        .forEach(telemetry => setTelemetry(telemetry.telemetry));
    }
  }, [ props.communicationId, telemetries, setTelemetry ])

  useEffect(() => {
    steps
      .filter(step => step.communication_id === props.communicationId)
      .forEach(step => setActiveStep(step.step));
  }, [ props.communicationId, steps, setActiveStep ])

  const handleStep = useCallback(step => {
    dispatchSteps({
      type: 'CHANGE_STEP',
      communication_id: props.communicationId,
      mission_id: props.missionId,
      step: step,
    });
  }, [ props.communicationId, props.missionId, dispatchSteps ]);

  useEffect(() => {
    handleStep(0);
  }, [ handleStep ])

  const onClickJump = () => {
    dispatchMapPosition({
      type: 'CURRENT',
      longitude: telemetry.longitude,
      latitude: telemetry.latitude,
      height: telemetry.altitude + 200,
    });
  }

  return (
    <>
      <Grid item xs={12}>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Box pt={1}>
            <Typography>Telemetry</Typography>
          </Box>
          <IconButton size="small" onClick={onClickJump}>
            <Visibility style={{ color: grey[50] }} />
          </IconButton>
        </Box>
      </Grid>
      <Grid item xs={3}>
        <Box p={1}>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>lat (deg)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>alt (m)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>rel alt (m)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>speed (m/s)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>state</Typography>
            </Grid>
          </Grid>
        </Box>
      </Grid>
      <Grid item xs={3}>
        <Box p={1}>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.latitude} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.altitude} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.relative_altitude} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.speed} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                {telemetry.flight_mode}
              </Typography>
            </Grid>
          </Grid>
        </Box>
      </Grid>
      <Grid item xs={3}>
        <Box p={1}>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>lon (deg)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>heading (deg)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>pitch (deg)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>roll (deg)</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>arm</Typography>
            </Grid>
          </Grid>
        </Box>
      </Grid>
      <Grid item xs={3}>
        <Box p={1}>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.longitude} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.heading} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.pitch} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <NumberValue value={telemetry.roll} />
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>
                <ArmValue value={telemetry.armed} />
              </Typography>
            </Grid>
          </Grid>
        </Box>
      </Grid>
      <Grid item xs={12}>
        <Divider/>
      </Grid>
      <Grid item xs={12}>
        <Box pt={1}>
          <Typography>Step</Typography>
        </Box>
      </Grid>
      <Grid item xs={12}>
        <Stepper alternativeLabel nonLinear activeStep={activeStep}>
          {allSteps.map((label, index) => (
            <Step key={label}>
              <StepButton onClick={() => handleStep(index)}>
                {label}
              </StepButton>
            </Step>
          ))}
        </Stepper>
      </Grid>
    </>
  );
}

export default FlightOperationCommunication;