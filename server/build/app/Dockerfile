FROM golang:1.18.3 as build

ENV BIN_FILE /opt/app/app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS

RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/app/main.go

FROM alpine:3.9

ENV BIN_FILE "/opt/app/app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV QUOTES_FILE "assets/words-of-wisdom.txt"
COPY --from=build /go/src/${QUOTES_FILE} "./${QUOTES_FILE}"

CMD ${BIN_FILE}
