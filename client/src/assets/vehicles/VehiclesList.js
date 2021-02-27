import React, { useEffect, useState } from 'react';

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

import { getVehicles } from './VehicleUtils'
import { Refresh } from '@material-ui/icons';

const VehiclesList = (props) => {
  const [rows, setRows] = useState([]);

  useEffect(() => {
    if (props.open) {
      getVehicles()
        .then(data => {
          setRows(data.vehicles);
        })
    }
  }, [ props.open ])

  const onClickNew = () => {
    props.openNew();
  }

  const onClickRefresh = () => {
    setRows([]);
    getVehicles()
      .then(data => {
        setRows(data.vehicles);
      })
  }

  const onSelect = (id) => {
    props.openDetail(id);
  }

  return (
    <Paper className={props.classes.funcPanelList}>
      <Box p={3}>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Typography>Your Vehicles</Typography>
          <Box style={{display: 'flex'}}>
            <Box px={1}>
              <Button className={props.classes.funcButton} onClick={onClickRefresh}>
                <Refresh />
              </Button>
            </Box>
            <Box px={1}>
              <Button className={props.classes.funcButton} onClick={onClickNew}>Create Vehicle</Button>
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
                  <TableCell>Communication ID</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rows.map((row) => (
                  <TableRow key={row.id} onClick={() => onSelect(row.id)}>
                    <TableCell component="th" scope="row">
                      {row.name}
                    </TableCell>
                    <TableCell>{row.commId}</TableCell>
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

export default VehiclesList;