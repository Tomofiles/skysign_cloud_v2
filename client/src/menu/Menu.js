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
  Badge
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Flight from '@material-ui/icons/Flight';
import Send from '@material-ui/icons/Send';
import Settings from '@material-ui/icons/Settings';

const list = (classes, missionOpen, assetsOpen, toggleMissions, toggleAssets) => (
  <div>
    <Card
        className={classes.menuLogoBackground}>
      <CardMedia
        className={classes.menuLogo}
        image="logo_transparent.png"
        title="Skysign Cloud"
      />
    </Card>
    <Divider />
    <List>
      <ListItem button onClick={toggleMissions}>
        <ListItemIcon>
          <Badge color="secondary" variant="dot" invisible={!missionOpen}>
            <Send style={{ color: grey[50] }} fontSize="large" />
          </Badge>
        </ListItemIcon>
        <ListItemText >Missions</ListItemText>
      </ListItem>
      <ListItem button onClick={toggleAssets}>
        <ListItemIcon>
          <Badge color="secondary" variant="dot" invisible={!assetsOpen}>
            <Flight style={{ color: grey[50] }} fontSize="large" />
          </Badge>
        </ListItemIcon>
        <ListItemText >Assets</ListItemText>
      </ListItem>
      <ListItem button>
        <ListItemIcon><Settings style={{ color: grey[50] }} fontSize="large" /></ListItemIcon>
        <ListItemText >Settings</ListItemText>
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
      {list(props.classes, props.missionOpen, props.assetsOpen, props.toggleMissions, props.toggleAssets)}
    </Drawer>
  );
}

export default Menu;