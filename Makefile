build-image:
	docker build -t victor-web-dev/financeapp .

run-app:
	docker-compose -f .devops/docker-compose.yml up -d

shutdown-app:
	docker-compose -f .devops/docker-compose.yml down