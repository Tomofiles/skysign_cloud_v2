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
  const [ missionUploadId, setMissionUploadId ] = useState("-");
  const [ communicationId, setCommunicationId ] = useState(undefined);

  useEffect(() => {
    vehicles
      .filter(vehicle => vehicle.id === props.vehicleId)
      .forEach(vehicle => {
        setVehicleName(vehicle.name);
        setCommunicationId(vehicle.communication_id);
      });
  }, [ props.vehicleId, vehicles, setVehicleName, setCommunicationId ])

  useEffect(() => {
    missions
      .filter(mission => mission.id === props.missionId)
      .forEach(mission => {
        setMissionName(mission.name);
        setMissionUploadId(mission.navigation.upload_id);
      });
  }, [ props.missionId, missions, setMissionName, setMissionUploadId ])

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
              missionId={missionUploadId} />
          </Grid>
        </Box>
      </ListItem>
    </Box>
  );
}

export default FlightOperationAssignment;