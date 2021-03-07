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
import { getFlight } from './FlightUtils';
import { getFlightplan, getAssignments } from '../../plans/flightplans/FlightplansUtils'
import FlightOperationAssignment from './FlightOperationAssignment';

const FlightOperation = (props) => {
  const { dispatchOperationMode, assignments, dispatchAssignments } = useContext(AppContext);
  const [ listsize, setListsize ] = useState("0vh");
  const [ flightplanName, setFlightplanName ] = useState("-");

  useEffect(() => {
    getFlight(props.id)
      .then(data => {
        getFlightplan(data.flightplan_id)
          .then(data => {
            setFlightplanName(data.name);
          })
        getAssignments(data.flightplan_id)
          .then(data => {
            dispatchAssignments({ type: 'ROWS', rows: data.assignments });
          })
      })
    dispatchOperationMode({ type: 'OPERATION' });
    return () => {
      dispatchAssignments({ type: 'NONE' });
      dispatchOperationMode({ type: 'NONE' });
    }
  }, [ props.id, dispatchOperationMode, dispatchAssignments ])

  useLayoutEffect(() => {
    // 仮画面サイズ調整
    let listsize = window.innerHeight - (32 + 32 + 24 + 36 + 16 + 16 + 24 + 48);
    setListsize(listsize + "px");
  }, [ setListsize ])

  const onClickReturn = () => {
    props.openList();  
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
            <Typography>{flightplanName}</Typography>
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