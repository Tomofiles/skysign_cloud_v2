import React, { useState, useEffect, useGlobal } from 'reactn';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  TextField,
  Grid,
  Box,
  List,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { getMission, updateMission, deleteMission } from './MissionUtils'
import { Mission } from './MissionHelper';
import WaypointItem from './WaypointItem';
import { EDIT_MODE } from '../../App';

const MissionsEdit = (props) => {
  const [ mission, setMission ] = useGlobal("editMission");
  const setEditMode = useGlobal("editMode")[1];
  const [ missionName, setMissionName ] = useState("");

  useEffect(() => {
    setMission(new Mission());
    getMission(props.id)
      .then(data => {
        setMissionName(data.name);
        setEditMode(EDIT_MODE.MISSION);
        let newMission = Object.assign(Object.create(Mission.prototype), data);
        setMission(newMission);
      })
  }, [ props.id, setMission, setMissionName, setEditMode ])

  const onClickCancel = () => {
    setEditMode(EDIT_MODE.NONE);
    setMission(new Mission());
    props.openDetail(props.id);
  }

  const onClickSave = () => {
    updateMission(props.id, mission)
      .then(ret => {
        setEditMode(EDIT_MODE.NONE);
        setMission(new Mission());
        props.openList();
      });
  }

  const onClickDelete = () => {
    deleteMission(props.id)
      .then(data => {
        setEditMode(EDIT_MODE.NONE);
        setMission(new Mission());
        props.openList();
      })
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
            <Typography>Edit Mission</Typography>
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
            onClick={onClickCancel}>
          Cancel
        </Button>
        <Button 
            className={props.classes.editVehicleButton}
            onClick={onClickDelete}>
          Delete
        </Button>
        <Button
            className={props.classes.editVehicleButton}
            onClick={onClickSave}>
          Save
        </Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default MissionsEdit;