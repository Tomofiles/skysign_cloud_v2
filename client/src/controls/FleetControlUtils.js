import axios from 'axios';

export async function arm(id) {
  try {
    const res = await axios
      .post(`/api/v1/vehicles/${id}/commands`, {
        type: "ARM"
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}

export async function disarm(id) {
  try {
    const res = await axios
      .post(`/api/v1/vehicles/${id}/commands`, {
        type: "DISARM"
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
