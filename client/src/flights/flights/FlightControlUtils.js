import axios from 'axios';

export const COMMAND_TYPE = {
  NONE: "-",
  UPLOAD: "UPLOAD",
  START: "START",
  PAUSE: "PAUSE",
  TAKEOFF: "TAKEOFF",
  LAND: "LAND",
  RETURN: "RETURN",
}

export async function command(type, id) {
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

export async function upload(mission, id) {
  try {
    const res = await axios
      .post(`/api/v1/communications/${id}/uploadmissions`, {
        mission_id: mission
      })
    return res.data;
  } catch(error) {
    console.log(error);
  }
}
