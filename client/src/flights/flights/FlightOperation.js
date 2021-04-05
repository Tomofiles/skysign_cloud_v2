import React, { useContext, useEffect, useLayoutEffect, useState } from 'react';

import {
  Button,
  Box,
  Typography,
  List,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

import { AppContext } from '../../context/Context';
import FlightOperationAssignment from './FlightOperationAssignment';
import { completeFlight } from './FlightUtils';

const FlightOperation = (props) => {
  const { dispatchOperationMode, flightplan, assignments, dispatchFlightoperation, dispatchFuncMode } = useContext(AppContext);
  const [ listsize, setListsize ] = useState("0vh");

  useEffect(() => {
    dispatchFlightoperation({ type: 'ID', id: props.id });
    dispatchOperationMode({ type: 'OPERATION' });
    return () => {
      dispatchFlightoperation({ type: 'NONE' });
      dispatchOperationMode({ type: 'NONE' });
    }
  }, [ props.id, dispatchOperationMode, dispatchFlightoperation ])

  useLayoutEffect(() => {
    // 仮画面サイズ調整
    let listsize = window.innerHeight - (32 + 32 + 24 + 36 + 16 + 16 + 24 + 48);
    setListsize(listsize + "px");
  }, [ setListsize ])

  const onClickReturn = () => {
    props.openList();  
  }

  const onClickComplete = () => {
    completeFlight(props.id)
      .then(data => dispatchFuncMode({ type: 'REPORTS' }));
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
            <Typography>{flightplan.name}</Typography>
          </Box>
          <Box px={1}>
            <Button className={props.classes.funcImportantButton} onClick={onClickComplete}>Complete</Button>
          </Box>
        </Box>
      </Box>
      <Box pb={4}>
        <List style={{overflowY: "auto", height: listsize}}>
          {assignments.map(assignment => (
            <FlightOperationAssignment
              key={assignment.id}
              classes={props.classes}
              vehicleId={assignment.vehicle_id}
              missionId={assignment.mission_id} />
          ))}
        </List>
      </Box>
    </div>
  );
}

export default FlightOperation;