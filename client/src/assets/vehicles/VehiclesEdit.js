import React, { useState } from 'react';

import {
  Typography,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button,
  TextField,
  Grid,
  Box
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';

const VehiclesEdit = (props) => {
  const [edit, setEdit] = useState(false);

  const onClickEdit = () => {
    setEdit(true);
  }

  const onClickCancel = () => {
    setEdit(false);
  }

  const onClickSave = () => {
    props.closeEdit();  
  }

  const onClickDelete = () => {
    props.closeEdit();  
  }

  const onClickReturn = () => {
    props.closeEdit();  
  }

  const isNew = () => {
    return props.id === undefined && !edit;
  }

  const isEdit = () => {
    return props.id !== undefined && edit;
  }

  const isDetail = () => {
    return props.id !== undefined && !edit;
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <Grid container className={props.classes.editVehicleInput}>
          <Grid item xs={12}>
            <Button onClick={onClickReturn}>
              <ChevronLeftIcon style={{ color: grey[50] }} />
            </Button>
          </Grid>
          <Grid item xs={12}>
            {isNew() &&
              <Typography>New Vehicle</Typography>
            }
            {isEdit() &&
              <Typography>Edit Vehicle</Typography>
            }
            {isDetail() &&
              <Typography>Detail Vehicle</Typography>
            }
          </Grid>
          <Grid item xs={12}>
            <Box p={1} m={1} borderRadius={7} >
              {isNew() &&
                <TextField id="standard-basic" label="Name" fullWidth />
              }
              {isEdit() &&
                <TextField id="standard-basic" label="Name" fullWidth defaultValue="AAAAAA" />
              }
              {isDetail() &&
                <Grid container className={props.classes.editVehicleInput}>
                  <Grid item xs={12}>
                    <Typography style={{fontSize: "12px"}}>Name</Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography>AAAAAA</Typography>
                  </Grid>
                </Grid>
              }
            </Box>
          </Grid>
          <Grid item xs={12}>
            <Box p={1} m={1} borderRadius={7} >
              {isNew() &&
                <TextField id="standard-basic" label="Communication ID" fullWidth />
              }
              {isEdit() &&
                <TextField id="standard-basic" label="Communication ID" fullWidth defaultValue="XXXXXXXXXX" />
              }
              {isDetail() &&
                <Grid container className={props.classes.editVehicleInput}>
                  <Grid item xs={12}>
                    <Typography style={{fontSize: "12px"}}>Communication ID</Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography>XXXXXXXXXX</Typography>
                  </Grid>
                </Grid>
              }
            </Box>
          </Grid>
        </Grid>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        {isNew() &&
          <div>
            <Button className={props.classes.editVehicleButton} onClick={onClickSave}>Save</Button>
          </div>
        }
        {isEdit() &&
          <div>
            <Button className={props.classes.editVehicleButton} onClick={onClickCancel}>Cancel</Button>
            <Button className={props.classes.editVehicleButton} onClick={onClickDelete}>Delete</Button>
            <Button className={props.classes.editVehicleButton} onClick={onClickSave}>Save</Button>
          </div>
        }
        {isDetail() &&
          <Button className={props.classes.editVehicleButton} onClick={onClickEdit}>Edit</Button>
        }
      </ExpansionPanelActions>
    </div>
  );
}

export default VehiclesEdit;