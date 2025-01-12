all: build web

build:
	go build -o trixie-the-truffler.exe
	butler push trixie-the-truffler.exe milk9111/trixie-the-truffler:windows

web:
	wasmnow -b
	tar -caf _web.zip wasmnow
	butler push _web.zip milk9111/trixie-the-truffler:html5

web-serve:
	wasmnow