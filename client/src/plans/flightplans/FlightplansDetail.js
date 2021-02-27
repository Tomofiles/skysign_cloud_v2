import React, { useEffect, useState } from 'react';

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

import { getFlightplan, deleteFlightplan, getAssignments } from './FlightplansUtils';

const CHANGE_NUMBER_OF_VEHICLES = "change_number_of_vehicles"
const ASSGN_ASSETS = "assign_assets"

const FlightplansDetail = (props) => {
  const [ flightplan, setFlightplan ] = useState({id: "-", name: "-", description: "-"});
  const [ numberOfVehicles, setNumberOfVehicles ] = useState("-");
  const [ openAction, setOpenAction ] = useState(false);

  useEffect(() => {
    getFlightplan(props.id)
      .then(data => {
        setFlightplan(data);
      })
    getAssignments(props.id)
      .then(data => {
        setNumberOfVehicles(data.assignments.length);
      })
  }, [ props.id ])

  const onClickDelete = () => {
    deleteFlightplan(props.id)
      .then(data => {
        props.openList();
      })
  }

  const onClickEdit = () => {
    props.openEdit(props.id);
  }

  const onClickReturn = () => {
    props.openList();  
  }

  const handleActionChange = e => {
    switch (e.target.value) {
      case CHANGE_NUMBER_OF_VEHICLES:
        props.openChangeNumberOfVehicles(props.id);
        break;
      case ASSGN_ASSETS:
        props.openAssignDetail(props.id);
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
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>{flightplan.name}</Typography>
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
      <Box>
        <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
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
                <MenuItem value={CHANGE_NUMBER_OF_VEHICLES}>Change Number Of Vehicles</MenuItem>
                <MenuItem value={ASSGN_ASSETS}>Assign Assets</MenuItem>
              </Select>
            </FormControl>
          </Box>
          <Box px={1}>
            <Button
                className={props.classes.funcButton}
                onClick={onClickDelete}>
              Delete
            </Button>
          </Box>
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

export default FlightplansDetail;