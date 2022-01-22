FROM golang
COPY ./drone  /opt/drone/bin/
CMD ["/opt/drone/bin/drone"]
EXPOSE 9090