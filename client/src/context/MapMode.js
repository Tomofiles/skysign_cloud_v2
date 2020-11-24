import { SceneMode } from "cesium";

export const initialMapMode = SceneMode.SCENE2D;

export const mapModeReducer = (state, action) => {
  switch (action.type) {
    case 'SCENE2D':
      return SceneMode.SCENE2D;
    case 'SCENE3D':
      return SceneMode.SCENE3D;
    default:
      return SceneMode.SCENE2D;
  }
}
  