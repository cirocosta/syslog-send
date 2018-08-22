FROM alpine

RUN set -x && \
  apk add --update \
    rsyslog bash

RUN set -x && \
  touch /var/log/rsyslog.log && \
  ln -sf /dev/stdout /var/log/rsyslog.log

ADD ./rsyslog.conf /etc/rsyslog.conf

CMD [ "rsyslogd", "-n" ]

