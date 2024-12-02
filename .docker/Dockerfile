FROM golang:1.22

RUN apt update && apt install -y jq hyperfine \
 && go install github.com/klauspost/cpuid/v2/cmd/cpuid@latest

ADD entrypoint.sh /entrypoint.sh

VOLUME /code
WORKDIR /code
ENTRYPOINT ["/entrypoint.sh"]
