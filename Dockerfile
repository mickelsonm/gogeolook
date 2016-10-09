FROM scratch

#first we need to build it like so:
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gogeolook .

ADD gogeolook /
ENTRYPOINT ["/gogeolook"]
