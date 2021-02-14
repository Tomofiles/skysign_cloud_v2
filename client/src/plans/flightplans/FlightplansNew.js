import React from 'react';

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

const FlightplansNew = (props) => {
  const { register, handleSubmit, errors } = useForm();

  const onClickSave = () => {
    props.openList();
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
                      label="The number of vehicles"
                      type="text"
                      name="commId"
                      fullWidth
                      inputRef={register({ required: true, maxLength: 50 })}
                      error={Boolean(errors.commId)}
                      helperText={errors.commId} />
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