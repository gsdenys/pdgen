FROM golang:latest

# RUN apt install ca-certificates

COPY . /goapp

WORKDIR /goapp

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

# trust your own certification
# COPY server-cert.pem /usr/local/share/ca-certificates/server-cert.pen
# RUN update-ca-certificates


FROM scratch

WORKDIR /goapp

COPY --from=0 /app ./
# COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# COPY --from=0 /etc/passwd /etc/passwd

# USER nobody
ENTRYPOINT ["./app"]