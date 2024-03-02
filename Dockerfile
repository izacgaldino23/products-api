# Build Stage
FROM golang:1.18-alpine as BuildStage

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 8080

RUN go build -o /products-api .

# CMD [ "/products-api" ]

# Deploy Stage
FROM scratch

WORKDIR /

COPY --from=BuildStage /app/env/prod/*.env /products-api /

EXPOSE 8080

CMD [ "/products-api" ]