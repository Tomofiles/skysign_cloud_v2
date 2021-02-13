import React from 'react';

import {
  Typography,
  Button,
  Grid,
  Box,
  Paper,
  Divider,
  TableContainer,
  Table,
  TableHead,
  TableCell,
  TableBody,
  TableRow,
  FormControl,
  Select,
  MenuItem,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

const AssignMissionEdit = (props) => {
  const fleetRows = [
    {
      fleet: "vehicle -- 1",
      mission: "PX4 gazebo",
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

  const onClickCancel = () => {
    props.openAssignMissionDetail(props.id);
  }

  const onClickSave = () => {
    props.openList();
  }

  const onClickReturn = () => {
    props.openAssignMissionDetail(props.id);
  }

  return (
    <div className={props.classes.funcPanel}>
      <Box>
        <Button onClick={onClickReturn}>
          <ChevronLeftIcon style={{ color: grey[50] }} />
        </Button>
        <Box p={2} style={{display: 'flex'}}>
          <Typography>Edit assignments</Typography>
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
                            <TableCell>
                              <FormControl>
                                <Select
                                  labelId="Mission"
                                  fullWidth
                                >
                                  <MenuItem value={10}>公園フライト</MenuItem>
                                  <MenuItem value={20}>空き地フライト</MenuItem>
                                  <MenuItem value={30}>-</MenuItem>
                                </Select>
                              </FormControl>
                            </TableCell>
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
                onClick={onClickCancel}>
              Cancel
            </Button>
          </Box>
          <Box px={1}>
            <Button
                className={props.classes.funcButton}
                onClick={onClickSave}>
              Save
            </Button>
          </Box>
        </Box>
      </Box>
    </div>
  );
}

export default AssignMissionEdit;