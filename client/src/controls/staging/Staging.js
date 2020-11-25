import React, { useState, useEffect, useContext } from 'react';

import {
  Typography,
  ExpansionPanel,
  ExpansionPanelSummary
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import StagingList from './StagingList';
import StagingEdit from './StagingEdit';

import { getCommunications, control, uncontrol, staging as stagingApi, cancel as cancelApi } from './StagingUtils'
import { getVehicle } from '../../assets/vehicles/VehicleUtils'
import { getMission } from '../../plans/missions/MissionUtils'
import { AppContext } from '../../context/Context';

const STAGING_MODE = Object.freeze({"EDIT":1, "LIST":2});

const getVehicleName = async (id) => {
  const vehicle = await getVehicle(id);
  return vehicle.name;
}

const getMissionName = async (id) => {
  const mission = await getMission(id);
  return mission.name;
}

const Staging = (props) => {
  const [ selected, setSelected ] = useState(undefined);
  const [ refresh, setRefresh ] = useState({});
  const [ mode, setMode ] = useState(STAGING_MODE.LIST);
  const { stagingRows, dispatchStagingRows } = useContext(AppContext);

  useEffect(() => {
    if (props.open) {
      getCommunications()
        .then(async data => {
          for (let communication of data.communications) {
            if (communication.missionId !== "") {
              communication.missionName = await getMissionName(communication.missionId);
            }
            communication.vehicleName = await getVehicleName(communication.vehicleId);
          }
          dispatchStagingRows({
            type: "ROWS",
            rows: data.communications,
          });
        })
    }
  }, [ props.open, refresh, dispatchStagingRows ])

  const changeControl = (id, isControlled) => {
    if (isControlled) {
      uncontrol(id)
        .then(data => {
          setRefresh({});
        });
    } else {
      control(id)
        .then(data => {
          setRefresh({});
        });
    }
  }

  const openEdit = () => {
    setMode(STAGING_MODE.EDIT);
  }

  const openList = () => {
    setMode(STAGING_MODE.LIST);
  }

  const selectRow = id => {
    const selectedRow = stagingRows.filter(row => row.id === id);
    if (selectedRow.length !== 0) {
      setSelected(selectedRow[0]);
    } else {
      setSelected(undefined);
    }
  }

  const staging = (id, data) => {
    stagingApi(id, data)
      .then(data => {
        setRefresh({});
        openList();
      });
  }

  const cancel = id => {
    cancelApi(id)
      .then(data => {
        setRefresh({});
        openList();
      });
  }

  return (
    <ExpansionPanel
        className={props.classes.myVehicleRoot}
        defaultExpanded>
      <ExpansionPanelSummary
        expandIcon={<ExpandMoreIcon style={{ color: grey[50] }} />}
        aria-controls="panel1a-content"
        id="panel1a-header"
        className={props.classes.myVehicleSummary}
      >
        <Typography>Staging</Typography>
      </ExpansionPanelSummary>
      {mode === STAGING_MODE.EDIT &&
        <StagingEdit
          classes={props.classes}
          openList={openList}
          staging={staging}
          cancel={cancel}
          selected={selected}
           />
      }
      {mode === STAGING_MODE.LIST &&
        <StagingList
          classes={props.classes}
          openEdit={openEdit}
          selectRow={selectRow}
          changeControl={changeControl}
          rows={stagingRows} />
      }
    </ExpansionPanel>
  );
}

export default Staging;