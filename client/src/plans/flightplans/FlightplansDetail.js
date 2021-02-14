import React from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  Paper,
  Divider,
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

  const onClickAssignAssets = () => {
    props.openAssignDetail(props.id);
  }

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Sample Flightplan</Typography>
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
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>The number of vehicles</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography>3</Typography>
                    </Grid>
                  </Grid>
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
                onClick={onClickAssignAssets}>
              Assign Assets
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

export default FlightplansDetail;