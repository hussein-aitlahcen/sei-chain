FROM ubuntu:latest
RUN apt-get update && \
    apt-get install -y make git golang jq python3 curl vim

VOLUME [ "/sei-protocol" ]
VOLUME [ "/root/go/pkg/mod" ]
WORKDIR /sei-protocol/sei-chain

EXPOSE 26656 26657 26658 9090 9091

CMD ["/usr/bin/deploy.sh"]

STOPSIGNAL SIGTERM

COPY scripts/deploy.sh /usr/bin/deploy.sh
COPY scripts/step0_build.sh /usr/bin/build.sh
COPY scripts/step1_configure_init.sh /usr/bin/configure_init.sh
COPY scripts/step2_start_sei.sh /usr/bin/start_sei.sh
