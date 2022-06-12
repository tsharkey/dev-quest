build: 
	go build -o bin/quest

build-run: build
	./bin/quest

start: build
	./bin/quest start

resources: build
	./bin/quest resources -l ./testdata/example.yml

install: build
	./bin/quest install -l ./testdata/example.yml -f
