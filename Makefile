all:
	$(MAKE) -C src all

build:
	$(MAKE) -C src build

clean:
	$(MAKE) -C src clean

.PHONY: all build clean
