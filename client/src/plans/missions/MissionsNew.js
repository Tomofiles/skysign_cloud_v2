import React, { useState, useGlobal } from 'reactn';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  TextField,
  Grid,
  Box,
  List
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { createMission } from './MissionUtils'
import WaypointItem from './WaypointItem';
import { Mission } from './MissionHelper';
import { useEffect } from 'react';
import { EDIT_MODE } from '../../App';

const MissionsNew = (props) => {
  const [ mission, setMission ] = useGlobal("editMission");
  const setEditMode = useGlobal("editMode")[1];
  const [ missionName, setMissionName ] = useState("");

  useEffect(() => {
    setEditMode(EDIT_MODE.MISSION);
    let newMission = new Mission();
    setMission(newMission);
  }, [ props, setMission, setEditMode ])

  const onClickSave = () => {
    createMission(mission)
      .then(ret => {
        setEditMode(EDIT_MODE.NONE);
        setMission(new Mission());
        props.openList();
      });
  }

  const onClickReturn = () => {
    setEditMode(EDIT_MODE.NONE);
    setMission(new Mission());
    props.openList();
  }

  const changeName = e => {
    setMissionName(e.target.value);

    let newMission = Object.assign(Object.create(Mission.prototype), mission);
    newMission.nameMission(e.target.value);

    setMission(newMission);
  }

  const changeRelativeHeight = async (index, height) => {
    let newMission = Object.assign(Object.create(Mission.prototype), mission);
    await newMission.changeRelativeHeight(index, height);

    setMission(newMission);
  }

  const changeSpeed = async (index, speed) => {
    let newMission = Object.assign(Object.create(Mission.prototype), mission);
    await newMission.changeSpeed(index, speed);

    setMission(newMission);
  }

  const removeWaypoint = index => {
    let newMission = Object.assign(Object.create(Mission.prototype), mission);
    newMission.removeWaypoint(index);

    setMission(newMission);
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <Grid container className={props.classes.editVehicleInput}>
          <Grid item xs={12}>
            <Button onClick={onClickReturn}>
              <ChevronLeftIcon style={{ color: grey[50] }} />
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Typography>New Mission</Typography>
          </Grid>
          <Grid item xs={12}>
            <Box className={props.classes.editVehicleInputText}
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
              <Grid container className={props.classes.editVehicleInput}>
                <Grid item xs={12}>
                  <Typography style={{fontSize: "12px"}}>Takeoff Ground Height</Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography>
                    {mission.takeoffPointGroundHeight === undefined ?
                      "-"
                    :
                      mission.takeoffPointGroundHeight} m
                    </Typography>
                </Grid>
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={12}>
            <Typography>Waypoints</Typography>
          </Grid>
          <Grid item xs={12}>
            <List 
              className={props.classes.myVehicleList} >
              {mission.items.length === 0 &&
                <Typography>No Waypoints</Typography>
              }
              {mission.items.map((waypoint, index) => (
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
        </Grid>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button
            className={props.classes.editVehicleButton}
            onClick={onClickSave}>
          Save
        </Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default MissionsNew;