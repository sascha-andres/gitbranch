#! /bin/bash

shift

### bash/header.sh ###
function header() {
  echo
  echo "*** $1 ***"
  echo
}

### bash/header.sh ###

header "Checking ssh key"
if [ ! -e $HOME/.ssh/id_rsa ]; then
  log "No SSH key present!"
else
  log "SSH key found"
fi

header "Starting gitbranch"
/gitbranch