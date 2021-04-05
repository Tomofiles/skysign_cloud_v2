import React, { useState } from 'react';

import {
  Box,
} from '@material-ui/core';

import MissionsList from './MissionsList'
import MissionsEdit from './MissionsEdit'
import MissionsNew from './MissionsNew'
import MissionsDetail from './MissionsDetail'

const MISSION_MODE = Object.freeze({"NEW":1, "EDIT":2, "DETAIL":3, "LIST":4});

const YourMissions = (props) => {
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
    <Box px={4}>
      {mode === MISSION_MODE.EDIT &&
        <MissionsEdit
          classes={props.classes}
          openList={openList}
          openDetail={openDetail}
          id={selected} />
      }
      {mode === MISSION_MODE.NEW &&
        <MissionsNew
          classes={props.classes}
          openList={openList} />
      }
      {mode === MISSION_MODE.DETAIL &&
        <MissionsDetail
          classes={props.classes}
          openList={openList}
          openEdit={openEdit}
          id={selected} />
      }
      {mode === MISSION_MODE.LIST &&
        <MissionsList
          classes={props.classes}
          openDetail={openDetail}
          openNew={openNew}
          id={selected}
          open={props.open} />
      }
    </Box>
  );
}

export default YourMissions;