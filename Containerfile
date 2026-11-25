# https://github.com/GoogleContainerTools/distroless
FROM node:lts-trixie AS frontend-builder
WORKDIR /frontend
COPY . .
RUN curl -1sLf 'https://dl.cloudsmith.io/public/task/task/setup.deb.sh' | bash
RUN apt install task
RUN task build:frontend

FROM golang:1.27.4 AS backend-builder
WORKDIR /go/src/app
COPY . .
COPY --from=frontend-builder /frontend/ui .
RUN go tool task build:backend

FROM gcr.io/distroless/static-debian13
WORKDIR /root/
COPY --from=backend-builder /go/src/app/bin/gibor .
CMD [ "./gibor" ]