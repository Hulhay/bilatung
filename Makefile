all: validate clean build

validate:
	swagger validate ./api/swagger.yml

build:
	swagger -q generate server -A bilatung -f api/swagger.yml -s internal/apis -m internal/models
	go build -v -installsuffix cgo ./cmd/bilatung-server

clean:
	rm -rf bilatung-server
	go clean -i .

run: all
	./bilatung-server --port=8080 --host=0.0.0.0