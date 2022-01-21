FROM golang
COPY ./drone /opt/drone
RUN  cd /opt && ./drone
