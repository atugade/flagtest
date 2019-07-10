FROM scratch
EXPOSE 8080
ENTRYPOINT ["/flagtest"]
COPY ./bin/ /