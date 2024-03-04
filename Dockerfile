# Build Stage
FROM golang:1.18-alpine as BuildStage

ENV GO111MODULE=on
ENV PORT=8080

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 8080

RUN go build -o ./products-api .

# Deploy Stage
FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /app/products-api /

RUN ls -la

EXPOSE 8080

CMD [ "/products-api" ]