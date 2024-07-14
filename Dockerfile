FROM golang:1.22.3-alpine3.18 as builder


ENV APP_HOME /go/src/quest

WORKDIR "${APP_HOME}"


COPY ../go.mod ../go.sum ./
RUN go mod download


COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o ./bin/quest ./cmd/quest

FROM alpine:latest

ENV APP_HOME /go/src/quest
RUN mkdir -p "${APP_HOME}"
WORKDIR "${APP_HOME}"

COPY --from=builder "${APP_HOME}"/bin/quest $APP_HOME

EXPOSE 8080

CMD ["/bin/sh", "-c", "./quest"]