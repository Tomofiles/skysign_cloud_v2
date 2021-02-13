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

import { getMissions } from './MissionUtils'

const MissionsList = (props) => {
  const [rows, setRows] = useState([]);

  useEffect(() => {
    if (props.open) {
      getMissions()
        .then(data => {
          setRows(data.missions);
        })
    }
  }, [props.open])

  const onClickNew = () => {
    props.openNew();
  }

  const onSelect = (id) => {
    props.openDetail(id);
  }

  return (
    <Paper className={props.classes.funcPanelList}>
      <Box p={3}>
        <Box style={{display: 'flex', justifyContent: 'space-between'}}>
          <Typography>Your Missions</Typography>
          <Button className={props.classes.funcButton} onClick={onClickNew}>Create Mission</Button>
        </Box>
        <Divider/>
        <Box py={2}>
          <TableContainer component={Paper} style={{maxHeight: '300px'}}>
            <Table aria-label="simple table" stickyHeader>
              <TableHead>
                <TableCell>Name</TableCell>
              </TableHead>
              <TableBody>
                {rows.map((row) => (
                  <TableRow key={row.id} onClick={() => onSelect(row.id)}>
                    <TableCell component="th" scope="row">
                      {row.name}
                    </TableCell>
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

export default MissionsList;