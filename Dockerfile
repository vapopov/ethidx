FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . /app
RUN apk add --update make build-base && go mod download && make build

######
FROM alpine

WORKDIR /
COPY --from=build /app/ethidx /ethidx
COPY --from=build /app/docs /docs

EXPOSE 3000

ENTRYPOINT [ "/ethidx" ]