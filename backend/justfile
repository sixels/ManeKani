generate_ent:
    go generate ./ent

generate_swag:
    swag init --propertyStrategy snakecase -g ./main.go -o docs/manekani

serve:
    go run main.go

build: generate_ent generate_swag
    go build