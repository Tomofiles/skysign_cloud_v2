import React from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  Grid,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Paper,
  Box,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

const AssignDetail = (props) => {
  const rows = [
    {
      fleet: "vehicle -- 1",
      vehicle: "PX4 gazebo",
    },
    {
      fleet: "vehicle -- 2",
      vehicle: "-",
    },
    {
      fleet: "vehicle -- 3",
      vehicle: "-",
    },
    {
      fleet: "vehicle -- 3",
      vehicle: "-",
    },
    {
      fleet: "vehicle -- 3",
      vehicle: "-",
    },
    {
      fleet: "vehicle -- 3",
      vehicle: "-",
    },
    {
      fleet: "vehicle -- 3",
      vehicle: "-",
    },
    {
      fleet: "vehicle -- 3",
      vehicle: "-",
    },
    {
      fleet: "vehicle -- 3",
      vehicle: "-",
    }
  ]

  const onClickEdit = () => {
    props.openAssignEdit(props.id);
  }

  const onClickReturn = () => {
    props.openDetail(props.id);  
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <Grid container className={props.classes.textLabel}>
          <Grid item xs={12}>
            <Button onClick={onClickReturn}>
              <ChevronLeftIcon style={{ color: grey[50] }} />
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Typography>Detail Flightplan Assignment</Typography>
          </Grid>
          <Grid item xs={12}>
            <Box p={2}>
              <TableContainer component={Paper} style={{maxHeight: '300px'}}>
                <Table aria-label="simple table">
                  <TableHead>
                    <TableCell>Fleet</TableCell>
                    <TableCell>Vehicle</TableCell>
                  </TableHead>
                  <TableBody>
                    {rows.map((row) => (
                      <TableRow key={row.fleet}>
                        <TableCell component="th" scope="row">
                          {row.fleet}
                        </TableCell>
                        <TableCell>{row.vehicle}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </Box>
          </Grid>
        </Grid>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button
            className={props.classes.funcButton}
            onClick={onClickEdit}>
          Edit
        </Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default AssignDetail;