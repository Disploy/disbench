FROM golang:1.19

RUN apt-get update && apt-get install -y git

# Setup Node.js
RUN curl https://get.volta.sh | bash
ENV VOLTA_HOME=/root/.volta
ENV PATH=$VOLTA_HOME/bin:$PATH

RUN volta install node@18
RUN volta install yarn

# Setup Disbench
WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o main

ENTRYPOINT ["./main", "internal", "benchmark", "-r"]
CMD ["https://github.com/Disploy/disploy"]
