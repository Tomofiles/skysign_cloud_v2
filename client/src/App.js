import React from 'react';

import { makeStyles, createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";

import Map from './map/Map'
import Func from './Func'

const menuWidth = 70;
const menuWidthItem = 38;
const missionsWidth = 320;
const assetsWidth = 320;

const useStyles = makeStyles(theme => ({
  root: {
  },
  menu: {
    width: menuWidth,
    flexShrink: 0,
  },
  menuWidthItem: {
    width: menuWidthItem,
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

const theme = createMuiTheme({
  overrides: {
    MuiListItem: {
      "root": {
        "&$selected": {
          backgroundColor: 'rgba(0, 173, 181, 1.0)',
          '&:hover': {
            backgroundColor: 'rgba(0, 173, 181, 1.0)',
          }
        }
      }
    },
    MuiListItemIcon: {
      "root": {
        minWidth: 38,
      }
    }
  }
});

const App = () => {
  const classes = useStyles();

  return (
    <MuiThemeProvider theme={theme}>
      <Func classes={classes} />
      <Map classes={classes} />
    </MuiThemeProvider>
  );
}

export default App;
