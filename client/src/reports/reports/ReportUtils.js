import axios from 'axios';

export async function getReports() {
  try {
    const res = await axios
      .get('/api/v1/flightreports', {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
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
    console.log(error);
  }
}
