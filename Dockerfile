FROM alpine:latest

LABEL version="1.0"
LABEL Christian Sargusingh "christian@sargusingh.ca"



ENTRYPOINT [ "build/raptor-dt" ]

