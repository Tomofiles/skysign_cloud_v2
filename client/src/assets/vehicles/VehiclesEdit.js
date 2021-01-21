import React, { useEffect, useState } from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  TextField,
  Grid,
  Box,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';
import { useForm, Controller } from 'react-hook-form';

import { getVehicle, updateVehicle, deleteVehicle } from './VehicleUtils'

const VehiclesEdit = (props) => {
  const [ id, setId ] = useState("");
  const { control, handleSubmit, setValue } = useForm();

  useEffect(() => {
    getVehicle(props.id)
      .then(data => {
        setId(data.id);
        setValue("name", data.name);
        setValue("commId", data.commId);
      })
  }, [props.id, setValue])

  const onClickCancel = (id) => {
    props.openDetail(id);
  }

  const onClickSave = (data) => {
    updateVehicle(id, data)
      .then(ret => {
        props.openList();
      });
  }

  const onClickDelete = (id) => {
    deleteVehicle(id)
      .then(data => {
        props.openList();
      })
  }

  const onClickReturn = () => {
    props.openList();  
  }

  return (
    <div>
      <form onSubmit={handleSubmit(onClickSave)}>
        <ExpansionPanelDetails>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Button onClick={onClickReturn}>
                <ChevronLeftIcon style={{ color: grey[50] }} />
              </Button>
            </Grid>
            <Grid item xs={12}>
              <Typography>Edit Vehicle</Typography>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.textInput}
                  p={1} m={1} borderRadius={7} >
                <Controller
                  as={TextField}
                  label="Name"
                  name="name"
                  control={control}
                  defaultValue=""
                  fullWidth />
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.textInput}
                  p={1} m={1} borderRadius={7} >
                <Controller
                  as={TextField}
                  label="Communication ID"
                  name="commId"
                  control={control}
                  defaultValue=""
                  fullWidth />
              </Box>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>
        <ExpansionPanelActions >
          <Button
              className={props.classes.funcButton}
              onClick={() => onClickCancel(id)}>
            Cancel
          </Button>
          <Button 
              className={props.classes.funcButton}
              onClick={() => onClickDelete(id)}>
            Delete
          </Button>
          <Button
              className={props.classes.funcButton}
              type="submit" >
            Save
          </Button>
        </ExpansionPanelActions>
      </form>
    </div>
  );
}

export default VehiclesEdit;