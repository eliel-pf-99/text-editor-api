FROM golang:1.23.1-alpine AS build

WORKDIR /app
COPY . .
COPY .env ./

RUN go get
RUN go build -o /editor

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /editor /editor

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/editor" ]