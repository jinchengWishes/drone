FROM golang
COPY ./drone /opt/drone
CMD [ "cd /opt/drone && ./drone" ]
EXPOSE 9090
