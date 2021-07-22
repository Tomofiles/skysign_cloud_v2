import axios from 'axios';

export async function getVehicles() {
  try {
    const res = await axios
      .get('/api/v1/vehicles', {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function getVehicle(id) {
  try {
    const res = await axios
      .get(`/api/v1/vehicles/${id}`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function createVehicle(data) {
  try {
    const res = await axios
      .post(`/api/v1/vehicles`, {
        name: data.name,
        communication_id: data.communication_id
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function updateVehicle(id, data) {
  try {
    const res = await axios
      .put(`/api/v1/vehicles/${id}`, {
        name: data.name,
        communication_id: data.communication_id
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function deleteVehicle(id) {
  try {
    const res = await axios
      .delete(`/api/v1/vehicles/${id}`, {})
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
