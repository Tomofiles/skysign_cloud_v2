import React, { useEffect } from 'react';

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
import { Controller, useForm } from 'react-hook-form';

import { getFlightplan, updateFlightplan } from './FlightplansUtils';

const FlightplansEdit = (props) => {
  const { control, handleSubmit, setValue } = useForm();

  useEffect(() => {
    getFlightplan(props.id)
      .then(data => {
        setValue("name", data.name);
        setValue("description", data.description);
      })
  }, [ props.id, setValue ])

  const onClickCancel = () => {
    props.openDetail(props.id);
  }

  const onClickSave = (data) => {
    updateFlightplan(props.id, data)
      .then(ret => {
        props.openList();
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
            <Typography>Edit Flightplan</Typography>
          </Box>
        </Box>
        <Box pb={2}>
          <Paper className={props.classes.funcPanelEdit}>
            <Box p={3}>
              <Grid container className={props.classes.textLabel}>
                <Grid item xs={12}>
                  <Typography>Basic configuration</Typography>
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
                      label="Description"
                      name="description"
                      multiline
                      rows={4}
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

export default FlightplansEdit;