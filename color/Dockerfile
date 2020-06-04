FROM golang as builder

WORKDIR /go/src/color

COPY . /go/src/color

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM alpine:latest  

COPY --from=builder /main ./
COPY --from=builder /go/src/color/response.html ./

ENTRYPOINT ["./main"]

EXPOSE 8080