FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /proxy

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /proxy /proxy

USER nonroot:nonroot

ENTRYPOINT ["/proxy"]