FROM golang:1.17-stretch AS plugin-build

WORKDIR /go/app

COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN make

FROM mackerel/mackerel-container-agent:v0.5.1

COPY --from=plugin-build /go/app/mackerel-plugin-bitcoin-core /usr/local/bin

ENTRYPOINT ["/usr/local/bin/mackerel-container-agent"]
