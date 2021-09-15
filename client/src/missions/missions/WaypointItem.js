import React from 'react';
import {
  Grid,
  Typography,
  Paper,
  Slider,
  ExpansionPanel,
  ExpansionPanelSummary,
  ExpansionPanelActions,
  TextField,
  Box,
  IconButton,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Delete from '@material-ui/icons/Delete';

const WaypointItem = (props) => {

  const onChangeHeightTextField = event => {
    let height = Number.parseFloat(event.target.value);
    if (!height) {
      height = 0;
    }
    props.changeRelativeHeight(props.index, height);
  }

  const onChangeHeightSlider = (event, newValue) => {
    props.changeRelativeHeight(props.index, newValue);
  }

  const onChangeSpeedTextField = event => {
    let speed = Number.parseFloat(event.target.value);
    if (!speed) {
      speed = 0;
    }
    props.changeSpeed(props.index, speed);
  }

  const onChangeSpeedSlider = (event, newValue) => {
    props.changeSpeed(props.index, newValue);
  }

  const onClickRemove = (id) => {
    props.onClickRemove(id);
  }

  return (
    <Box py={0.5} >
      <ExpansionPanel
          component={Paper}
          className={props.classes.missionListItem} >
        <ExpansionPanelSummary>
          <Grid container>
            <Grid item xs={12}>
              <Typography style={{ fontSize: "8px" }} >WP{props.index + 1}</Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography style={{ fontSize: "8px" }} >lat: {props.waypoint.latitude}</Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography style={{ fontSize: "8px" }} >lon: {props.waypoint.longitude}</Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography style={{ fontSize: "8px" }} >rel height: {props.waypoint.relative_altitude} m</Typography>
            </Grid>
            <Grid item xs={6}>
              <Typography style={{ fontSize: "8px" }} >speed: {props.waypoint.speed} m/s</Typography>
            </Grid>
          </Grid>
        </ExpansionPanelSummary>
        {props.editable &&
        <ExpansionPanelActions>
        <Grid container >
          <Grid item xs={12}>
            <Box className={props.classes.textInput}
                p={1} m={1} borderRadius={7} >
              <Grid container >
                <Grid item xs={4}>
                  <TextField
                    label="rel height"
                    type="number"
                    size="small"
                    value={props.waypoint.relative_altitude}
                    onChange={onChangeHeightTextField} />
                </Grid>
                <Grid item xs={8}>
                  <Box p={2} >
                    <Slider
                      color={"secondary"}
                      value={props.waypoint.relative_altitude}
                      onChange={onChangeHeightSlider} />
                  </Box>
                </Grid>
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={12}>
            <Box className={props.classes.textInput}
                p={1} m={1} borderRadius={7} >
              <Grid container >
                <Grid item xs={4}>
                  <TextField
                    label="speed"
                    type="number"
                    size="small"
                    value={props.waypoint.speed}
                    onChange={onChangeSpeedTextField} />
                </Grid>
                <Grid item xs={8}>
                  <Box p={2} >
                    <Slider
                      color={"secondary"}
                      value={props.waypoint.speed}
                      onChange={onChangeSpeedSlider} />
                  </Box>
                </Grid>
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={4}>
            <IconButton onClick={() => onClickRemove(props.index)}>
              <Delete style={{ color: grey[50] }} />
            </IconButton>
          </Grid>
        </Grid>
        </ExpansionPanelActions>
        }
      </ExpansionPanel>
    </Box>
  );
}

export default WaypointItem;