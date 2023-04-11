FROM alpine
ADD drone-seven /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/drone-seven