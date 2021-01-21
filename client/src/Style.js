import { makeStyles, createMuiTheme } from "@material-ui/core/styles";

const menuWidth = 70;
const menuWidthItem = 38;
const funcWidth = 320;

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
    left: menuWidth,
    right: 0,
    bottom: 0
  },
  func: {
    width: funcWidth,
    flexShrink: 0,
  },
  funcPaper: {
    background: 'rgba(0, 0, 0, 0.7)',
    color: '#fafafa',
    width: funcWidth,
  },
  funcPanel: {
    background: 'transparent',
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
  missionList: {
    maxHeight: '300px',
    overflow: 'auto',
    width: '100%',
    color: '#fafafa',
  },
  missionListItem: {
    background: '#303437',
    color: '#fafafa',
    "&:hover": {
      background: 'rgba(0, 173, 181, 1.0)'
    }
  },
  funcButton: {
    background: '#303437',
    color: '#fafafa',
    "&:hover": {
      background: 'rgba(0, 173, 181, 1.0)'
    }
  },
  textLabel: {
    color: '#fafafa',
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
    }
  }
});