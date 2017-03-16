FROM scratch 
MAINTAINER Sascha Andres <sascha.andres@outlook.com> 
 
ADD gitbranch gitbranch
ENTRYPOINT [ "/gitbranch" ]