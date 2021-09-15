import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  List,
  Divider,
  Select,
  Paper,
  FormControl,
  MenuItem,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ArrowDropDown from '@material-ui/icons/ArrowDropDown';
import ArrowDropUp from '@material-ui/icons/ArrowDropUp';
import { grey } from '@material-ui/core/colors';

import { getMission, deleteMission } from './MissionUtils'
import WaypointItem from './WaypointItem';
import { AppContext } from '../../context/Context';

const EDIT_MISSION = "edit_mission"
const DELETE_MISSION = "delete_mission"

const MissionsDetail = (props) => {
  const { editMission, dispatchEditMission, dispatchMapPosition } = useContext(AppContext);
  const [ openAction, setOpenAction ] = useState(false);

  useEffect(() => {
    getMission(props.id)
      .then(data => {
        dispatchEditMission({
          type: 'OPEN',
          mission: data,
        });
        if (data.navigation.waypoints.length > 0) {
          dispatchMapPosition({
            type: 'CURRENT',
            longitude: data.navigation.waypoints[0].longitude,
            latitude: data.navigation.waypoints[0].latitude,
            height: data.navigation.takeoff_point_ground_altitude + 200,
          })
        }
      })
    return () => {
      dispatchEditMission({ type: "CLEAR"});
    }
  }, [ props.id, dispatchEditMission, dispatchMapPosition ])

  useEffect(() => {
    console.log(editMission);
  }, [ editMission ])

  const onClickReturn = () => {
    props.openList();  
  }

  const handleActionChange = e => {
    switch (e.target.value) {
      case EDIT_MISSION:
        props.openEdit(props.id);
        break;
      case DELETE_MISSION:
        deleteMission(props.id)
          .then(data => {
            props.openList();
          })
        break;
      default:
        break;
    }
  };

  const handleActionClose = () => {
    setOpenAction(false);
  };

  const handleActionOpen = () => {
    setOpenAction(true);
  };

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Button onClick={onClickReturn}>
            <ChevronLeftIcon style={{ color: grey[50] }} />
          </Button>
        </Box>
        <Box m={2} style={{display: 'flex', justifyContent: 'space-between'}}>
          <Box>
            <Typography>{editMission.name}</Typography>
          </Box>
          <Box style={{display: 'flex'}}>
            <Box px={1}>
              <FormControl>
                <Button
                    id="openMenu"
                    className={props.classes.funcButton}
                    onClick={handleActionOpen} >
                  Action
                  {!openAction ? (
                    <ArrowDropDown fontSize="small"/>
                  ) : (
                    <ArrowDropUp fontSize="small"/>
                  )}
                </Button>
                <Select
                  onChange={handleActionChange}
                  style={{ display: "none" }}
                  open={openAction}
                  onClose={handleActionClose}
                  value=""
                  MenuProps={{
                    anchorEl: document.getElementById("openMenu"),
                    style: { marginTop: 60 }
                  }}
                >
                  <MenuItem value={EDIT_MISSION}>Edit Mission</MenuItem>
                  <MenuItem value={DELETE_MISSION}>Delete Mission</MenuItem>
                </Select>
              </FormControl>
            </Box>
          </Box>
        </Box>
      </Box>
      <Box pb={2}>
        <Paper className={props.classes.funcPanelEdit}>
          <Box p={3}>
            <Grid container className={props.classes.textLabel}>
              <Grid item xs={12}>
                <Typography>Mission details</Typography>
                <Divider/>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>Name</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography>{editMission.name}</Typography>
                    </Grid>
                  </Grid>
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>Takeoff Ground Height</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography>{editMission.navigation.takeoff_point_ground_altitude} m</Typography>
                    </Grid>
                  </Grid>
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Typography>Waypoints settings</Typography>
                <Divider/>
              </Grid>
              <Grid item xs={12}>
                <List
                  className={props.classes.missionList} >
                  {editMission.navigation.waypoints.length === 0 &&
                    <Typography>No Waypoints</Typography>
                  }
                  {editMission.navigation.waypoints.map((waypoint, index) => (
                    <WaypointItem
                      key={index}
                      classes={props.classes}
                      index={index}
                      waypoint={waypoint} />
                  ))}
                </List>
              </Grid>
            </Grid>
            <Divider/>
          </Box>
        </Paper>
      </Box>
    </div>
  );
}

export default MissionsDetail;