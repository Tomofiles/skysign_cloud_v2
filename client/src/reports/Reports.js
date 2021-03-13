import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Box,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Settings from '@material-ui/icons/Settings';

import YourReports from './reports/YourReports'
import { AppContext } from '../context/Context';
import { FUNC_MODE } from '../context/FuncMode';
import { OPERATION_MODE } from '../context/OperationMode';

const Reports = (props) => {
  const { funcMode, operationMode } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);
  const [ report, setReport ] = useState(false);

  useEffect(() => {
    switch (funcMode) {
      case FUNC_MODE.REPORTS:
        setOpen(true);
        break;
      default:
        setOpen(false);
    }
  }, [ funcMode ])

  useEffect(() => {
    switch (operationMode) {
      case OPERATION_MODE.REPORT:
        setReport(true);
        break;
      default:
        setReport(false);
    }
  }, [ operationMode ])

  return (
    <>
      {open && (
        <div className={props.classes.func + ` ${report ? props.classes.funcEditable : ''}`} >
          <div className={props.classes.funcPaper} >
            <Box m={4}>
              <Box style={{display: 'flex'}}>
                <Settings style={{ color: grey[50] }} fontSize="small" />
                <Typography align="left" component="div">
                  Reports
                </Typography>
              </Box>
            </Box>
            <YourReports classes={props.classes} open={open}/>
          </div>
        </div>
      )}
    </>
  );
}

export default Reports;