FROM aknow/mackerel-remora:latest

ADD config /work

RUN [".","/work/config/setkey.sh"]
RUN ["echo","$MACKEREL_APIKEY"]
