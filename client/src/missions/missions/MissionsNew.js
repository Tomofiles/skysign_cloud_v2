import React, { useState, useContext, useEffect } from 'react';

import {
  Typography,
  Button,
  TextField,
  Grid,
  Box,
  List,
  Paper,
  Divider
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { createMission } from './MissionUtils'
import WaypointItem from './WaypointItem';
import { AppContext } from '../../context/Context';

const MissionsNew = (props) => {
  const { editMission, dispatchEditMission, dispatchEditMode } = useContext(AppContext);
  const [ missionName, setMissionName ] = useState("");

  useEffect(() => {
    dispatchEditMode({ type: 'MISSION' });
    dispatchEditMission({ type: "CLEAR"});
    return () => {
      dispatchEditMode({ type: 'NONE' });
      dispatchEditMission({ type: "CLEAR"});
    }
  }, [ dispatchEditMode, dispatchEditMission ])

  const onClickSave = () => {
    createMission(editMission)
      .then(ret => {
        props.openList();
      });
  }

  const onClickReturn = () => {
    props.openList();
  }

  const onClickCancel = () => {
    props.openList();
  }

  const changeName = e => {
    setMissionName(e.target.value);
    dispatchEditMission({
      type: 'CHANGE_NAME',
      name: e.target.value,
    });
  }

  const changeRelativeHeight = async (index, height) => {
    dispatchEditMission({
      type: 'CHANGE_RELATIVE_HEIGHT',
      index: index,
      height: height,
    });
  }

  const changeSpeed = async (index, speed) => {
    dispatchEditMission({
      type: 'CHANGE_SPEED',
      index: index,
      speed: speed,
    });
  }

  const removeWaypoint = index => {
    if (editMission.items.length === 1) {
      dispatchEditMission({
        type: 'CHANGE_TAKEOFF_POINT_GROUND_HEIGHT',
        height: undefined,
      });
    }
    dispatchEditMission({
      type: 'REMOVE_WAYPOINT',
      index: index,
    });
  }

  return (
    <>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Create Mission</Typography>
        </Box>
      </Box>
      <Paper className={props.classes.funcPanelEdit}>
        <Box p={3}>
          <Grid container className={props.classes.textLabel}>
            <Grid item xs={12}>
              <Typography>Mission settings</Typography>
              <Divider/>
            </Grid>
            <Grid item xs={12}>
              <Box className={props.classes.textInput}
                  p={1} m={1} borderRadius={7} >
                <TextField
                  label="Name"
                  name="name"
                  value={missionName}
                  onChange={changeName}
                  fullWidth />
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Box  p={1} m={1} borderRadius={7} >
                <Grid container className={props.classes.textLabel}>
                  <Grid item xs={12}>
                    <Typography style={{fontSize: "12px"}}>Takeoff Ground Height</Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography>
                      {editMission.takeoffPointGroundHeight === undefined ?
                        "-"
                      :
                        editMission.takeoffPointGroundHeight} m
                      </Typography>
                  </Grid>
                </Grid>
              </Box>
            </Grid>
          </Grid>
          <Grid item xs={12}>
            <Typography>Waypoints settings</Typography>
            <Divider/>
          </Grid>
          <Grid item xs={12}>
            <List 
              className={props.classes.funcPanelDetails} >
              {editMission.items.length === 0 &&
                <Typography>No Waypoints</Typography>
              }
              {editMission.items.map((waypoint, index) => (
                <WaypointItem
                  key={index}
                  classes={props.classes}
                  index={index}
                  waypoint={waypoint}
                  editable
                  changeRelativeHeight={changeRelativeHeight}
                  changeSpeed={changeSpeed}
                  onClickRemove={removeWaypoint} />
              ))}
            </List>
          </Grid>
        </Box>
        <Box p={3}>
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
                  onClick={onClickSave}>
                Save
              </Button>
            </Box>
          </Box>
        </Box>
      </Paper>
    </>
  );
}

export default MissionsNew;