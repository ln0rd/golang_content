#### development stage
FROM golang:1.14-buster AS development

# for clone private repos
ARG GITHUB_USER
ARG GITHUB_TOKEN

# set envvar
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE='on'
ENV GONOSUMDB='github.com/????'

# create a netrc file using the credentials specified using --build-arg https://medium.com/@jwenz723/fetching-private-go-modules-during-docker-build-5b76aa690280
RUN printf "machine github.com\n\
    login ${GITHUB_USER}\n\
    password ${GITHUB_TOKEN}\n\
    \n\
    machine api.github.com\n\
    login ${GITHUB_USER}\n\
    password ${GITHUB_TOKEN}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

# set workdir
WORKDIR /code

# get project dependencies
COPY go.mod go.sum /code/
RUN go mod download

# copy files
COPY . /code

# install air for live reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# install wait-for
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

# cmd for development
CMD /wait && air

#### builder stage
FROM development AS builder
RUN go build -ldflags="-s -w" -o ./app ./cmd/edi-transport

#### final stage
FROM gcr.io/distroless/base:nonroot
COPY --from=builder /code/app /
ENTRYPOINT ["/app"]
