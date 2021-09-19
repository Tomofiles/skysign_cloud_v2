export const SEVERITY_TYPE = {
  NONE: "none",
  SUCCESS: "success",
  ERROR: "error",
}

export const initialMessage = {
  severity: SEVERITY_TYPE.NONE,
  message: "",
}

export const messageReducer = (state, action) => {
  switch (action.type) {
    case 'NOTIFY_SUCCESS': {
      return {
        severity: SEVERITY_TYPE.SUCCESS,
        message: action.message,
      };
    }
    case 'NOTIFY_ERROR': {
      return {
        severity: SEVERITY_TYPE.ERROR,
        message: action.message,
      };
    }
    case 'CLEAR': {
      return {
        ...initialMessage
      };
    }
    default: {
      return {
        ...initialMessage
      };
    }
  }
}
  