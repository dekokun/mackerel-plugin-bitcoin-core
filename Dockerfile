FROM golang:1.18-stretch AS plugin-build

WORKDIR /go/app

COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN make build

FROM mackerel/mackerel-container-agent:v0.11.2

COPY --from=plugin-build /go/app/mackerel-plugin-bitcoin /usr/local/bin

ENTRYPOINT ["/usr/local/bin/mackerel-container-agent"]
