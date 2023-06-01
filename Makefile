.PHONY: all
.SILENT: clean

all:
	mkdir build/{win,linux} -p
	go build -o build/linux/dem ./cmd/...

clean:
	rm -rf *.jpg *.png *.jpeg build
