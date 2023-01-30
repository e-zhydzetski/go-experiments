#!/bin/ash

tcpdump -i lo --immediate-mode -w dump.cap &
tcpdump_pid=$!
sleep 1

go test ./...

kill $tcpdump_pid
wait $tcpdump_pid