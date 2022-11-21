FROM golang

COPY . .

CMD ["go", "run", "./..."]
