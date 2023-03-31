compile-percolation-stats:
	CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -ldflags '-extldflags "-static"' -o ../../bin/percolation_stats ./percolation_stats.go

compile-percolation-stats-mac:
	CGO_ENABLED=0 GOOS=darwin GO111MODULE=on go build -a -ldflags '-extldflags "-static"' -o ../../bin/percolation_stats.mac ./percolation_stats.go

compile-percolation-stats-win:
	CGO_ENABLED=0 GOOS=windows GO111MODULE=on go build -a -ldflags '-extldflags "-static"' -o ../../bin/percolation_stats.win ./percolation_stats.go

