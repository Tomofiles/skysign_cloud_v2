import React, { useState, useEffect } from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  Paper,
  Divider
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { deleteVehicle, getVehicle } from './VehicleUtils'

const VehiclesDetail = (props) => {
  const [vehicle, setVehicle] = useState({id: "-", name: "-", commId: "-"});

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

  const onClickEdit = () => {
    props.openEdit(vehicle.id);
  }

  const onClickReturn = () => {
    props.openList();  
  }

  const onClickDelete = () => {
    deleteVehicle(vehicle.id)
      .then(data => {
        props.openList();
      })
  }

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>{vehicle.name}</Typography>
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
      <Box>
        <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
          <Box px={1}>
            <Button 
                className={props.classes.funcButton}
                onClick={onClickDelete}>
              Delete
            </Button>
          </Box>
          <Box px={1}>
            <Button
                className={props.classes.funcButton}
                onClick={onClickEdit}>
              Edit
            </Button>
          </Box>
        </Box>
      </Box>
    </div>
  );
}

export default VehiclesDetail;