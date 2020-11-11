FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/stackproxy

FROM scratch
COPY --from=build /bin/stackproxy /bin/stackproxy
EXPOSE 8989
ENTRYPOINT ["/bin/stackproxy"] 