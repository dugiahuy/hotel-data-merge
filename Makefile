.PHONY: dev

include .env
export $(shell sed 's/=.*//' .env)

dev:
	ENV=local go run cmd/*.go