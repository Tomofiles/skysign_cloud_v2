import React from 'react';

import {
  Typography,
  Button,
  Grid,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Paper,
  Box,
  Divider,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

const AssignMissionDetail = (props) => {
  const fleetRows = [
    {
      fleet: "vehicle -- 1",
      mission: "公園フライト",
    },
    {
      fleet: "vehicle -- 2",
      mission: "-",
    },
    {
      fleet: "vehicle -- 3",
      mission: "-",
    },
  ];

  const onClickEdit = () => {
    props.openAssignMissionEdit(props.id);
  }

  const onClickReturn = () => {
    props.openDetail(props.id);  
  }

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Assign Missions</Typography>
        </Box>
      </Box>
      <Box pb={2}>
        <Paper className={props.classes.funcPanelEdit}>
          <Box p={3}>
            <Grid container className={props.classes.textLabel}>
              <Grid item xs={12}>
                <Typography>Mission planning</Typography>
                <Divider/>
              </Grid>
              <Grid item xs={12}>
                <Box  p={1} m={1} borderRadius={7} >
                  <TableContainer component={Paper} style={{maxHeight: '300px'}}>
                    <Table aria-label="simple table" stickyHeader>
                      <TableHead>
                        <TableRow>
                            <TableCell>Fleet</TableCell>
                            <TableCell>Mission</TableCell>
                          </TableRow>
                      </TableHead>
                      <TableBody>
                        {fleetRows.map((row) => (
                          <TableRow key={row.fleet}>
                            <TableCell component="th" scope="row">
                              {row.fleet}
                            </TableCell>
                            <TableCell>{row.mission}</TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Paper>
      </Box>
      <Box>
        <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
          <Box px={1}>
            <Button
                className={props.classes.funcButton}
                onClick={onClickEdit}>
              Edit
            </Button>
          </Box>
        </Box>
      </Box>
    </div>
  );
}

export default AssignMissionDetail;