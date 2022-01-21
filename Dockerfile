FROM golang
COPY ./drone /opt/drone
RUN  ./opt/drone
