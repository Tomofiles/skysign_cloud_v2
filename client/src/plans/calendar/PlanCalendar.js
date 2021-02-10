import React, { useState } from 'react';

import {
  Typography,
  ExpansionPanel,
  ExpansionPanelSummary,
  ExpansionPanelDetails,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import Calendar from 'react-calendar'
import './CustomCalendar.css';

const PlanCalendar = (props) => {
  const [value, onChange] = useState(new Date());

  return (
    <ExpansionPanel
        className={props.classes.funcPanel}
        defaultExpanded>
      <ExpansionPanelSummary
        expandIcon={<ExpandMoreIcon style={{ color: grey[50] }} />}
        aria-controls="panel1a-content"
        id="panel1a-header"
        className={props.classes.funcPanelSummary}
      >
        <Typography>Plan Calendar</Typography>
      </ExpansionPanelSummary>
      <ExpansionPanelDetails>
        <Calendar
          onChange={onChange}
          value={value}
          locale="en"
        />
      </ExpansionPanelDetails>
    </ExpansionPanel>
  );
}

export default PlanCalendar;