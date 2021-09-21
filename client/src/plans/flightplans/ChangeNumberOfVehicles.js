import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  Paper,
  Divider,
  TextField,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';
import { Controller, useForm } from 'react-hook-form';

import { getFlightplan, changeNumberOfVehicles, getAssignments } from './FlightplansUtils';
import { AppContext } from '../../context/Context';

const ChangeNumberOfVehicles = (props) => {
  const { control, handleSubmit, setValue } = useForm();
  const [ flightplanName, setFlightplanName ] = useState("");
  const { dispatchMessage } = useContext(AppContext);

  useEffect(() => {
    if (props.open) {
      getFlightplan(props.id)
        .then(data => {
          setFlightplanName(data.name);
          getAssignments(data.fleet_id)
            .then(data => {
              setValue("number_of_vehicles", data.assignments.length);
            })
            .catch(message => {
              dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
            });
        })
        .catch(message => {
          dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
        });
    }
  }, [ props.open, props.id, setFlightplanName, setValue, dispatchMessage ])

  const onClickCancel = () => {
    props.openDetail(props.id);
  }

  const onClickSave = (data) => {
    changeNumberOfVehicles(props.id, data)
      .then(ret => {
        dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `Updated number_of_vehicles for ${flightplanName} successfully` });
        props.openList();
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }

  const onClickReturn = () => {
    props.openDetail(props.id);
  }

  return (
    <div className={props.classes.funcPanel}>
      <form onSubmit={handleSubmit(onClickSave)}>
        <Box>
          <Button onClick={onClickReturn}>
            <ChevronLeftIcon style={{ color: grey[50] }} />
          </Button>
          <Box p={2} style={{display: 'flex'}}>
            <Typography>Change number of vehicles</Typography>
          </Box>
        </Box>
        <Box pb={2}>
          <Paper className={props.classes.funcPanelEdit}>
            <Box p={3}>
              <Grid container className={props.classes.textLabel}>
                <Grid item xs={12}>
                  <Typography>Fleet formation configuration</Typography>
                  <Divider/>
                </Grid>
                <Grid item xs={12}>
                  <Box className={props.classes.textInput}
                      p={1} m={1} borderRadius={7} >
                    <Controller
                      as={TextField}
                      label="Number of vehicles"
                      name="number_of_vehicles"
                      control={control}
                      defaultValue=""
                      fullWidth />
                  </Box>
                </Grid>
              </Grid>
            </Box>
          </Paper>
        </Box>
        <Box>
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
      </form>
    </div>
  );
}

export default ChangeNumberOfVehicles;