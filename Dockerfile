FROM golang
COPY drone /opt
CMD [ "cd /opt && ./drone" ]
EXPOSE 9090
