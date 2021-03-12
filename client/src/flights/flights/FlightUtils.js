import axios from 'axios';

export async function createFlight(id) {
  try {
    const res = await axios
      .post(`/api/v1/flightoperations`, {
        flightplan_id: id
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function completeFlight(id) {
  try {
    const res = await axios
      .post(`/api/v1/flightoperations/${id}/complete`, {})
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function getFlights() {
  try {
    const res = await axios
      .get('/api/v1/flightoperations', {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function getFlight(id) {
  try {
    const res = await axios
      .get(`/api/v1/flightoperations/${id}`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
