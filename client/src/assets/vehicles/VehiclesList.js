import React, { useEffect, useState } from 'react';

import axios from 'axios';

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

async function getVehicles() {
  try {
    const res = await axios
      .get('/api/v1/vehicles', {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

const VehiclesList = (props) => {
  const [rows, setRows] = useState([]);

  useEffect(() => {
    getVehicles()
      .then(data => {
        setRows(data.vehicles);
      })
  }, [])

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
            <Typography>No Vehicles</Typography>
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

export default VehiclesList;