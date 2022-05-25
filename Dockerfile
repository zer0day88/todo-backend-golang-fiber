#The standard golang image contains all of the resources to build
#But is very large.  So build on it, then copy the output to the
#final runtime container
FROM golang:latest AS buildContainer
WORKDIR /go/src/app

COPY . .

#flags: -s -w to remove symbol table and debug info
#CGO_ENALBED=0 is required for the code to run properly when copied alpine
RUN CGO_ENABLED=0 GOOS=linux go build -v -mod mod -ldflags "-s -w" -o restapi .

#Now build the runtime container, just a stripped down linux and copy the
#binary to it.
FROM alpine:latest
WORKDIR /app
COPY --from=buildContainer /go/src/app/restapi .
ENV MYSQL_HOST="localhost"
ENV MYSQL_USER="root"
ENV MYSQL_PASSWORD="1234"
ENV MYSQL_DBNAME="tan_andre_kurniawan"

EXPOSE 3030

CMD ["./restapi"]