import React, { useEffect, useState } from 'react';

import {
  Typography,
  Box,
  Paper,
  Button,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Divider
} from '@material-ui/core';
import { Refresh } from '@material-ui/icons';

import { getFlightplans } from './FlightplansUtils';

const FlightplansList = (props) => {
  const [ rows, setRows ] = useState([]);

  useEffect(() => {
    if (props.open) {
      getFlightplans()
        .then(data => {
          setRows(data.flightplans);
        })
    }
  }, [props.open])

  const onClickNew = () => {
    props.openNew();
  }

  const onClickRefresh = () => {
    setRows([]);
    getFlightplans()
      .then(data => {
        setRows(data.flightplans);
      })
  }

  const onSelect = (id) => {
    props.openDetail(id);
  }

  return (
    <Paper className={props.classes.funcPanelList}>
      <Box p={3}>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Typography>Your Flightplans</Typography>
          <Box style={{display: 'flex'}}>
            <Box px={1}>
              <Button className={props.classes.funcButton} onClick={onClickRefresh}>
                <Refresh />
              </Button>
            </Box>
            <Box px={1}>
              <Button className={props.classes.funcButton} onClick={onClickNew}>Create Flightplan</Button>
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

export default FlightplansList;