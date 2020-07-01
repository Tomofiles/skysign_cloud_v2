import React from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  TextField,
  Grid,
  Box
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';
import { useForm } from 'react-hook-form';

import { createVehicle } from './VehicleUtils'

const VehiclesNew = (props) => {
  const { register, handleSubmit, errors } = useForm();

  const onClickSave = (data) => {
    createVehicle(data)
      .then(ret => {
        props.openList();
      });
  }

  const onClickReturn = () => {
    props.openList();  
  }

  return (
    <div>
      <form onSubmit={handleSubmit(onClickSave)}>
        <ExpansionPanelDetails>
          <Grid container className={props.classes.editVehicleInput}>
            <Grid item xs={12}>
              <Button onClick={onClickReturn}>
                <ChevronLeftIcon style={{ color: grey[50] }} />
              </Button>
            </Grid>
            <Grid item xs={12}>
              <Typography>New Vehicle</Typography>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.editVehicleInputText}
                  p={1} m={1} borderRadius={7} >
                <TextField
                  label="Name"
                  type="text"
                  name="name"
                  fullWidth
                  inputRef={register({ required: true, maxLength: 50 })}
                  error={Boolean(errors.name)}
                  helperText={errors.name} />
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.editVehicleInputText}
                  p={1} m={1} borderRadius={7} >
                <TextField
                  label="Communication ID"
                  type="text"
                  name="commId"
                  fullWidth
                  inputRef={register({ required: true, maxLength: 50 })}
                  error={Boolean(errors.commId)}
                  helperText={errors.commId} />
              </Box>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>
        <ExpansionPanelActions >
          <Button
              className={props.classes.editVehicleButton}
              type="submit" >
            Save
          </Button>
        </ExpansionPanelActions>
      </form>
    </div>
  );
}

export default VehiclesNew;