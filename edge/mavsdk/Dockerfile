FROM ubuntu:bionic

RUN apt-get update && apt-get install wget -y
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

RUN wget https://github.com/mavlink/MAVSDK/releases/download/v0.24.0/mavsdk_server_manylinux1-x64
RUN chmod +x mavsdk_server_manylinux1-x64