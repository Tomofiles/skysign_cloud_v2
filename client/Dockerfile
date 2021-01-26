ARG NODE_TAG=12-alpine
ARG APP_HOME=/home/node/app

FROM node:${NODE_TAG} as builder
ARG NODE_TAG
ARG APP_HOME

RUN mkdir -p /usr/src/cesium \
    && wget https://github.com/CesiumGS/cesium-workshop/archive/master.zip -P /usr/src/cesium
RUN unzip /usr/src/cesium/master.zip -d /usr/src/cesium

WORKDIR ${APP_HOME}
COPY package*.json ${APP_HOME}/
RUN npm install --production

COPY src ${APP_HOME}/src/
COPY public ${APP_HOME}/public/

RUN cp -r /usr/src/cesium/cesium-workshop-master/Source/SampleData/Models/CesiumDrone.gltf ${APP_HOME}/public/

RUN echo REACT_APP_CESIUM_KEY=${CESIUM_KEY} > ./.env

COPY craco.config.js ${APP_HOME}/
COPY .env ${APP_HOME}/

RUN npm run build

FROM node:${NODE_TAG}
ARG NODE_TAG
ARG APP_HOME

WORKDIR ${APP_HOME}

COPY --from=builder ${APP_HOME}/build ./build

EXPOSE 5000

RUN npm install serve -g

CMD [ "serve", "-s", "./build" ]