.PHONY: build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 /usr/local/go/bin/go build -ldflags="-w -s" -o ./oxcross-origin

install: build
	cp ./oxcross-origin /usr/local/bin/oxcross-origin
	cp ./oxcross-origin.service /etc/systemd/system/oxcross-origin.service
	chmod 644 /etc/systemd/system/oxcross-origin.service
	systemctl daemon-reload

uninstall:
	systemctl stop oxcross-origin
	systemctl disable oxcross-origin
	rm /etc/systemd/system/oxcross-origin.service
	rm /usr/local/bin/oxcross-origin

reinstall: uninstall install