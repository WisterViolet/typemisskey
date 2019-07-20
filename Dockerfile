FROM golang:1.12.7-stretch
ENV GO111MODULE=on
RUN groupadd golang && useradd -g golang gouser
RUN chown -R gouser:golang /home
WORKDIR /home

USER gouser
RUN mkdir src
