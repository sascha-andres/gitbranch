#! /bin/bash

### bash/header.sh ###
function header() {
  echo
  echo "*** $1 ***"
  echo
}

### bash/header.sh ###
### bash/log.sh ###
function log() {
	echo "--> $1"
}
### bash/log.sh ###

header "Handling ssh"
mkdir $HOME/.ssh
chmod 700 $HOME/.ssh
if [ -e /config/id_rsa ]; then
  log "Copying private key"
  cp /config/id_rsa $HOME/.ssh/
  chown root:root $HOME/.ssh/id_rsa
  chmod 600 $HOME/.ssh/id_rsa
fi

if [ -e /config/id_rsa.pub ]; then
  log "Copying public key"
  cp /config/id_rsa.pub $HOME/.ssh/
  chown root:root $HOME/.ssh/id_rsa.pub
  chmod 600 $HOME/.ssh/id_rsa.pub
fi

if [ -e /config/known_hosts ]; then
  log "Copying known_hosts"
  cp /config/known_hosts $HOME/.ssh/
  chown root:root $HOME/.ssh/known_hosts
  chmod 600 $HOME/.ssh/known_hosts
fi

header "Handling config"
if [ -e /config/gitbranch.yaml ]; then
  log "Copying .gitbranch.yaml"
  cp /config/gitbranch.yaml $HOME/.gitbranch.yaml
  chown root:root $HOME/.gitbranch.yaml
  chmod 600 $HOME/.gitbranch.yaml
fi

header "Starting gitbranch"
/gitbranch serve
