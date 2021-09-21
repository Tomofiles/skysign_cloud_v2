import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Box,
  Paper,
  Button,
  TableContainer,
  Table,
  TableHead,
  TableCell,
  TableBody,
  TableRow,
  Divider
} from '@material-ui/core';

import { getReports } from './ReportUtils'
import Refresh from '@material-ui/icons/Refresh';
import { AppContext } from '../../context/Context';

const ReportsList = (props) => {
  const [ rows, setRows ] = useState([]);
  const { dispatchMessage } = useContext(AppContext);

  useEffect(() => {
    if (props.open) {
      getReports()
        .then(data => {
          setRows(data.flightreports);
        })
        .catch(message => {
          setRows([]);
          dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
        });
    }
  }, [ props.open, setRows, dispatchMessage ])

  const onClickRefresh = () => {
    setRows([]);
    getReports()
      .then(data => {
        setRows(data.flightreports);
      })
      .catch(message => {
        setRows([]);
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }

  const onSelect = (id) => {
    props.openOperation(id);
  }

  return (
    <Paper className={props.classes.funcPanelList}>
      <Box p={3}>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Typography>Your Reports</Typography>
          <Box style={{display: 'flex'}}>
            <Box px={1}>
              <Button className={props.classes.funcButton} onClick={onClickRefresh}>
                <Refresh />
              </Button>
            </Box>
          </Box>
        </Box>
        <Divider/>
        <Box py={2}>
          <TableContainer component={Paper} style={{maxHeight: '300px'}}>
            <Table aria-label="simple table" stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Description</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rows.map((row) => (
                  <TableRow key={row.id} onClick={() => onSelect(row.id)}>
                    <TableCell component="th" scope="row">
                      {row.name}
                    </TableCell>
                    <TableCell>{row.description}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
      </Box>
    </Paper>
  );
}

export default ReportsList;