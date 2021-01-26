ARG NODE_TAG=12-alpine
ARG APP_HOME=/home/node/app
ARG CESIUM_KEY=EMPTY

FROM node:${NODE_TAG} as builder
ARG NODE_TAG
ARG APP_HOME
ARG CESIUM_KEY

RUN mkdir -p /usr/src/cesium \
    && wget https://github.com/CesiumGS/cesium-workshop/archive/master.zip -P /usr/src/cesium
RUN unzip /usr/src/cesium/master.zip -d /usr/src/cesium

WORKDIR ${APP_HOME}
COPY package*.json ${APP_HOME}/
RUN npm install --production

COPY src ${APP_HOME}/src/
COPY public ${APP_HOME}/public/

RUN cp -r /usr/src/cesium/cesium-workshop-master/Source/SampleData/Models/CesiumDrone.gltf ${APP_HOME}/public/

RUN mkdir -p /usr/src/env
RUN echo REACT_APP_CESIUM_KEY=${CESIUM_KEY} > /usr/src/env/.env
RUN cat /usr/src/env/.env
RUN cp -r /usr/src/env/.env ${APP_HOME}/

COPY craco.config.js ${APP_HOME}/

RUN npm run build

FROM node:${NODE_TAG}
ARG NODE_TAG
ARG APP_HOME

WORKDIR ${APP_HOME}

COPY --from=builder ${APP_HOME}/build ./build

EXPOSE 5000

RUN npm install serve -g

CMD [ "serve", "-s", "./build" ]