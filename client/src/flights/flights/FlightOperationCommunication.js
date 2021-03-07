import React, { useEffect, useState } from 'react';

import {
  Box,
  Typography,
  Grid,
  Stepper,
  Step,
  StepButton,
  Divider,
} from '@material-ui/core';
import { getTelemetry } from '../../map/MapUtils'

const initialTelemetry = {
    latitude: "-",
    longitude: "-",
    altitude: "-",
    relativeAltitude: "-",
    speed: "-",
    armed: "-",
    flightMode: "-",
    orientationX: "-",
    orientationY: "-",
    orientationZ: "-",
    orientationW: "-",
};

function getSteps() {
  return ['UPLOAD', 'TAKEOFF', 'START', 'LAND', 'END'];
}

const FlightOperationCommunication = props => {
  const [ telemetry, setTelemetry ] = useState(initialTelemetry);
  const steps = getSteps();

  useEffect(() => {
    if (props.communicationId) {
      getTelemetry(props.communicationId)
        .then(data => {
          setTelemetry(data.telemetry);
        })
    }
  }, [ props.communicationId, setTelemetry ])

  return (
    <>
      <Grid item xs={12}>
        <Box pt={1}>
          <Typography>Telemetry</Typography>
        </Box>
      </Grid>
      <Grid item xs={3}>
        <Box p={1}>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }} >lat (deg)</Typography>
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
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>{telemetry.latitude}</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>{telemetry.altitude}</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>{telemetry.relativeAltitude}</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>{telemetry.speed}</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>{telemetry.flightMode}</Typography>
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
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>{telemetry.longitude}</Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}></Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}></Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}></Typography>
            </Grid>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>{telemetry.armed}</Typography>
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
        <Stepper alternativeLabel>
          {steps.map((label) => (
            <Step key={label}>
              <StepButton>
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