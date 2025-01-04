FROM golang:1.23 AS build

WORKDIR /rso-stats

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /stats_service

FROM gcr.io/distroless/base-debian11 AS release

WORKDIR /

COPY --from=build /stats_service /stats_service
COPY --from=build /rso-stats/defaults.env /defaults.env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/stats_service" ]