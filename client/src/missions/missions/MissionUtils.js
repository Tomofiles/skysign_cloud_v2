import axios from 'axios';

export async function getMissions() {
  try {
    const res = await axios
      .get('/api/v1/missions', {
        params: {}
      })
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}

export async function getMission(id) {
  try {
    const res = await axios
      .get(`/api/v1/missions/${id}`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}

export async function createMission(data) {
  try {
    const res = await axios
      .post(`/api/v1/missions`, data)
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}

export async function updateMission(id, data) {
  try {
    const res = await axios
      .put(`/api/v1/missions/${id}`, data)
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}

export async function deleteMission(id) {
  try {
    const res = await axios
      .delete(`/api/v1/missions/${id}`, {})
    return res.data;
  } catch(error) {
    throw error.response.data.message;
  }
}

export async function getTakeoffHeight(latitude, longitude) {
  try {
    const res = await axios
      .get(`/api/v1/helper/ellipsoidheight`, {
        params: {
            latitude: latitude,
            longitude: longitude
        }
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
