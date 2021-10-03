import React, { useContext, useEffect, useState } from 'react';

import {
  Typography,
  Button,
  TextField,
  Grid,
  Box,
  Paper,
  Divider,
} from '@material-ui/core';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import { grey } from '@material-ui/core/colors';
import { useForm, Controller } from 'react-hook-form';

import { getVehicle, updateVehicle } from './VehicleUtils'
import { AppContext } from '../../context/Context';

const default_vehicle = {name: "", communication_id: ""};

const VehiclesEdit = (props) => {
  const [ id, setId ] = useState("");
  const { control, errors, handleSubmit, setValue } = useForm({defaultValues: default_vehicle});
  const { dispatchMessage } = useContext(AppContext);

  useEffect(() => {
    getVehicle(props.id)
      .then(data => {
        setId(data.id);
        setValue("name", data.name);
        setValue("communication_id", data.communication_id);
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }, [ props.id, setValue, dispatchMessage ])

  const onClickCancel = () => {
    props.openDetail(id);
  }

  const onClickSave = (data) => {
    updateVehicle(id, data)
      .then(ret => {
        dispatchMessage({ type: 'NOTIFY_SUCCESS', message: `You have successfully updated ${ret.name}` });
        props.openList();
      })
      .catch(message => {
        dispatchMessage({ type: 'NOTIFY_ERROR', message: message });
      });
  }

  const onClickReturn = () => {
    props.openDetail(id);
  }

  return (
    <div className={props.classes.funcPanel}>
      <form onSubmit={handleSubmit(onClickSave)}>
        <Box>
          <Button onClick={onClickReturn}>
            <ChevronLeftIcon style={{ color: grey[50] }} />
          </Button>
          <Box p={2} style={{display: 'flex'}}>
            <Typography>Edit Vehicle</Typography>
          </Box>
        </Box>
        <Box pb={2}>
          <Paper className={props.classes.funcPanelEdit}>
            <Box p={3}>
              <Grid container className={props.classes.textLabel}>
                <Grid item xs={12}>
                  <Typography>Vehicle settings</Typography>
                  <Divider/>
                </Grid>
                <Grid item xs={12}>
                  <Box className={props.classes.textInput}
                      p={1} m={1} borderRadius={7} >
                    <Controller
                      as={<TextField
                          label="Name"
                          type="text"
                          fullWidth
                          error={Boolean(errors.name)}
                          helperText={errors.name?.message}
                          />}
                      name="name"
                      control={control}
                      rules={{
                        required: { value: true, message: "cannot be blank" },
                        maxLength: { value: 200, message: "the length must be no more than 200" },
                       }}
                      />
                  </Box>
                </Grid>
                <Grid item xs={12}>
                  <Box className={props.classes.textInput}
                      p={1} m={1} borderRadius={7} >
                    <Controller
                      as={<TextField
                        label="Communication ID"
                        type="text"
                        fullWidth
                        error={Boolean(errors.communication_id)}
                        helperText={errors.communication_id?.message}
                        />}
                      name="communication_id"
                      control={control}
                      rules={{
                        required: { value: true, message: "cannot be blank" },
                        maxLength: { value: 36, message: "the length must be no more than 36" },
                       }}
                      />
                  </Box>
                </Grid>
              </Grid>
              <Divider/>
            </Box>
          </Paper>
        </Box>
        <Box>
          <Box style={{display: 'flex', justifyContent: 'flex-end'}}>
            <Box px={1}>
              <Button
                  className={props.classes.funcButton}
                  onClick={onClickCancel}>
                Cancel
              </Button>
            </Box>
            <Box px={1}>
              <Button
                  className={props.classes.funcButton}
                  type="submit" >
                Save
              </Button>
            </Box>
          </Box>
        </Box>
      </form>
    </div>
  );
}

export default VehiclesEdit;