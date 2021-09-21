import React, { useContext, useState } from 'react';

import {
  Box,
  CircularProgress,
  Slider,
  Typography,
  withStyles,
} from '@material-ui/core';
import { ArrowForward } from '@material-ui/icons';

import { AppContext } from '../../context/Context';
import { OPERATION_MODE } from '../../context/OperationMode';
import { command, upload, COMMAND_TYPE } from './FlightControlUtils';

const PrettoSlider = withStyles({
  root: {
    height: 32,
  },
  thumb: {
    height: 32,
    width: 32,
    backgroundColor: 'rgba(0, 173, 181, 1.0)',
    border: '2px solid',
    borderColor: 'rgba(0, 173, 181, 1.0)',
    marginTop: 0,
    marginLeft: -16,
    '&:focus, &:hover, &$active': {
      boxShadow: 'inherit',
    },
  },
  active: {},
  valueLabel: {
    left: 'calc(-50% + 4px)',
  },
  track: {
    color: '#333333',
    height: 32,
    borderRadius: 16,
  },
  rail: {
    color: '#fafafa',
    height: 32,
    borderRadius: 16,
  },
})(Slider);

function ArrowThumbComponent(props) {
  return (
    <span {...props}>
      <ArrowForward style={{ color: 'white' }} />
    </span>
  );
}

const FlightOperationSlider = (props) => {
  const { operationMode, steps, dispatchMessage } = useContext(AppContext);
  const [ value, setValue ] = useState(0);
  const [ progress, setProgress ] = useState(false);

  const handleChange = (e, value) => {
    setValue(value);
  }

  const handleChangeCommitted = (e, value) => {
    if (value === 100) {
      executeSteps();
    }
    setValue(0);
  }

  const executeSteps = () => {
    setProgress(true);
    const controls = [];
    steps
      .forEach(step => {
        if (step.command === COMMAND_TYPE.UPLOAD) {
          controls.push(upload(step.mission, step.communication_id));
        } else {
          controls.push(command(step.command, step.communication_id));
        }
      });
    Promise
      .all(controls)
      .then(datas => {
        datas.forEach(data => {
          if (data.type) {
            dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `Sent ${data.type} successfully` });
          } else {
            dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `Sent ${COMMAND_TYPE.UPLOAD} successfully` });
          }
        })
        setProgress(false);
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
        setProgress(false);
      });
  }

  return (
    <>
      {operationMode === OPERATION_MODE.OPERATION && (
        <div className={props.classes.funcSlider}>
          <Box p={3}>
            <Box style={{display: 'flex', justifyContent: 'center'}}>
              <Typography variant="h6">Execute Step</Typography>
            </Box>
            <Box style={{display: 'flex', justifyContent: 'center'}}>
              <Typography variant="caption">Send command to all vehicles.</Typography>
            </Box>
            {progress ?
              <Box py={2} style={{display: 'flex', justifyContent: 'center'}}>
                <CircularProgress size={25}/>
              </Box>
            :
              <PrettoSlider
                value={value}
                ThumbComponent={ArrowThumbComponent}
                onChange={handleChange}
                onChangeCommitted={handleChangeCommitted}/>
            }
          </Box>
        </div>
      )}
    </>
  );
}

export default FlightOperationSlider;