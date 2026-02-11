#!/bin/bash
set -e

# Compare container benchmark strategies using benchstat

cd "$(dirname "$0")"

COUNT=${1:-6}
echo "Running container benchmarks with $COUNT iterations each..."
go test -bench=BenchmarkContainer -benchtime=3x -count=$COUNT -run=^$ . 2>&1 > container.txt

echo ""
echo "=== Results ==="
go tool benchstat -col /container container.txt

echo "Running reuse benchmarks with $COUNT iterations each..."
go test -bench=BenchmarkReuse -benchtime=3x -count=$COUNT  -run=^$ . 2>&1 > reuse.txt

echo ""
echo "=== Results ==="
go tool benchstat -col /reuse reuse.txt
