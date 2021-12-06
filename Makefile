push:
	sudo docker build -t nightsilvertech/clockwerk:2.0.0 .
	sudo docker tag nightsilvertech/clockwerk:2.0.0 nightsilvertech/clockwerk:2.0.0
	sudo docker push nightsilvertech/clockwerk:2.0.0

docker run -d --name clockwerk -p 1929:1929 -e SCHEDULER_USERNAME=clockwerk -e SCHEDULER_PASSWORD=password -e REDIS_HOST=35.219.50.46 -e REDIS_PORT=6379 -e REDIS_PASS=root  nightsilvertech/clockwerk:2.0.0