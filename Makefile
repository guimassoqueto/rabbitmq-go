env:
	cat .env.sample 1> .env

ih:
	npm install && npx husky install

rmq:
	docker compose up rabbitmq -d
