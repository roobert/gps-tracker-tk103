all: gps-tracker-tk103-receiver gps-tracker-tk103-ui
gps-tracker-tk103-receiver:
	@go build -o gps-tracker-tk103-receiver cmd/gps-tracker-tk103-receiver/main.go
gps-tracker-tk103-ui:
	@go build -o gps-tracker-tk103-ui cmd/gps-tracker-tk103-ui/main.go
clean:
	@rm -vf gps-tracker-tk103-receiver
	@rm -vf gps-tracker-tk103-ui
