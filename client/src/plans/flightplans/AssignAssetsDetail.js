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
  const [ vehicles, setVehicles ] = useState([]);
  const [ missions, setMissions ] = useState([]);

  useEffect(() => {
    if (props.open) {
      getAssignments(props.id)
        .then(data => {
          setRows(data.assignments);
        })
      // getVehicles()
      //   .then(data => {
      //     setVehicles(data.vehicles);
      //   })
      // getMissions()
      //   .then(data => {
      //     setMissions(data.missions);
      //   })
    }
  }, [ props.open, props.id ])

  const onClickEdit = () => {
    props.openAssignEdit(props.id);
  }

  const onClickReturn = () => {
    props.openDetail(props.id);  
  }

  return (
    <div className={props.classes.funcPanel}>
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
                          <TableRow key={row.assignmentId}>
                            <TableCell component="th" scope="row">
                              {row.assignmentId}
                            </TableCell>
                            <TableCell>{row.vehicleId}</TableCell>
                            <TableCell>{row.missionId}</TableCell>
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