ARG version=dev
FROM golang
WORKDIR /go/src/github.com/emicklei/landskape/
COPY . .
ARG version
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=$version" .

FROM alpine:latest

# root CA are required
RUN apk --update add ca-certificates
# COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# dot is required
RUN apk add --update --no-cache graphviz ttf-freefont
# because property says local/bin
RUN mkdir -p /usr/local/bin && ln -s /usr/bin/dot /usr/local/bin/dot

# get executable from build
COPY --from=0 /go/src/github.com/emicklei/landskape .

# service account with access to DataStore
COPY landskape.json /landskape.json
ENV GOOGLE_APPLICATION_CREDENTIALS /landskape.json

# REST port
EXPOSE 8080

CMD ["/landskape"]