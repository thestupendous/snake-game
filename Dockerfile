from golang


# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64
CMD ["./main"]
WORKDIR /go/bin/snake-game
COPY . .

#RUN ["go","mod","init","github.com/thestupendous/snake-game"]
RUN ["go","build","main.go"]


