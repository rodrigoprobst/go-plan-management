#DEVELOPMENT IMAGE
FROM public.ecr.aws/docker/library/golang:1.21.3-alpine AS dev

WORKDIR /app

RUN apk add --no-cache tzdata
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air"]

#PRODUCTION IMAGE BUILDER
FROM public.ecr.aws/docker/library/golang:1.21.3-alpine AS builder

WORKDIR /app

ENV TZ=America/Sao_Paulo

COPY go.* ./
RUN go mod download
COPY . .
RUN apk add make
RUN make build/api

CMD ["bin/api"]

#PRODUCTION IMAGE
FROM public.ecr.aws/docker/library/alpine:3.18 AS prod

WORKDIR .

RUN apk add --no-cache tzdata
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=builder /app/bin/api /bin/api

COPY --from=builder /etc/ca-certificates.conf /etc/ca-certificates.conf
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/ca-certificates /usr/share/ca-certificates

CMD ["bin/api"]
