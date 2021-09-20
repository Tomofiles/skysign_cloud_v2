import React, { useContext } from 'react';

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
import { useForm } from 'react-hook-form';

import { createFlightplan } from './FlightplansUtils';
import { AppContext } from '../../context/Context';

const FlightplansNew = (props) => {
  const { register, handleSubmit, errors } = useForm();
  const { dispatchMessage } = useContext(AppContext);

  const onClickSave = (data) => {
    createFlightplan(data)
      .then(ret => {
        dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `${ret.name} was created successfully` });
        props.openList();
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }

  const onClickCancel = () => {
    props.openList();  
  }

  const onClickReturn = () => {
    props.openList();  
  }

  return (
    <div className={props.classes.funcPanel}>
      <form onSubmit={handleSubmit(onClickSave)}>
        <Box>
          <Button onClick={onClickReturn}>
            <ChevronLeftIcon style={{ color: grey[50] }} />
          </Button>
          <Box p={2} style={{display: 'flex'}}>
            <Typography>Create Flightplan</Typography>
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
                  <Box className={props.classes.textInput}
                      p={1} m={1} borderRadius={7} >
                    <TextField
                      label="Description"
                      type="text"
                      name="description"
                      multiline
                      rows={4}
                      fullWidth
                      inputRef={register({ required: true, maxLength: 50 })}
                      error={Boolean(errors.description)}
                      helperText={errors.description} />
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

export default FlightplansNew;