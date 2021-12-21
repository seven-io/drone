FROM alpine
ADD drone-sms77 /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/drone-sms77