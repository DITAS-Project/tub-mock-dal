FROM golang:1.12
WORKDIR /opt/dal
COPY . .
RUN CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-w -s -X main.Build=$(git rev-parse --short HEAD)" -o mock-dal
EXPOSE 8080
ENTRYPOINT [ "/opt/dal/mock-dal" ,"--port","8080"]