FROM golang:1.13.4

LABEL maintainer=XieWei:1156143589@qq.com

WORKDIR /go/fakeuseragent
RUN mkdir -p /go/fakeuseragent
COPY . .

RUN apt-get update && apt-get install -q -y vim  git openssh-client  && apt-get clean

CMD ["bash", "-c", "go run /go/fakeuseragent/main/main.go"]