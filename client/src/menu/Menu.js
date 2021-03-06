import React, { useContext, useEffect, useState } from 'react';

import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  Card,
  CardMedia,
  Badge,
  Typography,
  Grid,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Flight from '@material-ui/icons/Flight';
import EventNote from '@material-ui/icons/EventNote';
import Timeline from '@material-ui/icons/Timeline';
import Settings from '@material-ui/icons/Settings';
import Send from '@material-ui/icons/Send';
import MapMode from './MapMode';
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';
import { Description } from '@material-ui/icons';

const Menu = (props) => {
  const { funcMode, dispatchFuncMode } = useContext(AppContext);
  const [ reportsOpen, setReportsOpen ] = useState(false);
  const [ flightsOpen, setFlightsOpen ] = useState(false);
  const [ plansOpen, setPlansOpen ] = useState(false);
  const [ missionsOpen, setMissionsOpen ] = useState(false);
  const [ assetsOpen, setAssetsOpen ] = useState(false);

  useEffect(() => {
    setReportsOpen(funcMode === FUNC_MODE.REPORTS);
    setFlightsOpen(funcMode === FUNC_MODE.FLIGHTS);
    setPlansOpen(funcMode === FUNC_MODE.PLANS);
    setMissionsOpen(funcMode === FUNC_MODE.MISSIONS);
    setAssetsOpen(funcMode === FUNC_MODE.ASSETS);
  }, [ funcMode ])

  const openReports = () => {
    dispatchFuncMode({ type: 'REPORTS' });
  }

  const openFlights = () => {
    dispatchFuncMode({ type: 'FLIGHTS' });
  }

  const openPlans = () => {
    dispatchFuncMode({ type: 'PLANS' });
  }

  const openMissions = () => {
    dispatchFuncMode({ type: 'MISSIONS' });
  }

  const openAssets = () => {
    dispatchFuncMode({ type: 'ASSETS' });
  }

  return (
    <Drawer
      className={props.classes.menu}
      anchor='left'
      variant="permanent"
      classes={{
        paper: props.classes.menuPaper,
      }}
      open={true} >
      <Card
          className={props.classes.menuLogoBackground}>
        <Grid container>
          <Grid item xs={12}>
            <CardMedia
              className={props.classes.menuLogo}
              image="logo_transparent.png"
              title="Skysign Cloud"
            />
          </Grid>
          <Grid item xs={12}>
            <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Skysign Cloud</Typography>
          </Grid>
        </Grid>
      </Card>
      <Box px={1} py={2} />
      <List>
        <MapMode classes={props.classes} />
        <ListItem button onClick={openReports}>
          <ListItemIcon >
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12} >
                <Badge color="secondary" variant="dot" invisible={!reportsOpen}>
                  <Description style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Reports</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button onClick={openFlights}>
          <ListItemIcon >
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12} >
                <Badge color="secondary" variant="dot" invisible={!flightsOpen}>
                  <Send style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Flights</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button onClick={openPlans}>
          <ListItemIcon>
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12}>
                <Badge color="secondary" variant="dot" invisible={!plansOpen}>
                  <EventNote style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Plans</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button onClick={openMissions}>
          <ListItemIcon>
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12}>
                <Badge color="secondary" variant="dot" invisible={!missionsOpen}>
                  <Timeline style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Missions</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button onClick={openAssets}>
          <ListItemIcon>
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12}>
                <Badge color="secondary" variant="dot" invisible={!assetsOpen}>
                  <Flight style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Assets</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button>
          <ListItemIcon>
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12}>
                <Badge color="secondary" variant="dot" invisible>
                  <Settings style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Settings</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
      </List>
    </Drawer>
  );
}

export default Menu;