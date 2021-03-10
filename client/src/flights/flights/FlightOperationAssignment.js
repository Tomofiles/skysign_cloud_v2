import React, { useContext, useEffect, useState } from 'react';

import {
  Box,
  Paper,
  Typography,
  Divider,
  Grid,
  ListItem,
} from '@material-ui/core';
import FlightOperationCommunication from './FlightOperationCommunication';
import { AppContext } from '../../context/Context';

const FlightOperationAssignment = props => {
  const { vehicles, missions } = useContext(AppContext);
  const [ vehicleName, setVehicleName ] = useState("-");
  const [ missionName, setMissionName ] = useState("-");
  const [ communicationId, setCommunicationId ] = useState(undefined);

  useEffect(() => {
    vehicles
      .filter(vehicle => vehicle.id === props.vehicleId)
      .forEach(vehicle => {
        setVehicleName(vehicle.name);
        setCommunicationId(vehicle.commId);
      });
  }, [ props.vehicleId, vehicles, setVehicleName, setCommunicationId ])

  useEffect(() => {
    missions
      .filter(mission => mission.id === props.missionId)
      .forEach(mission => setMissionName(mission.name));
  }, [ props.missionId, missions, setMissionName ])

  return (
    <Box pb={1}>
      <ListItem component={Paper} className={props.classes.funcPanelEdit}>
        <Box>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={6}>
              <Box>
                <Typography>Vehicle</Typography>
              </Box>
            </Grid>
            <Grid item xs={6}>
              <Box>
                <Typography>{vehicleName}</Typography>
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Divider/>
            </Grid>
            <Grid item xs={6}>
              <Box pt={1}>
                <Typography>Mission</Typography>
              </Box>
            </Grid>
            <Grid item xs={6}>
              <Box pt={1}>
                <Typography>{missionName}</Typography>
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Divider/>
            </Grid>
            <FlightOperationCommunication 
              classes={props.classes} 
              communicationId={communicationId}
              missionId={props.missionId} />
          </Grid>
        </Box>
      </ListItem>
    </Box>
  );
}

export default FlightOperationAssignment;