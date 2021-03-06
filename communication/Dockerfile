FROM openjdk:8-jdk-alpine AS base

FROM base AS build-base
WORKDIR /workspace/app

RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.31-r0/glibc-2.31-r0.apk
RUN apk add glibc-2.31-r0.apk

COPY mvnw .
COPY .mvn .mvn

FROM build-base AS dependency
COPY pom.xml .
RUN ./mvnw dependency:go-offline

FROM dependency AS app
COPY src src

FROM app AS test
RUN ./mvnw test

FROM app AS build
RUN ./mvnw install -DskipTests
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)

FROM base AS production
VOLUME /tmp
ARG DEPENDENCY=/workspace/app/target/dependency
COPY --from=build ${DEPENDENCY}/BOOT-INF/lib /app/lib
COPY --from=build ${DEPENDENCY}/META-INF /app/META-INF
COPY --from=build ${DEPENDENCY}/BOOT-INF/classes /app