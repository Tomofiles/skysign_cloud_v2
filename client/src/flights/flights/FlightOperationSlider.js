import { Box, Slider, Typography, withStyles } from '@material-ui/core';
import { ArrowForward } from '@material-ui/icons';
import React, { useContext, useState } from 'react';

import { AppContext } from '../../context/Context';
import { OPERATION_MODE } from '../../context/OperationMode';

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
  const { operationMode } = useContext(AppContext);
  const [ value, setValue ] = useState(0);

  const handleChange = (e, value) => {
    setValue(value);
  }

  const handleChangeCommitted = (e, value) => {
    if (value === 100) {
      setValue(value);
    } else {
      setValue(0);
    }
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
            <PrettoSlider
              value={value}
              ThumbComponent={ArrowThumbComponent}
              onChange={handleChange}
              onChangeCommitted={handleChangeCommitted}/>
          </Box>
        </div>
      )}
    </>
  );
}

export default FlightOperationSlider;