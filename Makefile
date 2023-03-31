.PHONY: all percolation-stats 

ifdef V
TESTARGS = -v -args -alsologtostderr -v 5
else
TESTARGS =
endif

all: percolation-stats

percolation-stats:
	mkdir -p bin
	$(MAKE) -C ./applications/percolation compile-percolation-stats

percolation-stats-mac:
	mkdir -p bin
	$(MAKE) -C ./applications/percolation compile-percolation-stats-mac

