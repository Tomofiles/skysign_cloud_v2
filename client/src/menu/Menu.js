import React from 'react';

import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Divider,
  Card,
  CardMedia,
  Badge,
  Typography,
  Grid,
  Box
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Flight from '@material-ui/icons/Flight';
import Send from '@material-ui/icons/Send';
import Settings from '@material-ui/icons/Settings';
import Games from '@material-ui/icons/Games';

const list = (classes, controlsOpen, missionsOpen, assetsOpen, toggleControls, toggleMissions, toggleAssets) => (
  <div>
    <Card
        className={classes.menuLogoBackground}>
      <Grid container>
        <Grid item xs={12}>
          <CardMedia
            className={classes.menuLogo}
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
      <ListItem button onClick={toggleControls}>
        <ListItemIcon >
          <Grid container className={classes.menuWidthItem} >
            <Grid item xs={12} >
              <Badge color="secondary" variant="dot" invisible={!controlsOpen}>
                <Games style={{ color: grey[50] }} fontSize="large" />
              </Badge>
            </Grid>
            <Grid item xs={12}>
              <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Controls</Typography>
            </Grid>
          </Grid>
        </ListItemIcon>
      </ListItem>
      <ListItem button onClick={toggleMissions}>
        <ListItemIcon>
          <Grid container className={classes.menuWidthItem} >
            <Grid item xs={12}>
              <Badge color="secondary" variant="dot" invisible={!missionsOpen}>
                <Send style={{ color: grey[50] }} fontSize="large" />
              </Badge>
            </Grid>
            <Grid item xs={12}>
              <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Missions</Typography>
            </Grid>
          </Grid>
        </ListItemIcon>
      </ListItem>
      <ListItem button onClick={toggleAssets}>
        <ListItemIcon>
          <Grid container className={classes.menuWidthItem} >
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
          <Grid container className={classes.menuWidthItem} >
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
  </div>
);

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
      {list(
        props.classes,
        props.controlsOpen,
        props.missionsOpen,
        props.assetsOpen,
        props.toggleControls,
        props.toggleMissions,
        props.toggleAssets
      )}
    </Drawer>
  );
}

export default Menu;