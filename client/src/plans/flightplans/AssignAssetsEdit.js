import React, { useEffect, useState } from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  Paper,
  Divider,
  TableContainer,
  Table,
  TableHead,
  TableCell,
  TableBody,
  TableRow,
  FormControl,
  Select,
  MenuItem,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { getAssignments, updateAssignments } from './FlightplansUtils';
import { getVehicles } from '../../assets/vehicles/VehicleUtils';
import { getMissions } from '../../missions/missions/MissionUtils';

const AssignAssetsEdit = (props) => {
  const [ rows, setRows ] = useState([]);
  const [ vehicles, setVehicles ] = useState([]);
  const [ missions, setMissions ] = useState([]);

  useEffect(() => {
    if (props.open) {
      getAssignments(props.id)
        .then(data => {
          setRows(data.assignments);
        })
      getVehicles()
        .then(data => {
          setVehicles(data.vehicles);
        })
      getMissions()
        .then(data => {
          setMissions(data.missions);
        })
    }
  }, [ props.open, props.id ])

  const onClickCancel = () => {
    props.openAssignDetail(props.id);
  }

  const onClickSave = () => {
    let assignments = { assignments: rows };
    updateAssignments(props.id, assignments)
    .then(ret => {
      props.openList();
    });
  }

  const onClickReturn = () => {
    props.openAssignDetail(props.id);
  }

  const hancleVehicleChange = (id, e) => {
    setRows(rows => {
      let newRows = rows.slice(0, rows.length);
      newRows
        .filter(row => row.id === id)
        .forEach(row => row.vehicleId = e.target.value);
      return newRows;
    })
  }

  const hancleMissionChange = (id, e) => {
    setRows(rows => {
      let newRows = rows.slice(0, rows.length);
      newRows
        .filter(row => row.id === id)
        .forEach(row => row.missionId = e.target.value);
      return newRows;
    })
  }

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Edit assignments</Typography>
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
                            <TableCell>
                              <FormControl>
                                <Select
                                  labelId="Vehicle"
                                  value={row.vehicleId}
                                  onChange={e => hancleVehicleChange(row.id, e)}
                                  fullWidth
                                >
                                  <MenuItem
                                    key={""}
                                    value={""}>
                                      {"-"}
                                  </MenuItem>
                                  {vehicles.map(vehicle => (
                                    <MenuItem
                                      key={vehicle.id}
                                      value={vehicle.id}>
                                        {vehicle.name}
                                    </MenuItem>
                                  ))}
                                </Select>
                              </FormControl>
                            </TableCell>
                            <TableCell>
                              <FormControl>
                                <Select
                                  labelId="Mission"
                                  value={row.missionId}
                                  onChange={e => hancleMissionChange(row.id, e)}
                                  fullWidth
                                >
                                  <MenuItem
                                    key={""}
                                    value={""}>
                                      {"-"}
                                  </MenuItem>
                                  {missions.map(mission => (
                                    <MenuItem
                                      key={mission.id}
                                      value={mission.id}>
                                        {mission.name}
                                    </MenuItem>
                                  ))}
                                </Select>
                              </FormControl>
                            </TableCell>
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
                onClick={onClickCancel}>
              Cancel
            </Button>
          </Box>
          <Box px={1}>
            <Button
                className={props.classes.funcButton}
                onClick={onClickSave}>
              Save
            </Button>
          </Box>
        </Box>
      </Box>
    </div>
  );
}

export default AssignAssetsEdit;