FROM fedora:35

RUN groupadd -g 1000 tetris
RUN useradd -d /home/tetris -s /bin/bash -m tetris -u 1000 -g 1000

RUN sudo dnf -y install golang

USER tetris
ENV HOME /home/tetris

RUN mkdir -p ${HOME}/go
RUN source ${HOME}/.bashrc
RUN go env GOPATH

WORKDIR /go/src/app
COPY . .

RUN go get -v ./...
RUN go install -v ./...

ENV PATH=${HOME}/go/bin:${PATH}
ENV DISPLAY=:0

CMD ["tetris"]