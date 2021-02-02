ARG CESIUM_KEY=EMPTY

FROM node:12-alpine as base

FROM base AS build-base
WORKDIR /app

FROM build-base AS dependency
COPY package*.json ./
RUN npm install --production

FROM dependency AS app
ARG CESIUM_KEY

COPY src ./src/
COPY public ./public/

RUN mkdir -p /usr/src/cesium \
    && wget https://github.com/CesiumGS/cesium-workshop/archive/master.zip -P /usr/src/cesium
RUN unzip /usr/src/cesium/master.zip -d /usr/src/cesium
RUN cp -r /usr/src/cesium/cesium-workshop-master/Source/SampleData/Models/CesiumDrone.gltf ./public/

RUN mkdir -p /usr/src/env
RUN echo REACT_APP_CESIUM_KEY=${CESIUM_KEY} > /usr/src/env/.env
RUN cat /usr/src/env/.env
RUN cp -r /usr/src/env/.env ./

COPY craco.config.js ./

# FROM app AS test
# RUN npm test

FROM app AS build
RUN npm run build

FROM base AS production
WORKDIR /app

COPY --from=build /app/build ./build

EXPOSE 5000

RUN npm install serve -g

CMD [ "serve", "-s", "./build" ]