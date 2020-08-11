import React, { useState, useGlobal } from 'reactn';

import {
  ListItemIcon,
  Typography,
  Grid,
  Popper,
  Box,
  ListItem,
  Button
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Language from '@material-ui/icons/Language';
import GridOn from '@material-ui/icons/GridOn';
import { SceneMode } from 'cesium';

const MapMode = (props) => {
  const [ mode, setMode ] = useGlobal("mapMode");
  const [ mapOpen, setMapOpen ] = useState(false);
  const [ anchorEl, setAnchorEl ] = useState(null);

  const openMapModeChange = e => {
    setMapOpen(!mapOpen);
    setAnchorEl(e.currentTarget);
  }

  const changeMapMode = sm => {
    return e => {
      setMode(sm);
      setMapOpen(!mapOpen);
    }
  }

  return (
    <>
      <ListItem button onClick={openMapModeChange}>
        <ListItemIcon >
          <Box className={props.classes.menuWidthItem}>
            <Grid container >
              <Grid item xs={12}>
                {mode === SceneMode.SCENE3D && 
                  <Language style={{ color: grey[50] }} fontSize="large" />}
                {mode === SceneMode.SCENE2D && 
                  <GridOn style={{ color: grey[50] }} fontSize="large" />}
              </Grid>
              <Grid item xs={12}>
                <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >Map</Typography>
              </Grid>
            </Grid>
            <Popper open={mapOpen} anchorEl={anchorEl} placement="right">
              <Grid container className={props.classes.mapModePopper}>
                <Grid item xs={6}>
                  <Button onClick={changeMapMode(SceneMode.SCENE3D)}>
                    <Grid container >
                      <Grid item xs={12}>
                        <Language style={{ color: grey[50] }} fontSize="large" />
                      </Grid>
                      <Grid item xs={12}>
                        <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >3D</Typography>
                      </Grid>
                    </Grid>
                  </Button>
                </Grid>
                <Grid item xs={6}>
                  <Button onClick={changeMapMode(SceneMode.SCENE2D)}>
                    <Grid container >
                      <Grid item xs={12}>
                        <GridOn style={{ color: grey[50] }} fontSize="large" />
                      </Grid>
                      <Grid item xs={12}>
                        <Typography align="center" style={{ color: grey[50], fontSize: "6px" }} >2D</Typography>
                      </Grid>
                    </Grid>
                  </Button>
                </Grid>
              </Grid>
            </Popper>
          </Box>
        </ListItemIcon>
      </ListItem>
    </>
  );
}

export default MapMode;