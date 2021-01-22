import React from 'react';

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
import Send from '@material-ui/icons/Send';
import Archive from '@material-ui/icons/Archive';
import Settings from '@material-ui/icons/Settings';
import Games from '@material-ui/icons/Games';
import MapMode from './MapMode';

const Menu = (props) => {
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
        <ListItem button onClick={props.toggleControls}>
          <ListItemIcon >
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12} >
                <Badge color="secondary" variant="dot" invisible={!props.controlsOpen}>
                  <Games style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Controls</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button onClick={props.togglePlans}>
          <ListItemIcon>
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12}>
                <Badge color="secondary" variant="dot" invisible={!props.plansOpen}>
                  <Send style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Plans</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button onClick={props.toggleActual}>
          <ListItemIcon>
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12}>
                <Badge color="secondary" variant="dot" invisible={!props.actualOpen}>
                  <Archive style={{ color: grey[50] }} fontSize="large" />
                </Badge>
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Actual</Typography>
              </Grid>
            </Grid>
          </ListItemIcon>
        </ListItem>
        <ListItem button onClick={props.toggleAssets}>
          <ListItemIcon>
            <Grid container className={props.classes.menuItem} >
              <Grid item xs={12}>
                <Badge color="secondary" variant="dot" invisible={!props.assetsOpen}>
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