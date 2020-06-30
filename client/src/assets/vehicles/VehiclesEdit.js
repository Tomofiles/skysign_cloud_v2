import React, { useState, useEffect } from 'react';

import axios from 'axios';

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
import { useForm } from 'react-hook-form';

async function getVehicle(id) {
  try {
    const res = await axios
      .get(`/api/v1/vehicles/${id}`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

async function updateVehicle(id, data) {
  try {
    const res = await axios
      .put(`/api/v1/vehicles/${id}`, {
        name: data.vname,
        commId: data.commId
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

async function deleteVehicle(id) {
  try {
    const res = await axios
      .delete(`/api/v1/vehicles/${id}`, {})
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

const VehiclesEdit = (props) => {
  const { register, handleSubmit, errors } = useForm();
  const [vehicle, setVehicle] = useState({id: "-", vname: "-", commId: "-"});

  useEffect(() => {
    getVehicle(props.id)
      .then(data => {
        setVehicle({
          id: data.id,
          vname: data.name,
          commId: data.commId
        });
      })
  }, [props.id])

  const onClickCancel = (id) => {
    props.openDetail(id);
  }

  const onClickSave = (data) => {
    updateVehicle(vehicle.id, data)
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

  const handleChangeVname = (event) => {
    setVehicle({
      vname: event.target.vname,
      commId: vehicle.commId,
    });
  }

  const handleChangeCommId = (event) => {
    setVehicle({
      vname: vehicle.vname,
      commId: event.target.commId,
    });
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
              <Typography>Edit Vehicle</Typography>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.editVehicleInputText}
                  p={1} m={1} borderRadius={7} >
                <TextField
                  label="Name"
                  type="text"
                  name="vname"
                  fullWidth
                  inputRef={register({ required: true, maxLength: 50 })}
                  error={Boolean(errors.vname)}
                  helperText={errors.vname}
                  onChange={handleChangeVname}
                  value={vehicle.vname} />
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
                  helperText={errors.commId}
                  onChange={handleChangeCommId}
                  value={vehicle.commId} />
              </Box>
            </Grid>
          </Grid>
        </ExpansionPanelDetails>
        <ExpansionPanelActions >
          <Button
              className={props.classes.editVehicleButton}
              onClick={() => onClickCancel(vehicle.id)}>
            Cancel
          </Button>
          <Button 
              className={props.classes.editVehicleButton}
              onClick={() => onClickDelete(vehicle.id)}>
            Delete
          </Button>
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

export default VehiclesEdit;