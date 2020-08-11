import React, { useEffect, useState } from 'react';
import { useForm, Controller } from 'react-hook-form';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  Grid,
  Box,
  Select,
  MenuItem,
  InputLabel,
  FormControl
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { v4 as uuidv4 } from 'uuid';

import { getVehicles } from '../assets/vehicles/VehicleUtils'
import { getMissions } from '../plans/missions/MissionUtils'

const noSelected = {
  "id": "no-selected",
  "name": "-"
};

const StagingNew = (props) => {
  const { control, handleSubmit } = useForm({
    defaultValues: {
      mission: "no-selected",
      vehicle: "no-selected"
    }
  });
  const [ missions, setMissions ] = useState([noSelected]);
  const [ vehicles, setVehicles ] = useState([noSelected]);

  useEffect(() => {
    getMissions()
      .then(data => {
        setMissions([noSelected, ...data.missions]);
      });
    getVehicles()
      .then(data => {
        setVehicles([noSelected, ...data.vehicles]);
      });
  }, [])

  const onClickReturn = () => {
    props.openList();  
  }

  const onClickOk = (data) => {
    data.id = uuidv4();
    data.selected = false;
    props.addRow(data);
    props.openList();  
  }

  return (
    <div>
      <form onSubmit={handleSubmit(onClickOk)}>
        <ExpansionPanelDetails>
          <Grid container className={props.classes.editVehicleInput}>
            <Grid item xs={12}>
              <Button onClick={onClickReturn}>
                <ChevronLeftIcon style={{ color: grey[50] }} />
              </Button>
            </Grid>
            <Grid item xs={12}>
              <Typography>New Staging</Typography>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.editVehicleInputText}
                  p={1} m={1} borderRadius={7} >
                <FormControl fullWidth>
                  <InputLabel id="mission-label">Mission</InputLabel>
                  <Controller
                    as={
                      <Select
                        labelId="mission-label"
                        id="mission"
                      >
                        {missions.map((mission) => (
                          <MenuItem key={mission.id} value={mission.id}>{mission.name}</MenuItem>
                        ))}
                      </Select>
                    }
                    name="mission"
                    control={control} />
                </FormControl>
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.editVehicleInputText}
                  p={1} m={1} borderRadius={7} >
                <FormControl fullWidth>
                  <InputLabel id="vehicle-label">Vehicle</InputLabel>
                  <Controller
                    as={
                      <Select
                        labelId="vehicle-label"
                        id="vehicle"
                      >
                        {vehicles.map((vehicle) => (
                          <MenuItem key={vehicle.id} value={vehicle.id}>{vehicle.name}</MenuItem>
                        ))}
                      </Select>
                    }
                    name="vehicle"
                    control={control} />
                </FormControl>
              </Box>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>
        <ExpansionPanelActions >
          <Button
              className={props.classes.editVehicleButton}
              type="submit" >
            OK
          </Button>
        </ExpansionPanelActions>
      </form>
    </div>
  );
}

export default StagingNew;