import React, { useContext, useEffect } from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  List,
  Divider,
  Paper
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { getMission, deleteMission } from './MissionUtils'
import WaypointItem from './WaypointItem';
import { AppContext } from '../../context/Context';

const MissionsDetail = (props) => {
  const { editMission, dispatchEditMission, dispatchMapPosition } = useContext(AppContext);

  useEffect(() => {
    getMission(props.id)
      .then(data => {
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
      dispatchEditMission({ type: "CLEAR"});
    }
  }, [ props.id, dispatchEditMission, dispatchMapPosition ])

  const onClickEdit = () => {
    props.openEdit(props.id);
  }

  const onClickDelete = () => {
    deleteMission(props.id)
      .then(data => {
        props.openList();
      })
  }

  const onClickReturn = () => {
    props.openList();  
  }

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>{editMission.name}</Typography>
        </Box>
      </Box>
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
                    <Typography>{editMission.takeoffPointGroundHeight} m</Typography>
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
                {editMission.items.length === 0 &&
                  <Typography>No Waypoints</Typography>
                }
                {editMission.items.map((waypoint, index) => (
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
        <Box p={3}>
          <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
            <Box px={1}>
              <Button 
                  className={props.classes.funcButton}
                  onClick={onClickDelete}>
                Delete
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
      </Paper>
    </div>
  );
}

export default MissionsDetail;