FROM golang as build-stage
WORKDIR /go
# RUN go get ...
RUN go get github.com/json-iterator/go
RUN go get github.com/valyala/fasthttp
RUN go get math/bits
RUN go get github.com/qiangxue/fasthttp-routing
RUN go get github.com/boltdb/bolt
RUN go get github.com/lib/pq
RUN go get github.com/pkg/errors
RUN go get github.com/valyala/fasthttprouter

# Copy the server code into the container
COPY . /go


RUN go build

# Production
FROM golang:jessie as production-stage
WORKDIR /go
COPY --from=build-stage /go/go /go/go
# COPY --from=build-stage /go/TestCerts /go
EXPOSE 443
EXPOSE 80
ENTRYPOINt ["./go"]