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

import { createVehicle } from './VehicleUtils'
import { AppContext } from '../../context/Context';

const default_vehicle = {name: "", communication_id: ""};

const VehiclesNew = (props) => {
  const { register, handleSubmit, errors } = useForm({defaultValues: default_vehicle});
  const { dispatchMessage } = useContext(AppContext);

  const onClickSave = (data) => {
    createVehicle(data)
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
            <Typography>Create Vehicle</Typography>
          </Box>
        </Box>
        <Box pb={2}>
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
                    <TextField
                      label="Name"
                      type="text"
                      name="name"
                      fullWidth
                      inputRef={register({
                        required: { value: true, message: "cannot be blank" },
                        maxLength: { value: 200, message: "the length must be no more than 200" },
                       })}
                      error={Boolean(errors.name)}
                      helperText={errors.name?.message}
                      />
                  </Box>
                </Grid>
                <Grid item xs={12}>
                  <Box className={props.classes.textInput}
                      p={1} m={1} borderRadius={7} >
                    <TextField
                      label="Communication ID"
                      type="text"
                      name="communication_id"
                      fullWidth
                      inputRef={register({
                        required: { value: true, message: "cannot be blank" },
                        maxLength: { value: 36, message: "the length must be no more than 36" },
                       })}
                      error={Boolean(errors.communication_id)}
                      helperText={errors.communication_id?.message}
                      />
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

export default VehiclesNew;