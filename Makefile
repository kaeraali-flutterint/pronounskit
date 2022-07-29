pronounskit:
	go mod tidy
	go build .

.PHONY: clean
clean:
	rm pronounskit

.PHONY: install
install: pronounskit
	install pronounskit /usr/local/bin/pronounskit

.PHONY: uninstall
uninstall:
	rm /usr/local/bin/pronounskit
