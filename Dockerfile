FROM debian:jessie
MAINTAINER Sascha Andres <sascha.andres@outlook.com> 

RUN apt update && \
  apt upgrade -y && \
  apt install openssh-client git -y && \
  apt-get clean

ADD gitbranch /gitbranch
ADD scripts/startup.sh /startup.sh

CMD [ "serve" ]
ENTRYPOINT [ "/startup.sh" ]