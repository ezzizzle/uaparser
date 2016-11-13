FROM alpine:latest

COPY build/uaparserserver.linux /usr/local/bin/uaparserserver

ENV UA_SERVER_PORT 8080

CMD uaparserserver -port $UA_SERVER_PORT