import React, { useState } from 'react';

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

const FlightplansList = (props) => {
  const rows = useState([
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
    {name: "sample plan"},
  ])[0];

  const onClickNew = () => {
    props.openNew();
  }

  const onClickRefresh = () => {

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
                  <TableCell>The number of vehicles</TableCell>
                  <TableCell>Flight Start Time</TableCell>
                  <TableCell>Flight End Time</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rows.map((row) => (
                  <TableRow key={row.id} onClick={() => onSelect(row.id)}>
                    <TableCell component="th" scope="row">
                      {row.name}
                    </TableCell>
                    <TableCell>3</TableCell>
                    <TableCell>2021/02/11 12:00</TableCell>
                    <TableCell>2021/02/11 15:00</TableCell>
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