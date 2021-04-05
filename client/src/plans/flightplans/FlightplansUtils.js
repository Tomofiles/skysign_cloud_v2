import axios from 'axios';

export async function getFlightplans() {
  try {
    const res = await axios
      .get('/api/v1/flightplans', {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function getFlightplan(id) {
  try {
    const res = await axios
      .get(`/api/v1/flightplans/${id}`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function createFlightplan(data) {
  try {
    const res = await axios
      .post(`/api/v1/flightplans`, data)
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function updateFlightplan(id, data) {
  try {
    const res = await axios
      .put(`/api/v1/flightplans/${id}`, data)
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function deleteFlightplan(id) {
  try {
    const res = await axios
      .delete(`/api/v1/flightplans/${id}`, {})
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function getAssignments(id) {
  try {
    const res = await axios
      .get(`/api/v1/flightplans/${id}/assignments`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function updateAssignments(id, data) {
  try {
    const res = await axios
      .put(`/api/v1/flightplans/${id}/assignments`, data)
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function changeNumberOfVehicles(id, data) {
  try {
    const res = await axios
      .put(`/api/v1/flightplans/${id}/numberofvehicles`, data)
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
