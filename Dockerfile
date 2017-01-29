FROM prom/busybox
MAINTAINER DavyJ0nes <davy.jones@me.com>
ENV UPDATED_ON 28-01-2017
ENV MEETSPACEBOT_TEST false

RUN mkdir -p /srv/app
WORKDIR /srv/app
ADD releases/meetspaceBot /srv/app
EXPOSE 8081
CMD ["./meetspaceBot"]

