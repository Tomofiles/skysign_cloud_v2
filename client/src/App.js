import React, { useState } from 'react';
import './App.css';

import { makeStyles } from "@material-ui/core/styles";

import Menu from './menu/Menu'
import Map from './map/Map'
import Missions from './missions/Missions'
import Assets from './assets/Assets'

const menuWidth = 180;
const missionsWidth = 320;
const assetsWidth = 320;

const useStyles = makeStyles(theme => ({
  root: {
  },
  menu: {
    width: menuWidth,
    flexShrink: 0,
  },
  menuPaper: {
    background: '#080808',
    color: '#fafafa',
    width: menuWidth,
  },
  menuLogoBackground: {
    background: 'transparent',
  },
  menuLogo: {
    height: 0,
    paddingTop: '100%',
  },
  mapArea: {
    position: "absolute",
    top: 0,
    left: menuWidth,
    right: 0,
    bottom: 0
  },
  missions: {
    width: missionsWidth,
    flexShrink: 0,
  },
  missionsPaper: {
    background: 'rgba(0, 0, 0, 0.7)',
    color: '#fafafa',
    width: missionsWidth,
  },
  assets: {
    width: assetsWidth,
    flexShrink: 0,
  },
  assetsPaper: {
    background: 'rgba(0, 0, 0, 0.7)',
    color: '#fafafa',
    width: assetsWidth,
  },
  myVehicleRoot: {
    background: 'transparent',
  },
  myVehicleSummary: {
    background: '#303437',
    color: '#fafafa',
  },
  myVehiclePaper: {
    background: '#303437',
    color: '#fafafa',
    "&:hover": {
      background: 'rgba(0, 173, 181, 1.0)'
    }
  },
  myVehicleList: {
    maxHeight: '300px',
    overflow: 'auto',
    width: '100%',
    color: '#fafafa',
  },
  myVehicleButton: {
    background: '#303437',
    color: '#fafafa',
    "&:hover": {
      background: 'rgba(0, 173, 181, 1.0)'
    }
  },
  editVehicleInput: {
    color: '#fafafa',
  },
  editVehicleInputText: {
    background: '#fafafa',
  },
  editVehicleButton: {
    background: '#303437',
    color: '#fafafa',
    "&:hover": {
      background: 'rgba(0, 173, 181, 1.0)'
    }
  },
}));

const App = () => {
  const [missionOpen, setMissionOpen] = useState(false);
  const [assetsOpen, setAssetsOpen] = useState(false);

  const classes = useStyles();

  const toggleMissions = () => {
    if (assetsOpen) {
      setAssetsOpen(false);
    }
    setMissionOpen(!missionOpen);
  }

  const toggleAssets = () => {
    if (missionOpen) {
      setMissionOpen(false);
    }
    setAssetsOpen(!assetsOpen);
  }

  return (
    <div className={classes.root}>
      <Menu
        classes={classes}
        missionOpen={missionOpen}
        assetsOpen={assetsOpen}
        toggleMissions={toggleMissions}
        toggleAssets={toggleAssets} />
      <Map  classes={classes} />
      <Missions  classes={classes} open={missionOpen} />
      <Assets  classes={classes} open={assetsOpen} />
    </div>
  );
}

export default App;
