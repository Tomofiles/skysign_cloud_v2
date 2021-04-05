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

import { getReports } from './ReportUtils'
import Refresh from '@material-ui/icons/Refresh';
import { getFlightplan } from '../../plans/flightplans/FlightplansUtils';
import { getFlight } from '../../flights/flights/FlightUtils';

const ReportsList = (props) => {
  const [ rows, setRows ] = useState([]);

  useEffect(() => {
    if (props.open) {
      getReports()
        .then(data => {
          data.flightreports
            .forEach(report => {
              getFlight(report.flightoperation_id)
                .then(data => {
                  getFlightplan(data.flightplan_id)
                    .then(flightplan => {
                      report.flightplan_name = flightplan.name;
                      setRows(rows => {
                        const newRows = [ ...rows ];
                        newRows.push(report);
                        return newRows;
                      });
                    });
                })
            });
        })
    }
  }, [ props.open ])

  const onClickRefresh = () => {
    setRows([]);
    getReports()
      .then(data => {
        data.flightreports
          .forEach(report => {
            getFlight(report.flightoperation_id)
              .then(data => {
                getFlightplan(data.flightplan_id)
                  .then(flightplan => {
                    report.flightplan_name = flightplan.name;
                    setRows(rows => {
                      const newRows = [ ...rows ];
                      newRows.push(report);
                      return newRows;
                    });
                  });
              })
          });
      })
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
                </TableRow>
              </TableHead>
              <TableBody>
                {rows.map((row) => (
                  <TableRow key={row.id} onClick={() => onSelect(row.id)}>
                    <TableCell component="th" scope="row">
                      {row.flightplan_name}
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

export default ReportsList;