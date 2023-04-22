include .env

run up:
	./script/app.sh run

stop down:
	./script/app.sh stop

deps:
	docker compose -f ./tools/compose.yaml run --rm go-mod

logs-%:
	./script/app.sh logs $*

dashboard:
	./script/app.sh dashboard

destroy:
	./script/app.sh destroy
