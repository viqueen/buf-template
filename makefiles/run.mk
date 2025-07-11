.PHONY: backend-run
backend-run: # Run the application
ifeq ($(shell which modd),)
	go install github.com/cortesi/modd/cmd/modd
endif
	modd -f backend/modd.conf

.PHONY: backend-debug
backend-debug: # Run the application in debug mode
ifeq ($(shell which modd),)
	go install github.com/cortesi/modd/cmd/modd
endif
	modd -f backend/modd-debug.conf
