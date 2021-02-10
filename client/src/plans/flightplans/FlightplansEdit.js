import React from 'react';

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

const FlightplansEdit = (props) => {

  const onClickCancel = () => {
    props.openDetail(props.id);
  }

  const onClickSave = () => {
    props.openList();
  }

  const onClickDelete = () => {
    props.openList();
  }

  const onClickReturn = () => {
    props.openList();
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <Grid container className={props.classes.textLabel}>
          <Grid item xs={12}>
            <Button onClick={onClickReturn}>
              <ChevronLeftIcon style={{ color: grey[50] }} />
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Typography>Edit Flightplan</Typography>
          </Grid>
          <Grid item xs={12}>
            <Box className={props.classes.textInput}
                p={1} m={1} borderRadius={7} >
              <TextField
                label="Name"
                name="name"
                fullWidth />
            </Box>
          </Grid>
        </Grid>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button
            className={props.classes.funcButton}
            onClick={onClickCancel}>
          Cancel
        </Button>
        <Button 
            className={props.classes.funcButton}
            onClick={onClickDelete}>
          Delete
        </Button>
        <Button
            className={props.classes.funcButton}
            onClick={onClickSave}>
          Save
        </Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default FlightplansEdit;