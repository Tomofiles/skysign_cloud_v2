import React from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  Grid,
  Box,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

const FlightplansDetail = (props) => {

  const onClickEdit = () => {
    props.openEdit(props.id);
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
            <Typography>Detail Flightplan</Typography>
          </Grid>
          <Grid item xs={12}>
            <Box  p={1} m={1} borderRadius={7} >
              <Grid container className={props.classes.textLabel}>
                <Grid item xs={12}>
                  <Typography style={{fontSize: "12px"}}>Name</Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography>Sample Flightplan</Typography>
                </Grid>
              </Grid>
            </Box>
          </Grid>
        </Grid>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button
            className={props.classes.funcButton}
            onClick={onClickEdit}>
          Edit
        </Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default FlightplansDetail;