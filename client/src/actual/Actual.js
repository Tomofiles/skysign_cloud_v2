import React from 'react';

import {
  Drawer,
  Grid,
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

const Actual = (props) => {
  return (
    <Drawer
        className={props.classes.func}
        anchor='right'
        variant="persistent"
        classes={{
          paper: props.classes.funcPaper,
        }}
        open={props.open} >
      <Box m={2} alignContent="center">
        <Box >
          <Grid container>
            <Grid item xs={4} />
            <Grid item xs={1}>
              <Settings style={{ color: grey[50] }} fontSize="small" />
            </Grid>
            <Grid item xs={3}>
              <Typography align="left" component="div">
                Actual
              </Typography>
            </Grid>
            <Grid item xs={4} />
          </Grid>
        </Box>
      </Box>
    </Drawer>
  );
}

export default Actual;