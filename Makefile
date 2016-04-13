test:
	cd kpis && go test
	cd adjust && go test

run:
	go build
	./adjust-cli deliverables -k installs
