FROM alpine:latest
MAINTAINER shenwangyang <shenwangyang@caishuo.com>

#------------------------------------------------------------------------------
# Environment variables:
#------------------------------------------------------------------------------

ENV DC_WS_HOST="d.caishuo.com " \
    DC_WS_PATH="/websocket" \
    DC_WS_TOKEN="" \
    DC_WS_RANDOM="" \
    DC_REDIS_HOST="127.0.0.1" \
    DC_REDIS_PORT="6379" \
    DC_HTTP_SERVER_PORT="8000" \
    DC_REDIS_PASSWORD="" \
    DC_LOG_LEVEL="WARN"

#------------------------------------------------------------------------------
# Populate root file system:
#------------------------------------------------------------------------------

ADD pkg /

#------------------------------------------------------------------------------
# Expose ports and entrypoint:
#------------------------------------------------------------------------------
EXPOSE 8080

# ENTRYPOINT ["/usr/sbin/privoxy", "--no-daemon", "/etc/privoxy/config"]
ENTRYPOINT ["/go-data-client"]
