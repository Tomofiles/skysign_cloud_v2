import React, { useState, useEffect } from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  Paper,
  Divider,
  Select,
  FormControl,
  MenuItem,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ArrowDropDown from '@material-ui/icons/ArrowDropDown';
import ArrowDropUp from '@material-ui/icons/ArrowDropUp';
import { grey } from '@material-ui/core/colors';

import { deleteVehicle, getVehicle } from './VehicleUtils'

const EDIT_VEHICLE = "edit_vehicle"
const DELETE_VEHICLE = "delete_vehicle"

const VehiclesDetail = (props) => {
  const [vehicle, setVehicle] = useState({id: "-", name: "-", commId: "-"});
  const [ openAction, setOpenAction ] = useState(false);

  useEffect(() => {
    getVehicle(props.id)
      .then(data => {
        setVehicle({
          id: data.id,
          name: data.name,
          commId: data.commId
        });
      })
  }, [ props.id ])

  const onClickReturn = () => {
    props.openList();  
  }

  const handleActionChange = e => {
    switch (e.target.value) {
      case EDIT_VEHICLE:
        props.openEdit(vehicle.id);
        break;
      case DELETE_VEHICLE:
        deleteVehicle(vehicle.id)
          .then(data => {
            props.openList();
          })
        break;
      default:
        break;
    }
  };

  const handleActionClose = () => {
    setOpenAction(false);
  };

  const handleActionOpen = () => {
    setOpenAction(true);
  };

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Button onClick={onClickReturn}>
            <ChevronLeftIcon style={{ color: grey[50] }} />
          </Button>
        </Box>
        <Box m={2} style={{display: 'flex', justifyContent: 'space-between'}}>
          <Box>
            <Typography>{vehicle.name}</Typography>
          </Box>
          <Box style={{display: 'flex'}}>
            <Box px={1}>
              <FormControl>
                <Button
                    id="openMenu"
                    className={props.classes.funcButton}
                    onClick={handleActionOpen} >
                  Action
                  {!openAction ? (
                    <ArrowDropDown fontSize="small"/>
                  ) : (
                    <ArrowDropUp fontSize="small"/>
                  )}
                </Button>
                <Select
                  onChange={handleActionChange}
                  style={{ display: "none" }}
                  open={openAction}
                  onClose={handleActionClose}
                  value=""
                  MenuProps={{
                    anchorEl: document.getElementById("openMenu"),
                    style: { marginTop: 60 }
                  }}
                >
                  <MenuItem value={EDIT_VEHICLE}>Edit Vehicle</MenuItem>
                  <MenuItem value={DELETE_VEHICLE}>Delete Vehicle</MenuItem>
                </Select>
              </FormControl>
            </Box>
          </Box>
        </Box>
      </Box>
      <Box pb={2}>
        <Paper className={props.classes.funcPanelEdit}>
          <Box p={3}>
            <Grid container className={props.classes.textLabel}>
              <Grid item xs={12}>
                <Typography>Vehicle details</Typography>
                <Divider/>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>Name</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography>{vehicle.name}</Typography>
                    </Grid>
                  </Grid>
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>Communication ID</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography>{vehicle.commId}</Typography>
                    </Grid>
                  </Grid>
                </Box>
              </Grid>
            </Grid>
            <Divider/>
          </Box>
        </Paper>
      </Box>
    </div>
  );
}

export default VehiclesDetail;