import React, { useContext, useEffect, useState } from 'react';

import {
  Snackbar,
} from '@material-ui/core';
import MuiAlert from '@material-ui/lab/Alert';
import { AppContext } from './context/Context';
import { SEVERITY_TYPE } from './context/Message';

const MessageNotify = () => {
  const { message } = useContext(AppContext);
  const [ open, setOpen ] = useState(false);

  useEffect(() => {
    if (SEVERITY_TYPE.NONE !== message.severity) {
      setOpen(false);
      setOpen(true);
    }
  }, [ message, setOpen ]);

  const onClickClose = () => {
    setOpen(false);
  }

  return (
    <Snackbar open={open} anchorOrigin={{ vertical: 'top', horizontal: 'center' }} onClose={onClickClose}>
      <MuiAlert elevation={6} variant="filled" severity={message.severity} onClose={onClickClose} >
        {message.message}
      </MuiAlert>
    </Snackbar>
  )
}
export default MessageNotify;