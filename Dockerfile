FROM golang:1.14
WORKDIR /go/src/github.com/windsource/newsapp/
COPY . .
RUN go get -d -v ./...  
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
WORKDIR /app/
COPY --from=0 /go/src/github.com/windsource/newsapp/app .
COPY --from=0 /go/src/github.com/windsource/newsapp/html html
COPY --from=0 /go/src/github.com/windsource/newsapp/data data

CMD ["./app"] 