gen-mock:
	mockgen -source ./paginator/iface.go -destination ./paginator/mock_paginator/mock_iface.go
test:
	go test ./paginator