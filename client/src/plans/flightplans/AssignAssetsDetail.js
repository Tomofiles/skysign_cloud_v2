import React, { useEffect, useState } from 'react';

import {
  Typography,
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
  Divider,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { getAssignments } from './FlightplansUtils';
import { getVehicles } from '../../assets/vehicles/VehicleUtils';
import { getMissions } from '../../missions/missions/MissionUtils';

const AssignAssetsDetail = (props) => {
  const [ rows, setRows ] = useState([]);

  useEffect(() => {
    if (props.open) {
      Promise.all([
        getVehicles(),
        getMissions(),
      ]).then(vm => {
        getAssignments(props.id)
          .then(data => {
            let rows = data.assignments.slice(0, data.assignments.length);
            rows.forEach(row => {
              let vehicle = vm[0].vehicles.find(v => v.id === row.vehicle_id);
              let mission = vm[1].missions.find(m => m.id === row.mission_id);
              row.vehicleName = vehicle === undefined ? "-" : vehicle.name;
              row.missionName = mission === undefined ? "-" : mission.name;
            });
            setRows(data.assignments);
          })
      })
    }
  }, [ props.open, props.id ])

  const onClickEdit = () => {
    props.openAssignEdit(props.id);
  }

  const onClickReturn = () => {
    props.openDetail(props.id);  
  }

  return (
    <div className={props.classes.funcPanel2}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Assign Assets</Typography>
        </Box>
      </Box>
      <Box pb={2}>
        <Paper className={props.classes.funcPanelEdit}>
          <Box p={3}>
            <Grid container className={props.classes.textLabel}>
              <Grid item xs={12}>
                <Typography>Fleet formation</Typography>
                <Divider/>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <TableContainer component={Paper} style={{maxHeight: '300px'}}>
                    <Table aria-label="simple table" stickyHeader>
                      <TableHead>
                        <TableRow>
                            <TableCell>Fleet</TableCell>
                            <TableCell>Vehicle</TableCell>
                            <TableCell>Mission</TableCell>
                          </TableRow>
                      </TableHead>
                      <TableBody>
                        {rows.map((row) => (
                          <TableRow key={row.assignment_id}>
                            <TableCell component="th" scope="row">
                              {row.assignment_id}
                            </TableCell>
                            <TableCell>{row.vehicleName}</TableCell>
                            <TableCell>{row.missionName}</TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Paper>
      </Box>
      <Box>
        <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
          <Box px={1}>
            <Button
                className={props.classes.funcButton}
                onClick={onClickEdit}>
              Edit
            </Button>
          </Box>
        </Box>
      </Box>
    </div>
  );
}

export default AssignAssetsDetail;