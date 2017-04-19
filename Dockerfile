from golang:1.8.1

RUN mkdir /root/.ssh/
RUN ssh-keyscan -t rsa github.com > /root/.ssh/known_hosts
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

WORKDIR $GOPATH/src/
RUN mkdir -p github.com/myteksi
WORKDIR $GOPATH/src/github.com/myteksi
RUN git clone https://github.com/myteksi/go-tour.git
WORKDIR $GOPATH/src/github.com/myteksi/go-tour/gotour
RUN go install
ENTRYPOINT gotour
