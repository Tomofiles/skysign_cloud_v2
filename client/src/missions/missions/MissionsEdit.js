import React, { useState, useEffect, useContext } from 'react';

import {
  Typography,
  Button,
  TextField,
  Grid,
  Box,
  List,
  Paper,
  Divider,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { getMission, updateMission } from './MissionUtils'
import WaypointItem from './WaypointItem';
import { AppContext } from '../../context/Context';

const MissionsEdit = (props) => {
  const { editMission, dispatchEditMission, dispatchEditMode, dispatchMapPosition } = useContext(AppContext);
  const [ missionName, setMissionName ] = useState("");

  useEffect(() => {
    getMission(props.id)
      .then(data => {
        setMissionName(data.name);
        dispatchEditMode({ type: 'MISSION' });
        dispatchEditMission({
          type: 'OPEN',
          mission: data,
        });
        if (data.items.length > 0) {
          dispatchMapPosition({
            type: 'CURRENT',
            longitude: data.items[0].longitude,
            latitude: data.items[0].latitude,
            height: data.takeoffPointGroundHeight + 200,
          })
        }
      })
    return () => {
      dispatchEditMode({ type: 'NONE' });
      dispatchEditMission({ type: "CLEAR"});
    }
  }, [ props.id, setMissionName, dispatchEditMode, dispatchEditMission, dispatchMapPosition ])

  const onClickCancel = () => {
    props.openDetail(props.id);
  }

  const onClickSave = () => {
    updateMission(props.id, editMission)
      .then(ret => {
        props.openList();
      });
  }

  const onClickReturn = () => {
    props.openDetail(props.id);
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
    <div className={props.classes.funcPanel}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Edit Mission</Typography>
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
    </div>
  );
}

export default MissionsEdit;