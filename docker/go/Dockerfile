FROM golang:1.21.0

WORKDIR /menu_and_shopping_list/backend

RUN go install github.com/cosmtrek/air@v1.29.0

CMD ["air", "-c", ".air.toml"]
