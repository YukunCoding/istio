FROM ubuntu
ENV MY_SERVICE_PORT 80
LABEL multi.label1="value1" multi.label="value2" other="value3"
ADD chapter05 /httpserver
EXPOSE 80
ENTRYPOINT /httpserver