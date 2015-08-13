FROM golang:onbuild
RUN go get github.com/francoishill/test-godockerize
EXPOSE 45678