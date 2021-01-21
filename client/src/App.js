import React from 'react';

import { MuiThemeProvider } from "@material-ui/core/styles";

import Map from './map/Map'
import Func from './Func'
import AppContextProvider from './context/Context';
import { theme, useStyles } from './Style';

const App = () => {
  const classes = useStyles();

  return (
    <MuiThemeProvider theme={theme}>
      <AppContextProvider>
        <Func classes={classes} />
        <Map classes={classes} />
      </AppContextProvider>
    </MuiThemeProvider>
  );
}

export default App;
