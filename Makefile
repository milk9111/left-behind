all: build web

build:
	go build -o trixie-the-truffler.exe

web:
	wasmnow -b
	tar -caf _web.zip wasmnow

web-serve:
	wasmnow