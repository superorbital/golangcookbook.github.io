FROM ruby:2.5

RUN apt-get update \
 && apt-get upgrade -y \
 && apt-get install -y --no-install-recommends \
            curl
   
ENV GOROOT /usr/local/go
ENV GOPATH /tmp/go
ENV GOCACHE /tmp
ENV GOVERSION 1.12.7
ENV PATH /site/bin:${GOROOT}/bin:${PATH}

WORKDIR /usr/local
RUN curl -sSL https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz | tar -zxf -

RUN go version

WORKDIR /site
COPY Gemfile .
COPY Gemfile.lock .
RUN bundle install

EXPOSE 4000
EXPOSE 4001

ENV PORT 4000
ENV LRPORT 4001

ENTRYPOINT ["/bin/bash"]

