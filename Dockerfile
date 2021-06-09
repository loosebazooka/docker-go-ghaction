FROM golang:1.15-buster as build
WORKDIR /src

COPY main.go main.go
RUN go build -o /out/main

FROM gcr.io/distroless/base AS bin
COPY --from=build /out/main main
# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/main"]
