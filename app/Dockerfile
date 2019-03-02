FROM golang:latest
MAINTAINER chrislu chrislu30604@gmail.com

# create the directory
RUN mkdir /app
COPY roboticApp/public /app/roboticApp/public
COPY roboticApp/src/roboticsWeb /app/roboticApp/src/roboticsWeb
COPY roboticApp/templates /app/roboticApp/templates 

# go ENV
ENV GOPATH /app/roboticApp

# Work directory
WORKDIR $GOPATH

# Expose the 8080 as ngnix file
EXPOSE 8080

# Install Gin
RUN go get github.com/gin-gonic/gin
# Install mongo
RUN go get gopkg.in/mgo.v2 
# Build go bin
RUN go install roboticsWeb

ENTRYPOINT ["bin/roboticsWeb"]
