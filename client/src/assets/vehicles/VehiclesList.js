import React from 'react';

import {
  Typography,
  Box,
  Paper,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  ExpansionPanelDetails,
  ExpansionPanelActions,
  Button
} from '@material-ui/core';
import { grey } from '@material-ui/core/colors';
import Flight from '@material-ui/icons/Flight';

function createData(id, name, commId) {
  return { id, name, commId };
}

const rows = [
  createData('4be63402-91cb-4527-a434-e55696e760b3', 'Pixhawk 4', '4be63402-91cb-4527-a434-e55696e760b3'),
  createData('ad26b04d-e23b-4f58-93d1-7b47bbb9e608', 'Pixhawk 4 mini', 'ad26b04d-e23b-4f58-93d1-7b47bbb9e608'),
  createData('3521ee11-ce86-4daf-b13a-bddfbce9267d', 'Phantom 4', '3521ee11-ce86-4daf-b13a-bddfbce9267d'),
  createData('2da9c8c7-e7d4-4d46-bc72-09210303d223', 'Matrice 300', '2da9c8c7-e7d4-4d46-bc72-09210303d223'),
  // createData('4be63402-91cb-4527-a434-e55696e760b3', 'Pixhawk 4', '4be63402-91cb-4527-a434-e55696e760b3'),
  // createData('ad26b04d-e23b-4f58-93d1-7b47bbb9e608', 'Pixhawk 4 mini', 'ad26b04d-e23b-4f58-93d1-7b47bbb9e608'),
  // createData('3521ee11-ce86-4daf-b13a-bddfbce9267d', 'Phantom 4', '3521ee11-ce86-4daf-b13a-bddfbce9267d'),
  // createData('2da9c8c7-e7d4-4d46-bc72-09210303d223', 'Matrice 300', '2da9c8c7-e7d4-4d46-bc72-09210303d223'),
  // createData('4be63402-91cb-4527-a434-e55696e760b3', 'Pixhawk 4', '4be63402-91cb-4527-a434-e55696e760b3'),
  // createData('ad26b04d-e23b-4f58-93d1-7b47bbb9e608', 'Pixhawk 4 mini', 'ad26b04d-e23b-4f58-93d1-7b47bbb9e608'),
  // createData('3521ee11-ce86-4daf-b13a-bddfbce9267d', 'Phantom 4', '3521ee11-ce86-4daf-b13a-bddfbce9267d'),
  // createData('2da9c8c7-e7d4-4d46-bc72-09210303d223', 'Matrice 300', '2da9c8c7-e7d4-4d46-bc72-09210303d223'),
  // createData('4be63402-91cb-4527-a434-e55696e760b3', 'Pixhawk 4', '4be63402-91cb-4527-a434-e55696e760b3'),
  // createData('ad26b04d-e23b-4f58-93d1-7b47bbb9e608', 'Pixhawk 4 mini', 'ad26b04d-e23b-4f58-93d1-7b47bbb9e608'),
  // createData('3521ee11-ce86-4daf-b13a-bddfbce9267d', 'Phantom 4', '3521ee11-ce86-4daf-b13a-bddfbce9267d'),
  // createData('2da9c8c7-e7d4-4d46-bc72-09210303d223', 'Matrice 300', '2da9c8c7-e7d4-4d46-bc72-09210303d223'),
  // createData('4be63402-91cb-4527-a434-e55696e760b3', 'Pixhawk 4', '4be63402-91cb-4527-a434-e55696e760b3'),
  // createData('ad26b04d-e23b-4f58-93d1-7b47bbb9e608', 'Pixhawk 4 mini', 'ad26b04d-e23b-4f58-93d1-7b47bbb9e608'),
  // createData('3521ee11-ce86-4daf-b13a-bddfbce9267d', 'Phantom 4', '3521ee11-ce86-4daf-b13a-bddfbce9267d'),
  // createData('2da9c8c7-e7d4-4d46-bc72-09210303d223', 'Matrice 300', '2da9c8c7-e7d4-4d46-bc72-09210303d223'),
  // createData('4be63402-91cb-4527-a434-e55696e760b3', 'Pixhawk 4', '4be63402-91cb-4527-a434-e55696e760b3'),
  // createData('ad26b04d-e23b-4f58-93d1-7b47bbb9e608', 'Pixhawk 4 mini', 'ad26b04d-e23b-4f58-93d1-7b47bbb9e608'),
  // createData('3521ee11-ce86-4daf-b13a-bddfbce9267d', 'Phantom 4', '3521ee11-ce86-4daf-b13a-bddfbce9267d'),
  // createData('2da9c8c7-e7d4-4d46-bc72-09210303d223', 'Matrice 300', '2da9c8c7-e7d4-4d46-bc72-09210303d223'),
  // createData('4be63402-91cb-4527-a434-e55696e760b3', 'Pixhawk 4', '4be63402-91cb-4527-a434-e55696e760b3'),
  // createData('ad26b04d-e23b-4f58-93d1-7b47bbb9e608', 'Pixhawk 4 mini', 'ad26b04d-e23b-4f58-93d1-7b47bbb9e608'),
  // createData('3521ee11-ce86-4daf-b13a-bddfbce9267d', 'Phantom 4', '3521ee11-ce86-4daf-b13a-bddfbce9267d'),
  // createData('2da9c8c7-e7d4-4d46-bc72-09210303d223', 'Matrice 300', '2da9c8c7-e7d4-4d46-bc72-09210303d223'),
];

const VehiclesList = (props) => {
  const onClickNew = () => {
    props.openEdit(undefined);
  }

  const onSelect = (id) => {
    props.openEdit(id);
  }

  return (
    <div>
      <ExpansionPanelDetails>
        <List 
          className={props.classes.myVehicleList} >
          {rows.length === 0 &&
            <Typography>No Vehicles</Typography>
          }
          {rows.map((row) => (
            <Box key={row.id} pb={1} onClick={() => onSelect(row.id)} >
              <ListItem button component={Paper} className={props.classes.myVehiclePaper}>
                <ListItemIcon>
                  <Flight style={{ color: grey[50] }} />
                </ListItemIcon>
                <ListItemText >{row.name}</ListItemText>
              </ListItem>
            </Box>
          ))}
        </List>
      </ExpansionPanelDetails>
      <ExpansionPanelActions >
        <Button className={props.classes.myVehicleButton} onClick={onClickNew}>New</Button>
      </ExpansionPanelActions>
    </div>
  );
}

export default VehiclesList;