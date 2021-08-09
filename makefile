.PHONY:	demo

demo:
	@export GO111MODULE=on && \
	export GOPROXY=https://goproxy.io && \
	go build demo.go
	@chmod 777 demo


.PHONY: clean
clean:
	@rm -rf demo
	@echo "[clean Done]"
