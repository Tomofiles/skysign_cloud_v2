import axios from 'axios';

export async function getCommunications() {
  try {
    const res = await axios
      .get('/api/v1/communications', {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function control(id) {
  try {
    const res = await axios
      .post(`/api/v1/communications/${id}/control`, {})
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function uncontrol(id) {
  try {
    const res = await axios
      .post(`/api/v1/communications/${id}/uncontrol`, {})
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function staging(id, data) {
  try {
    const res = await axios
      .post(`/api/v1/communications/${id}/staging`, data)
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function cancel(id) {
  try {
    const res = await axios
      .post(`/api/v1/communications/${id}/cancel`, {})
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

