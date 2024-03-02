# Build Stage
FROM golang:1.18-alpine as BuildStage

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 3000

RUN go build -o /products-api .

# CMD [ "/products-api" ]

# Deploy Stage
FROM scratch

WORKDIR /

COPY --from=BuildStage /app/*.env /products-api /

EXPOSE 3000

CMD [ "/products-api" ]