import React, { useContext, useEffect, useLayoutEffect, useState } from 'react';

import {
  Button,
  Box,
  Paper,
  Typography,
  Divider,
  Grid,
  List,
  ListItem,
  IconButton,
  Stepper,
  Step,
  StepButton,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import Visibility from '@material-ui/icons/Visibility';
import { grey } from '@material-ui/core/colors';

import { AppContext } from '../../context/Context';

function getSteps() {
  return ['UPLOAD', 'TAKEOFF', 'START', 'LAND', 'END'];
}

const FlightOperation = (props) => {
  const { dispatchOperationMode } = useContext(AppContext);
  const [ listsize, setListsize ] = useState("0vh");
  const fleet = ["","","","",""];
  const steps = getSteps();

  useEffect(() => {
    dispatchOperationMode({ type: 'OPERATION' });
    return () => {
      dispatchOperationMode({ type: 'NONE' });
    }
  }, [ props.id, dispatchOperationMode ])

  useLayoutEffect(() => {
    // 仮画面サイズ調整
    let listsize = window.innerHeight - (32 + 32 + 24 + 36 + 16 + 16 + 24 + 48);
    setListsize(listsize + "px");
  }, [ setListsize ])

  const onClickReturn = () => {
    props.openList();  
  }

  const onClickJump = () => {

  }

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
            <Typography>flightplan name</Typography>
          </Box>
        </Box>
      </Box>
      <Box pb={4}>
        <List style={{overflowY: "auto", height: listsize}}>
        {fleet.map(() => (
          <Box pb={1}>
          <ListItem component={Paper} className={props.classes.funcPanelEdit}>
          {/* <Box pb={1}>
            <Paper className={props.classes.funcPanelEdit}> */}
              <Box>
                <Grid container className={props.classes.textLabel}>
                  <Grid item xs={12}>
                    <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
                      <IconButton size="small" onClick={onClickJump}>
                        <Visibility style={{ color: grey[50] }} />
                      </IconButton>
                    </Box>
                  </Grid>
                  <Grid item xs={6}>
                    <Box>
                      <Typography>Vehicle</Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={6}>
                    <Box>
                      <Typography>vehicle name -- 1</Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={12}>
                    <Divider/>
                  </Grid>
                  <Grid item xs={12}>
                    <Box pt={1}>
                      <Typography>Telemetry</Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={3}>
                    <Box p={1}>
                      <Grid container className={props.classes.textLabel}>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }} >lat (deg)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>alt (m)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>rel alt (m)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>speed (m/s)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>state</Typography>
                        </Grid>
                      </Grid>
                    </Box>
                  </Grid>
                  <Grid item xs={3}>
                    <Box p={1}>
                      <Grid container className={props.classes.textLabel}>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>34.566894</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>12.9554</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>3.5998</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>1.6998</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>LANDING</Typography>
                        </Grid>
                      </Grid>
                    </Box>
                  </Grid>
                  <Grid item xs={3}>
                    <Box p={1}>
                      <Grid container className={props.classes.textLabel}>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>lon (deg)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>heading (deg)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>pitch (deg)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>roll (deg)</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ textAlign: "right", fontSize: "13px", whiteSpace: "nowrap" }}>arm</Typography>
                        </Grid>
                      </Grid>
                    </Box>
                  </Grid>
                  <Grid item xs={3}>
                    <Box p={1}>
                      <Grid container className={props.classes.textLabel}>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>142.966587</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>45.955</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>0.158</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>0.002</Typography>
                        </Grid>
                        <Grid item xs={12}>
                          <Typography style={{ fontSize: "13px", whiteSpace: "nowrap" }}>disarmed</Typography>
                        </Grid>
                      </Grid>
                    </Box>
                  </Grid>
                  <Grid item xs={12}>
                    <Divider/>
                  </Grid>
                  <Grid item xs={6}>
                    <Box py={1}>
                      <Typography>Mission</Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={6}>
                    <Box py={1}>
                      <Typography>mission name -- 1</Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={12}>
                    <Divider/>
                  </Grid>
                  <Grid item xs={12}>
                    <Box pt={1}>
                      <Typography>Step</Typography>
                    </Box>
                  </Grid>
                  <Grid item xs={12}>
                    <Stepper alternativeLabel>
                      {steps.map((label) => (
                        <Step key={label}>
                          <StepButton>
                            {label}
                          </StepButton>
                        </Step>
                      ))}
                    </Stepper>
                  </Grid>
                </Grid>
              </Box>
            {/* </Paper>
          </Box> */}
          </ListItem>
          </Box>
        ))}
        </List>
      </Box>
    </div>
  );
}

export default FlightOperation;