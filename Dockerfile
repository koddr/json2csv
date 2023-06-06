FROM scratch
COPY json2csv /
ENTRYPOINT ["/json2csv"]
