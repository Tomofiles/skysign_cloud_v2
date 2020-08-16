import React from 'react';

import {
  Typography,
  Box,
  Paper,
  List,
  ListItem,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Grid,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Send from '@material-ui/icons/Send';
import Flight from '@material-ui/icons/Flight';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';

const StagingList = (props) => {

  const onClickRow = id => {
    props.selectRow(id);
    props.openEdit();
  }

  const onClickControl = (id, isControlled) => {
    props.changeControl(id, isControlled);
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
            <Box key={row.id} pb={1} >
              <ListItem button selected={row.selected} component={Paper} className={props.classes.myVehiclePaper}>
                <Grid container>
                  <Grid item xs={10}>
                    <Box onClick={() => onClickRow(row.id)}>
                      <Grid container>
                        <Grid item xs={2}>
                          <Flight style={{ color: grey[50] }} fontSize="small" />
                        </Grid>
                        <Grid item xs={10}>
                          <Typography >{row.vehicleName}</Typography>
                        </Grid>
                        <Grid item xs={2}>
                          <Send style={{ color: grey[50] }} fontSize="small" />
                        </Grid>
                        <Grid item xs={10}>
                          <Typography >{row.missionName}</Typography>
                        </Grid>
                      </Grid>
                    </Box>
                  </Grid>
                  <Grid item xs={2}>
                    <Box px={1.5} py={1.5} onClick={() => onClickControl(row.id, row.isControlled)}>
                      {row.isControlled ?
                        <Visibility style={{ color: grey[50] }} fontSize="small" />
                      :
                        <VisibilityOff style={{ color: grey[50] }} fontSize="small" />}
                    </Box>
                  </Grid>
                </Grid>
              </ListItem>
            </Box>
          ))}
        </List>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
      </ExpansionPanelActions>
    </div>
  );
}

export default StagingList;