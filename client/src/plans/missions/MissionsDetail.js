import React, { useGlobal, useEffect } from 'reactn';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  Grid,
  Box,
  List
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { getMission } from './MissionUtils'
import { Mission } from './MissionHelper';
import WaypointItem from './WaypointItem';

const MissionsDetail = (props) => {
  const [ mission, setMission ] = useGlobal("editMission");

  useEffect(() => {
    setMission(new Mission());
    getMission(props.id)
      .then(data => {
        let newMission = Object.assign(Object.create(Mission.prototype), data);
        setMission(newMission);
      })
  }, [ props.id, setMission ])

  const onClickEdit = () => {
    setMission(new Mission());
    props.openEdit(props.id);
  }

  const onClickReturn = () => {
    setMission(new Mission());
    props.openList();  
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
            <Typography>Detail Mission</Typography>
          </Grid>
          <Grid item xs={12}>
            <Box  p={1} m={1} borderRadius={7} >
              <Grid container className={props.classes.editVehicleInput}>
                <Grid item xs={12}>
                  <Typography style={{fontSize: "12px"}}>Name</Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography>{mission.name}</Typography>
                </Grid>
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={12}>
            <Box  p={1} m={1} borderRadius={7} >
              <Grid container className={props.classes.editVehicleInput}>
                <Grid item xs={12}>
                  <Typography style={{fontSize: "12px"}}>Takeoff Ground Height</Typography>
                </Grid>
                <Grid item xs={12}>
                  <Typography>{mission.takeoffPointGroundHeight} m</Typography>
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
                  waypoint={waypoint} />
              ))}
            </List>
          </Grid>
        </Grid>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button
            className={props.classes.editVehicleButton}
            onClick={onClickEdit}>
          Edit
        </Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default MissionsDetail;