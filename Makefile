
generate:
	go generate ./...

docker:
	docker compose -f ./docker-compose-local.yml up --build --force-recreate

stop-docker:
	docker compose -f ./docker-compose-local.yml down

ngrok:
	ngrok http --url=globally-live-crappie.ngrok-free.app 80

serve:
	go run main.go serve

ts:
	go run main.go ts