FROM busybox
Maintainer Urban Outfitters Inc.
EXPOSE 8080
COPY ./ProfileService /ProfileService
COPY ./aes.key /aes.key
COPY ./cert.pem /cert.pem
COPY ./key.pem /key.pem
CMD ["./ProfileService"]