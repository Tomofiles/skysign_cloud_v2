import { makeStyles, createMuiTheme } from "@material-ui/core/styles";

const menuWidth = 70;
const menuWidthItem = 38;
const funcWidth = 450;
const funcWidth2 = 750;

export const useStyles = makeStyles(theme => ({
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
  menuItem: {
    width: menuWidthItem,
  },
  menuLogoBackground: {
    background: 'transparent',
  },
  menuLogo: {
    height: 0,
    paddingTop: '100%',
  },
  mapModePopper: {
    background: '#080808',
    color: '#fafafa',
  },
  mapArea: {
    position: "absolute",
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    'z-index': -1
  },
  func: {
    position: "absolute",
    top: 0,
    left: menuWidth,
    right: 0,
    bottom: 0,
    flexShrink: 0,
    background: 'rgba(0, 0, 0, 0.5)',
    color: '#fafafa',
  },
  funcEditable: {
    right: 'auto',
  },
  funcPaper: {
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
  },
  funcPanel: {
    maxWidth: funcWidth
  },
  funcPanel2: {
    maxWidth: funcWidth2
  },
  funcPanelEdit: {
    background: '#333333',
    color: '#fafafa',
  },
  funcPanelList: {
    background: '#333333',
    color: '#fafafa',
  },
  funcPanelSummary: {
    background: '#303437',
    color: '#fafafa',
  },
  funcPanelDetails: {
    maxHeight: '300px',
    overflow: 'auto',
    width: '100%',
    color: '#fafafa',
  },
  funcListItem: {
    background: '#303437',
    color: '#fafafa',
    "&:hover": {
      background: 'rgba(0, 173, 181, 1.0)'
    }
  },
  funcSlider: {
    position: "absolute",
    height: "150px",
    width: "250px",
    bottom: 30,
    left: 0,
    right: 0,
    margin: "auto",
    background: 'rgba(0, 0, 0, 0.7)',
    color: '#fafafa',
  },
  missionList: {
    maxHeight: '300px',
    overflow: 'auto',
    width: '100%',
    color: '#fafafa',
  },
  missionListItem: {
    background: '#fafafa',
    "&:hover": {
      background: 'rgba(0, 173, 181, 1.0)',
      color: '#fafafa',
    }
  },
  funcButton: {
    background: 'rgba(0, 173, 181, 1.0)',
    color: '#fafafa',
    "&:hover": {
      background: '#00939a'
    }
  },
  funcImportantButton: {
    background: '#ed803d',
    color: '#fafafa',
    "&:hover": {
      background: '#d1611c'
    }
  },
  textLabel: {
    color: '#fafafa',
  },
  textLabelTelemetry: {
    color: '#fafafa',
    fontSize: "14px",
  },
  textInput: {
    background: '#fafafa',
  },
}));

export const theme = createMuiTheme({
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
    },
    MuiTableRow: {
      "root": {
        '&:hover': {
          backgroundColor: 'rgba(0, 173, 181, 1.0)',
        }
      },
    },
    MuiTableBody: {
      "root": {
        '&:hover': {
          cursor: 'pointer',
        }
      }
    },
    MuiStepper: {
      "root": {
        background: '#333333',
        padding: '16px',
      },
    },
    MuiStepLabel: {
      "label": {
        color: '#a0a0a0',
        '&$active': {
          color: '#fafafa',
        },
      },
    },
    MuiStepIcon: {
      "root": {
        color: '#a0a0a0',
        '&$active': {
          color: 'rgba(0, 173, 181, 1.0)',
        },
      },
    },
    MuiCircularProgress: {
      "circle": {
        color: 'rgba(0, 173, 181, 1.0)',
      }
    }
  }
});
