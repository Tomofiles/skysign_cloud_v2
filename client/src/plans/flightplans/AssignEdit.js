import React from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  Grid,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

const AssignEdit = (props) => {

  const onClickCancel = () => {
    props.openAssignDetail(props.id);
  }

  const onClickSave = () => {
    props.openList();
  }

  const onClickReturn = () => {
    props.openAssignDetail(props.id);
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
            <Typography>Edit Flightplan Assignment</Typography>
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
            onClick={onClickSave}>
          Save
        </Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default AssignEdit;