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

import { getMissions } from '../../plans/missions/MissionUtils'

const noSelected = {
  "id": "no-selected",
  "name": "-"
};

const StagingEdit = (props) => {
  const { control, handleSubmit, setValue } = useForm({
    defaultValues: {
      missionId: "no-selected"
    }
  });
  const [ missions, setMissions ] = useState([noSelected]);

  useEffect(() => {
    getMissions()
      .then(data => {
        setMissions([noSelected, ...data.missions]);
      });
  }, [ setMissions ])

  useEffect(() => {
    if (props.selected.missionId) {
      setValue("missionId", props.selected.missionId);
    }
  }, [ props.selected, setValue ])

  const onClickReturn = () => {
    props.openList();  
  }

  const onClickStaging = data => {
    props.staging(props.selected.id, data);
  }

  const onClickCancel = () => {
    props.cancel(props.selected.id);
  }

  return (
    <div>
      <form onSubmit={handleSubmit(onClickStaging)}>
        <ExpansionPanelDetails>
          <Grid container className={props.classes.editVehicleInput}>
            <Grid item xs={12}>
              <Button onClick={onClickReturn}>
                <ChevronLeftIcon style={{ color: grey[50] }} />
              </Button>
            </Grid>
            <Grid item xs={12}>
              <Typography>Staging / Cancle</Typography>
            </Grid>
            <Grid item xs={12}>
              <Box  p={1} m={1} borderRadius={7} >
                <Grid container className={props.classes.editVehicleInput}>
                  <Grid item xs={12}>
                    <Typography style={{fontSize: "12px"}}>Name</Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography>{props.selected.vehicleName}</Typography>
                  </Grid>
                </Grid>
              </Box>
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
                        id="missionId"
                      >
                        {missions.map((mission) => (
                          <MenuItem key={mission.id} value={mission.id}>{mission.name}</MenuItem>
                        ))}
                      </Select>
                    }
                    name="missionId"
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
            Staging
          </Button>
          <Button
              className={props.classes.editVehicleButton}
              onClick={onClickCancel} >
            Cancel
          </Button>
        </ExpansionPanelActions>
      </form>
    </div>
  );
}

export default StagingEdit;