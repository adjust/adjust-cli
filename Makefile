test:
	cd kpis && go test

run:
	go build
	./adjust-cli deliverables -k installs
