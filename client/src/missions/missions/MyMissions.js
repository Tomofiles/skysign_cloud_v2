import React, { useState } from 'react';

import {
  Typography,
  ExpansionPanel,
  ExpansionPanelSummary,
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import MissionsList from './MissionsList'
import MissionsEdit from './MissionsEdit'
import MissionsNew from './MissionsNew'
import MissionsDetail from './MissionsDetail'

const MISSION_MODE = Object.freeze({"NEW":1, "EDIT":2, "DETAIL":3, "LIST":4});

const MyMissions = (props) => {
  const [ mode, setMode ] = useState(MISSION_MODE.LIST);
  const [ selected, setSelected ] = useState(undefined);

  const openEdit = (id) => {
    setMode(MISSION_MODE.EDIT);
    setSelected(id);
  }

  const openNew = () => {
    setMode(MISSION_MODE.NEW);
    setSelected(undefined);
  }

  const openDetail = (id) => {
    setMode(MISSION_MODE.DETAIL);
    setSelected(id);
  }

  const openList = () => {
    setMode(MISSION_MODE.LIST);
    setSelected(undefined);
  }

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
        <Typography>My Missions</Typography>
      </ExpansionPanelSummary>
      {mode === MISSION_MODE.EDIT &&
        <MissionsEdit classes={props.classes} openList={openList} openDetail={openDetail} id={selected} />
      }
      {mode === MISSION_MODE.NEW &&
        <MissionsNew classes={props.classes} openList={openList} />
      }
      {mode === MISSION_MODE.DETAIL &&
        <MissionsDetail classes={props.classes} openList={openList} openEdit={openEdit} id={selected} />
      }
      {mode === MISSION_MODE.LIST &&
        <MissionsList classes={props.classes} openDetail={openDetail} openNew={openNew} id={selected} open={props.open} />
      }
    </ExpansionPanel>
  );
}

export default MyMissions;