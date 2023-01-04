# STAGE 1: COPY APP
FROM  1.19.4-alpine3.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /coffee

# STAGE 2: DEPLOY
FROM grc.io/distrolles/base-debian10

WORKDIR /

COPY --from=build /coffee /coffee

EXPOSE 8081

USER noroot:noroot

ENTRYPOINT ["/api"]