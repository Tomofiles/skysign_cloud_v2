import React from 'react';

import {
  Typography,
  Box,
  Paper,
  List,
  ListItem,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  Grid,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Send from '@material-ui/icons/Send';
import Flight from '@material-ui/icons/Flight';

const StagingList = (props) => {
  const onClickNew = () => {
    props.openNew();
  }

  const onClickRow = (id) => {
    props.toggleSelectRow(id);
  }

  const onClickRemove = () => {
    props.removeRows();
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <List 
          className={props.classes.myVehicleList} >
          {props.rows.length === 0 &&
            <Typography>No Staging</Typography>
          }
          {props.rows.map((row) => (
            <Box key={row.id} pb={1} onClick={() => onClickRow(row.id)} >
              <ListItem button selected={row.selected} component={Paper} className={props.classes.myVehiclePaper}>
                <Grid container>
                  <Grid item  xs={2}>
                    <Send style={{ color: grey[50] }} fontSize="small" />
                  </Grid>
                  <Grid item xs={10}>
                    <Typography >{row.mission}</Typography>
                  </Grid>
                  <Grid item xs={2}>
                    <Flight style={{ color: grey[50] }} fontSize="small" />
                  </Grid>
                  <Grid item xs={10}>
                    <Typography >{row.vehicleName}</Typography>
                  </Grid>
                </Grid>
              </ListItem>
            </Box>
          ))}
        </List>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button className={props.classes.myVehicleButton} onClick={onClickNew}>Add</Button>
        <Button className={props.classes.myVehicleButton} onClick={onClickRemove}>Remove</Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default StagingList;