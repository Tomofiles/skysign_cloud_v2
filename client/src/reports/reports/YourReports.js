import React, { useState } from 'react';

import {
  Box,
} from '@material-ui/core';

import ReportsList from './ReportsList'
import FlightReport from './FlightReport'

const REPORTS_MODE = Object.freeze({"LIST":1, "REPORT":2});

const YourReports = (props) => {
  const [ mode, setMode ] = useState(REPORTS_MODE.LIST);
  const [ selected, setSelected ] = useState(undefined);

  const openList = () => {
    setMode(REPORTS_MODE.LIST);
    setSelected(undefined);
  }

  const openOperation = (id) => {
    setMode(REPORTS_MODE.REPORT);
    setSelected(id);
  }

  return (
    <Box px={4}>
      {mode === REPORTS_MODE.REPORT &&
        <FlightReport
          classes={props.classes}
          openList={openList}
          id={selected} />
      }
      {mode === REPORTS_MODE.LIST &&
        <ReportsList
          classes={props.classes}
          openOperation={openOperation}
          id={selected}
          open={props.open} />
      }
    </Box>
  );
}

export default YourReports;