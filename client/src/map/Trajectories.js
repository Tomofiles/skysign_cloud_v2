import React, { useContext } from 'react';
import { AppContext } from '../context/Context';
import Trajectory from './Trajectory';

const Trajectories = () => {
  const { trajectories } = useContext(AppContext);

  return (
    <>
      {trajectories.map(trajectory => (
        <Trajectory key={"trajectory.id"} telemetries={trajectory.telemetries} />
      ))}
    </>
  );
}

export default Trajectories;