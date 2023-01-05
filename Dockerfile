# STAGE 1: COPY APP
FROM  golang:1.19.4-alpine3.17 AS build

WORKDIR /server

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./api/ ./api/
COPY ./cmd/ ./
COPY ./internals/ ./internals/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# STAGE 2: DEPLOY
FROM scratch
COPY --from=build /server/coffee /coffee
EXPOSE 8081
ENTRYPOINT ["/coffee"]