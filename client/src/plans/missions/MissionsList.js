import React, { useEffect, useState } from 'react';

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
import Flight from '@material-ui/icons/Flight';

import { getMissions } from './MissionUtils'

const MissionsList = (props) => {
  const [rows, setRows] = useState([]);

  useEffect(() => {
    if (props.open) {
      getMissions()
        .then(data => {
          setRows(data.missions);
        })
    }
  }, [props.open])

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
          className={props.classes.myVehicleList} >
          {rows.length === 0 &&
            <Typography>No Missions</Typography>
          }
          {rows.map((row) => (
            <Box key={row.id} pb={1} onClick={() => onSelect(row.id)} >
              <ListItem button component={Paper} className={props.classes.myVehiclePaper}>
                <ListItemIcon>
                  <Flight style={{ color: grey[50] }} />
                </ListItemIcon>
                <ListItemText >{row.name}</ListItemText>
              </ListItem>
            </Box>
          ))}
        </List>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button className={props.classes.myVehicleButton} onClick={onClickNew}>New</Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default MissionsList;