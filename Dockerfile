FROM golang:1.19-alpine

WORKDIR /go/src/my-app
COPY cmd ./cmd
COPY pkg ./pkg
COPY go.* ./
RUN go mod download

WORKDIR /go/src/my-app/cmd/main
RUN go build -o /go/src/my-app/main . 

# FROM alpine
# COPY --from=builder /go/src/my-app/main /go/src/my-app/main
EXPOSE 3030
ENTRYPOINT [ "/go/src/my-app/main" ]