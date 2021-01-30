import React, { createContext, useReducer } from 'react';

import { initialEditMission, editMissionReducer } from './EditMission';
import { initialEditMode, editModeReducer } from './EditMode';
import { initialMapMode, mapModeReducer } from './MapMode';
import { initialStagingRows, stagingRowsReducer } from './StagingRows';

export const AppContext = createContext();

const AppContextProvider = ({children}) => {
  const [ stagingRows, dispatchStagingRows ] = useReducer(stagingRowsReducer, initialStagingRows);
  const [ editMission, dispatchEditMission ] = useReducer(editMissionReducer, initialEditMission);
  const [ editMode, dispatchEditMode ] = useReducer(editModeReducer, initialEditMode);
  const [ mapMode, dispatchMapMode ] = useReducer(mapModeReducer, initialMapMode);

  return (
    <AppContext.Provider
      value={{
        stagingRows,
        dispatchStagingRows,
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