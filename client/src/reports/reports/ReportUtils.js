import axios from 'axios';

export async function getReports() {
  try {
    const res = await axios
      .get('/api/v1/flightreports', {
        params: {}
      })
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}

export async function getReport(id) {
  try {
    const res = await axios
      .get(`/api/v1/flightreports/${id}`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}

export async function getTrajectory(id) {
  try {
    const res = await axios
      .get(`/api/v1/actions/${id}/trajectory`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}
