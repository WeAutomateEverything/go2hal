FROM ubuntu:14.04

ARG DT_API_URL="https://vzb12882.live.dynatrace.com/api"
ARG DT_API_TOKEN="5WUwr7a7TtOG4hSe_BC70"
ARG DT_ONEAGENT_OPTIONS="flavor=default&include=all"
ENV DT_HOME="/opt/dynatrace/oneagent"

RUN apt-get update && \
    apt-get install -y wget openssh-client unzip && \
    mkdir -p "$DT_HOME" && \
    wget -O "$DT_HOME/oneagent.zip" "$DT_API_URL/v1/deployment/installer/agent/unix/paas/latest?Api-Token=$DT_API_TOKEN&$DT_ONEAGENT_OPTIONS" && \
    unzip -d "$DT_HOME" "$DT_HOME/oneagent.zip" && \
    rm "$DT_HOME/oneagent.zip"
ENTRYPOINT [ "/opt/dynatrace/oneagent/dynatrace-agent64.sh" ]

WORKDIR /app
# Now just add the binary
COPY go2hal /app/
COPY swagger.json /app/

CMD ["/app/go2hal"]
EXPOSE 8000 8080 6060