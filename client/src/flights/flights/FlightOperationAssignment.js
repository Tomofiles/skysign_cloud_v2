import React, { useContext, useEffect, useState } from 'react';

import {
  Box,
  Paper,
  Typography,
  Divider,
  Grid,
  ListItem,
  IconButton,
} from '@material-ui/core';
import Visibility from '@material-ui/icons/Visibility';
import { grey } from '@material-ui/core/colors';

import { getVehicle } from '../../assets/vehicles/VehicleUtils'
import FlightOperationCommunication from './FlightOperationCommunication';
import { AppContext } from '../../context/Context';

const FlightOperationAssignment = props => {
  const { missions } = useContext(AppContext);
  const [ vehicleName, setVehicleName ] = useState("-");
  const [ missionName, setMissionName ] = useState("-");
  const [ communicationId, setCommunicationId ] = useState(undefined);

  useEffect(() => {
    getVehicle(props.vehicleId)
      .then(data => {
        setVehicleName(data.name);
        setCommunicationId(data.commId);
      })
  }, [ props.vehicleId, setVehicleName, setCommunicationId ])

  useEffect(() => {
    console.log(props.missionId);
    console.log(missions);
    missions
      .filter(mission => mission.id === props.missionId)
      .forEach(mission => setMissionName(mission.name));
  }, [ props.missionId, missions, setMissionName ])

  const onClickJump = () => {

  }

  return (
    <Box pb={1}>
      <ListItem component={Paper} className={props.classes.funcPanelEdit}>
        <Box>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
                <IconButton size="small" onClick={onClickJump}>
                  <Visibility style={{ color: grey[50] }} />
                </IconButton>
              </Box>
            </Grid>
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
              communicationId={communicationId} />
          </Grid>
        </Box>
      </ListItem>
    </Box>
  );
}

export default FlightOperationAssignment;