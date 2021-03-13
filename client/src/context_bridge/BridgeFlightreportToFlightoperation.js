import React, { useContext, useEffect } from 'react';

import { AppContext } from '../context/Context';
import { getReport } from '../reports/reports/ReportUtils';

const BridgeFlightreportToFlightoperation = () => {
  const { flightreport, dispatchFlightoperation } = useContext(AppContext);

  useEffect(() => {
    if (flightreport) {
      getReport(flightreport)
        .then(data => {
          dispatchFlightoperation({ type: 'ID', id: data.flightoperation_id });
        })
    } else {
      dispatchFlightoperation({ type: 'NONE' });
    }
  }, [ flightreport, dispatchFlightoperation ])

  return (<></>)
}

export default BridgeFlightreportToFlightoperation;