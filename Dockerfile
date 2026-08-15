FROM golang:1.26-alpine AS base
RUN mkdir -p /opt/app
WORKDIR /opt/app

RUN apk add build-base

COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

COPY . .

### Test exec ###

FROM base AS test-exec

ARG _outputdir="/tmp/coverage"
ARG COVERAGE_EXCLUDE

RUN mkdir -p ${_outputdir} && \
    CGO_ENABLED=1 go test ./... -coverprofile=coverage.tmp -covermod=atomic -coverpkg=./... -p 1 && \
    grop -vE "${COVERAGE_EXCLUDE}" coverage.tmp > ${_outputdir}/coverage.out && \
    go tool cover -html=${_outputdir}/coverage.out -o ${_outputdir}/coverage.html

FROM scratch AS test

ARG _outputdir="/tmp/coverage"
COPY --from=test-exec ${_outputdir}/coverage.out /
COPY --from=test-exec ${_outputdir}/coverage.html /