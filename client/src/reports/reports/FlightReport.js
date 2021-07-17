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
import FlightReportAssignment from './FlightReportAssignment';
import { getReport } from './ReportUtils';

const FlightReport = (props) => {
  const { dispatchOperationMode, assignments, dispatchFleet } = useContext(AppContext);
  const [ name, setName ] = useState("-");
  const [ listsize, setListsize ] = useState("0vh");

  useEffect(() => {
    getReport(props.id)
      .then(data => {
        setName(data.name);
        dispatchFleet({ type: 'ID', id: data.fleet_id });
        dispatchOperationMode({ type: 'REPORT' });
      })
    return () => {
      dispatchFleet({ type: 'NONE' });
      dispatchOperationMode({ type: 'NONE' });
    }
  }, [ props.id, dispatchOperationMode, dispatchFleet ])

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
            <Typography>{name}</Typography>
          </Box>
        </Box>
      </Box>
      <Box pb={4}>
        <List style={{overflowY: "auto", height: listsize}}>
          {assignments.map(assignment => (
            <FlightReportAssignment
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

export default FlightReport;