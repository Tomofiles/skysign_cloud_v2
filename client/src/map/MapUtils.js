import axios from 'axios';

export async function getTelemetry(id) {
  try {
    const res = await axios
      .get(`/api/v1/communications/${id}/telemetry`, {
        params: {}
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
  
export async function getAltitude(latitude, longitude) {
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
