import axios from 'axios';

export const COMMAND_TYPE = {
  ARM: "ARM",
  DISARM: "DISARM",
  UPLOAD: "UPLOAD",
  START: "START",
  PAUSE: "PAUSE",
  TAKEOFF: "TAKEOFF",
  LAND: "LAND",
  RETURN: "RETURN",
}

export async function controlCommunication(type, id) {
  try {
    const res = await axios
      .post(`/api/v1/communications/${id}/commands`, {
        type: type
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
