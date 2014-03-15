build:
	go-bindata -nomemcopy -pkg fileindex templates/
	gofmt -s -l -w .
