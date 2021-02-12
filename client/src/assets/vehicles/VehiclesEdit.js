import React, { useEffect, useState } from 'react';

import {
  Typography,
  Button,
  TextField,
  Grid,
  Box,
  Paper,
  Divider,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';
import { useForm, Controller } from 'react-hook-form';

import { getVehicle, updateVehicle } from './VehicleUtils'

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
  }, [ props.id, setValue ])

  const onClickCancel = () => {
    props.openDetail(id);
  }

  const onClickSave = (data) => {
    updateVehicle(id, data)
      .then(ret => {
        props.openList();
      });
  }

  const onClickReturn = () => {
    props.openList();  
  }

  return (
    <form onSubmit={handleSubmit(onClickSave)}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Edit Vehicle</Typography>
        </Box>
      </Box>
      <Paper className={props.classes.funcPanelEdit}>
        <Box p={3}>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Typography>Vehicle settings</Typography>
              <Divider/>
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
          <Divider/>
        </Box>
        <Box p={3}>
          <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
            <Box px={1}>
              <Button
                  className={props.classes.funcButton}
                  onClick={onClickCancel}>
                Cancel
              </Button>
            </Box>
            <Box px={1}>
              <Button
                  className={props.classes.funcButton}
                  type="submit" >
                Save
              </Button>
            </Box>
          </Box>
        </Box>
      </Paper>
    </form>
  );
}

export default VehiclesEdit;