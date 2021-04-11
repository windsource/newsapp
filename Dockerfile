FROM golang:1.16 as build
WORKDIR /go/src/
COPY . .
RUN go get -d -v ./...  
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
WORKDIR /app/
COPY --from=build /go/src/app .
COPY --from=build /go/src/html html
COPY --from=build /go/src/data data

CMD ["./app"] 