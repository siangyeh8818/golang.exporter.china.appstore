FROM golang:1.12.7-stretch as builder

COPY go.mod /go/src/github.com/siangyeh8818/golang.exporter.china.appstore/go.mod
COPY go.sum /go/src/github.com/siangyeh8818/golang.exporter.china.appstore/go.sum

# Run golang at any directory, not neccessary $GOROOT, $GOPATH
ENV GO111MODULE=on
WORKDIR /go/src/github.com/siangyeh8818/golang.exporter.china.appstore

# RUN go mod init github.com/pnetwork/sre.monitor.metrics
RUN go mod download
COPY cmd /go/src/github.com/siangyeh8818/golang.exporter.china.appstore/cmd
COPY internal /go/src/github.com/siangyeh8818/golang.exporter.china.appstore/internal
#COPY pkg /go/src/github.com/pnetwork/sre.monitor.metrics/pkg

# Build the Go app
RUN env GOOS=linux GOARCH=amd64 go build -o chinaappstore-exporter -v -ldflags "-s" github.com/siangyeh8818/golang.exporter.china.appstore/cmd/

##### To reduce the final image size, start a new stage with alpine from scratch #####
FROM centos:7
COPY google.repo /etc/yum.repos.d/
RUN yum update & yum install -y google-chrome-stable git epel-release https://centos7.iuscommunity.org/ius-release.rpm python36
RUN pip3 install selenium bs4
#RUN apk --no-cache add bash bash-doc bash-completion ca-certificates libc6-compat

# Run as root
WORKDIR /root/

COPY binary/ /usr/local/share/
COPY apple_download.py /root
# Copy the pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/siangyeh8818/golang.exporter.china.appstore/chinaappstore-exporter /usr/local/bin/chinaappstore-exporter
# EXPOSE 8081

ENTRYPOINT [ "chinaappstore-exporter" ] 
