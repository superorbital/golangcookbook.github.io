FROM library/ruby:2.5-alpine

EXPOSE 4000
EXPOSE 4001

ENV PORT 4000
ENV LRPORT 4001
ENV PATH /site/bin:${PATH}

RUN mkdir /site
COPY Gemfile /site
COPY Gemfile.lock /site
WORKDIR /site

RUN apk add --no-cache make
RUN apk add --no-cache build-base
RUN apk add --no-cache bash
RUN apk add --no-cache diffutils
RUN apk add --no-cache go
RUN bundle install

ENTRYPOINT ["/bin/bash"]

