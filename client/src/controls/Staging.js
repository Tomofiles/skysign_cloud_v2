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
  Button,
  ExpansionPanel,
  ExpansionPanelSummary
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Flight from '@material-ui/icons/Flight';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

const Staging = (props) => {
  const [rows, setRows] = useState([]);

  return (
    <ExpansionPanel
        className={props.classes.myVehicleRoot}
        defaultExpanded>
      <ExpansionPanelSummary
        expandIcon={<ExpandMoreIcon style={{ color: grey[50] }} />}
        aria-controls="panel1a-content"
        id="panel1a-header"
        className={props.classes.myVehicleSummary}
      >
        <Typography>Staging</Typography>
      </ExpansionPanelSummary>
      <ExpansionPanelDetails>
        <List 
          className={props.classes.myVehicleList} >
          {rows.length === 0 &&
            <Typography>No Vehicles</Typography>
          }
          {rows.map((row) => (
            <Box key={row.id} pb={1} onClick={() => {}} >
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
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Add</Button>
        <Button className={props.classes.myVehicleButton} onClick={() => {}}>Delete</Button>
      </ExpansionPanelActions>
    </ExpansionPanel>
  );
}

export default Staging;