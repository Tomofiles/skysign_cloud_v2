import React, { useState } from 'react';

import {
  Typography,
  Box,
  Paper,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import EventNote from '@material-ui/icons/EventNote';

const FlightplansList = (props) => {
  const rows = useState([
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
  ])[0];

  const onClickNew = () => {
    props.openNew();
  }

  const onSelect = (id) => {
    props.openDetail(id);
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <List 
          className={props.classes.funcPanelDetails} >
          {rows.length === 0 &&
            <Typography>No Flightplans</Typography>
          }
          {rows.map((row) => (
            <Box key={row.id} pb={1} onClick={() => onSelect(row.id)} >
              <ListItem button component={Paper} className={props.classes.funcListItem}>
                <ListItemIcon>
                  <EventNote style={{ color: grey[50] }} />
                </ListItemIcon>
                <ListItemText >{row.name}</ListItemText>
              </ListItem>
            </Box>
          ))}
        </List>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button className={props.classes.funcButton} onClick={onClickNew}>New</Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default FlightplansList;