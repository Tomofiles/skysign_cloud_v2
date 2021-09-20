import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  Paper,
  Divider,
  Select,
  MenuItem,
  FormControl,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ArrowDropDown from '@material-ui/icons/ArrowDropDown';
import ArrowDropUp from '@material-ui/icons/ArrowDropUp';
import { grey } from '@material-ui/core/colors';

import { getFlightplan, executeFlightplan, deleteFlightplan, getAssignments } from './FlightplansUtils';
import { AppContext } from '../../context/Context';

const EDIT_FLIGHTPLAN = "edit_flightplan"
const DELETE_FLIGHTPLAN = "delete_flightplan"
const CHANGE_NUMBER_OF_VEHICLES = "change_number_of_vehicles"
const ASSGN_ASSETS = "assign_assets"

const FlightplansDetail = (props) => {
  const { dispatchFuncMode, dispatchMessage } = useContext(AppContext);
  const [ flightplan, setFlightplan ] = useState({id: "-", name: "-", description: "-", fleet_id: "-"});
  const [ numberOfVehicles, setNumberOfVehicles ] = useState("-");
  const [ openAction, setOpenAction ] = useState(false);

  useEffect(() => {
    getFlightplan(props.id)
      .then(data => {
        setFlightplan(data);
        getAssignments(data.fleet_id)
          .then(data => {
            setNumberOfVehicles(data.assignments.length);
          })
          .catch(message => {
            dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
          });
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }, [ props.id, setFlightplan, setNumberOfVehicles, dispatchMessage ])

  const onClickReturn = () => {
    props.openList();  
  }

  const onClickFlight = () => {
    executeFlightplan(props.id)
      .then(ret => {
        dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `Executed ${flightplan.name} successfully` });
        dispatchFuncMode({ type: 'FLIGHTS' });
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }

  const handleActionChange = e => {
    switch (e.target.value) {
      case CHANGE_NUMBER_OF_VEHICLES:
        props.openChangeNumberOfVehicles(props.id);
        break;
      case ASSGN_ASSETS:
        props.openAssignDetail(props.id);
        break;
      case EDIT_FLIGHTPLAN:
        props.openEdit(props.id);
        break;
      case DELETE_FLIGHTPLAN:
        deleteFlightplan(props.id)
          .then(data => {
            dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `You have successfully deleted ${flightplan.name}` });
            props.openList();
          })
          .catch(message => {
            dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
          });
        break;
      default:
        break;
    }
  };

  const handleActionClose = () => {
    setOpenAction(false);
  };

  const handleActionOpen = () => {
    setOpenAction(true);
  };

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Button onClick={onClickReturn}>
            <ChevronLeftIcon style={{ color: grey[50] }} />
          </Button>
        </Box>
        <Box m={2} style={{display: 'flex', justifyContent: 'space-between'}}>
          <Box>
            <Typography>{flightplan.name}</Typography>
          </Box>
          <Box style={{display: 'flex'}}>
            <Box px={1}>
              <Button className={props.classes.funcImportantButton} onClick={onClickFlight}>Flight</Button>
            </Box>
            <Box px={1}>
              <FormControl>
                <Button
                    id="openMenu"
                    className={props.classes.funcButton}
                    onClick={handleActionOpen} >
                  Action
                  {!openAction ? (
                    <ArrowDropDown fontSize="small"/>
                  ) : (
                    <ArrowDropUp fontSize="small"/>
                  )}
                </Button>
                <Select
                  onChange={handleActionChange}
                  style={{ display: "none" }}
                  open={openAction}
                  onClose={handleActionClose}
                  value=""
                  MenuProps={{
                    anchorEl: document.getElementById("openMenu"),
                    style: { marginTop: 60 }
                  }}
                >
                  <MenuItem value={EDIT_FLIGHTPLAN}>Edit Flightplan</MenuItem>
                  <MenuItem value={CHANGE_NUMBER_OF_VEHICLES}>Change Number Of Vehicles</MenuItem>
                  <MenuItem value={ASSGN_ASSETS}>Assign Assets</MenuItem>
                  <MenuItem value={DELETE_FLIGHTPLAN}>Delete Flightplan</MenuItem>
                </Select>
              </FormControl>
            </Box>
          </Box>
        </Box>
      </Box>
      <Box pb={2}>
        <Paper className={props.classes.funcPanelEdit}>
          <Box p={3}>
            <Grid container className={props.classes.textLabel}>
              <Grid item xs={12}>
                <Typography>Basic configuration</Typography>
                <Divider/>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>Name</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography>{flightplan.name}</Typography>
                    </Grid>
                  </Grid>
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>Description</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <div>
                        {flightplan.description.split("\n").map((t, i) => {
                          return <div key={i}>{t}</div>
                        })}
                      </div>
                    </Grid>
                  </Grid>
                </Box>
              </Grid>
              <Grid item xs={12}>
                <Typography>Fleet formation configuration</Typography>
                <Divider/>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <Grid container className={props.classes.textLabel}>
                    <Grid item xs={12}>
                      <Typography style={{fontSize: "12px"}}>Number of vehicles</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography>{numberOfVehicles}</Typography>
                    </Grid>
                  </Grid>
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Paper>
      </Box>
    </div>
  );
}

export default FlightplansDetail;