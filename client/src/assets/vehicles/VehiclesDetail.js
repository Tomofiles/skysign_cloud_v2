import React, { useState, useEffect } from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  Grid,
  Box
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { getVehicle } from './VehicleUtils'

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
  }, [props.id])

  const onClickEdit = (id) => {
    props.openEdit(id);
  }

  const onClickReturn = () => {
    props.openList();  
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <Grid container className={props.classes.editVehicleInput}>
          <Grid item xs={12}>
            <Button onClick={onClickReturn}>
              <ChevronLeftIcon style={{ color: grey[50] }} />
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Typography>Detail Vehicle</Typography>
          </Grid>
          <Grid item xs={12}>
            <Box  p={1} m={1} borderRadius={7} >
              <Grid container className={props.classes.editVehicleInput}>
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
              <Grid container className={props.classes.editVehicleInput}>
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
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button className={props.classes.editVehicleButton} onClick={() => onClickEdit(vehicle.id)}>Edit</Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default VehiclesDetail;