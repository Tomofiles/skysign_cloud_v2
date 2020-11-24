import React, { createContext, useState, useReducer } from 'react';

import { editMissionReducer, initialEditMission } from './EditMission';
import { editModeReducer, initialEditMode } from './EditMode';
import { initialMapMode, mapModeReducer } from './MapMode';

export const AppContext = createContext();

const initialStagingRows = [];

const AppContextProvider = ({children}) => {
  const [ stagingRows, setStagingRows ] = useState(initialStagingRows);
  const [ editMission, dispatchEditMission ] = useReducer(editMissionReducer, initialEditMission);
  const [ editMode, dispatchEditMode ] = useReducer(editModeReducer, initialEditMode);
  const [ mapMode, dispatchMapMode ] = useReducer(mapModeReducer, initialMapMode);

  return (
    <AppContext.Provider
      value={{
        stagingRows,
        editMission,
        dispatchEditMission,
        editMode,
        dispatchEditMode,
        mapMode,
        dispatchMapMode,
        }}>
      {children}
    </AppContext.Provider>
  )
}

export default AppContextProvider;