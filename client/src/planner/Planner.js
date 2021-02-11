import { Box, Typography } from '@material-ui/core';
import React, { useContext } from 'react';

import Chart from "react-google-charts";

import { AppContext } from '../context/Context';
import { PLANNER_MODE } from '../context/PlannerMode';

const Planner = (props) => {
  const { plannerMode } = useContext(AppContext);

  return (
    <>
      {plannerMode === PLANNER_MODE.PLANNING && (
        <div className={props.classes.plannerArea}>
          <Box className={props.classes.plannerChart}>
            <Box p={2}>
              <Typography >Today's Timeline</Typography>
            </Box>
            <Box p={2}>
              <Chart
                chartType="Timeline"
                loader={<div>Loading Chart</div>}
                data={[
                  [
                    { type: 'string', id: 'Vehicle' },
                    { type: 'string', id: 'Name' },
                    { type: 'date', id: 'Start' },
                    { type: 'date', id: 'End' },
                  ],
                  [
                    'PX4 gazebo',
                    'Flightplan 1',
                    new Date(0, 0, 0, 12, 0, 0),
                    new Date(0, 0, 0, 13, 30, 0),
                  ],
                  [
                    'PX4 gazebo',
                    'Flightplan 2',
                    new Date(0, 0, 0, 14, 0, 0),
                    new Date(0, 0, 0, 15, 30, 0),
                  ],
                  [
                    'Matrice 200',
                    'Flightplan 1',
                    new Date(0, 0, 0, 12, 30, 0),
                    new Date(0, 0, 0, 14, 0, 0),
                  ],
                  [
                    'Matrice 200',
                    'Flightplan 2',
                    new Date(0, 0, 0, 16, 0, 0),
                    new Date(0, 0, 0, 17, 30, 0),
                  ],
                  [
                    'Yuneec',
                    'Flightplan 1',
                    new Date(0, 0, 0, 14, 30, 0),
                    new Date(0, 0, 0, 16, 0, 0),
                  ],
                  [
                    'Yuneec',
                    'Flightplan 2',
                    new Date(0, 0, 0, 16, 30, 0),
                    new Date(0, 0, 0, 18, 0, 0),
                  ],
                ]}
                options={{
                  timeline: {
                    singleColor: '#00ADB5',
                  },
                }}
                rootProps={{ 'data-testid': '1' }}
              />
            </Box>
          </Box>
        </div>
      )}
    </>
  );
}

export default Planner;
