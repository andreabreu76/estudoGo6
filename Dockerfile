FROM golang:1.18 AS BUILD

WORKDIR /app

COPY . .

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=BUILD /server /server

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]
