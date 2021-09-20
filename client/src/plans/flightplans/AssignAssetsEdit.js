import React, { useContext, useEffect, useState } from 'react';

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

import { getFlightplan, getAssignments, updateAssignments } from './FlightplansUtils';
import { getVehicles } from '../../assets/vehicles/VehicleUtils';
import { getMissions } from '../../missions/missions/MissionUtils';
import { AppContext } from '../../context/Context';

const AssignAssetsEdit = (props) => {
  const [ rows, setRows ] = useState([]);
  const [ flightplanName, setFlightplanName ] = useState("");
  const [ fleetId, setFleetId ] = useState("");
  const [ vehicles, setVehicles ] = useState([]);
  const [ missions, setMissions ] = useState([]);
  const { dispatchMessage } = useContext(AppContext);

  useEffect(() => {
    if (props.open) {
      getFlightplan(props.id)
        .then(data => {
          setFlightplanName(data.name);
          setFleetId(data.fleet_id);
          getAssignments(data.fleet_id)
            .then(data => {
              setRows(data.assignments);
            })
            .catch(message => {
              dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
            });
        })
        .catch(message => {
          dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
        });
      getVehicles()
        .then(data => {
          setVehicles(data.vehicles);
        })
        .catch(message => {
          dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
        });
      getMissions()
        .then(data => {
          setMissions(data.missions);
        })
        .catch(message => {
          dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
        });
    }
  }, [ props.open, props.id, setRows, setFlightplanName, setFleetId, setVehicles, setMissions, dispatchMessage ])

  const onClickCancel = () => {
    props.openAssignDetail(props.id);
  }

  const onClickSave = () => {
    let assignments = { assignments: rows };
    updateAssignments(fleetId, assignments)
      .then(ret => {
        dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `Updated assignments for ${flightplanName} successfully` });
        props.openList();
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
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
        .forEach(row => row.vehicle_id = e.target.value);
      return newRows;
    })
  }

  const hancleMissionChange = (id, e) => {
    setRows(rows => {
      let newRows = rows.slice(0, rows.length);
      newRows
        .filter(row => row.id === id)
        .forEach(row => row.mission_id = e.target.value);
      return newRows;
    })
  }

  return (
    <div className={props.classes.funcPanel2}>
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
                          <TableRow key={row.assignment_id}>
                            <TableCell component="th" scope="row">
                              {row.assignment_id}
                            </TableCell>
                            <TableCell>
                              <FormControl>
                                <Select
                                  labelId="Vehicle"
                                  value={row.vehicle_id}
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
                                  value={row.mission_id}
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