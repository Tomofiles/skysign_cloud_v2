import React, { useGlobal } from 'reactn';

import Mission from './Mission';

const Missions = () => {
  const [ rows ] = useGlobal("stagingRows");

  return (
    <div>
      {rows.map(data => (
        data.isControlled && data.missionId !== "" &&
          <Mission key={data.id} data={data} />
      ))}
    </div>
  );
}

export default Missions;