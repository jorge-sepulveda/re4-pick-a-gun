run-ui:
	go run ui/ui.go ui/bundled.go

run-cli:
	go run cli/cli.go

build-mac:
	fyne package -os darwin -icon icon.png

build-windows:
	fyne-cross windows -output re4-pick-a-gun -arch=amd64

