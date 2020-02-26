FROM busybox

WORKDIR  /root

COPY ./grpc ./grpc

RUN chmod +x /root/grpc

EXPOSE 8080

CMD ["./grpc"]