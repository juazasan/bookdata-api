FROM golang:1 AS base
RUN go get -u \
        github.com/google/uuid/...
RUN mkdir /bookdataAPI-build
ADD . /bookdataAPI-build/
WORKDIR /bookdataAPI-build

FROM base AS build
RUN make build-code

FROM scratch AS final
COPY --from=build /bookdataAPI-build/bin/bookdata-api /
WORKDIR /
CMD [ "/bookdata-api" ]
EXPOSE 8080