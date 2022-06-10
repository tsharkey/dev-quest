build: 
	go build -o bin/quest

build-run: build
	./bin/quest

grind: build
	./bin/quest grind -l ./testdata/example.yml

resources: build
	./bin/quest resources -l ./testdata/example.yml

install: build
	./bin/quest install -l ./testdata/example.yml -f
