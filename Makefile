push:
	sudo docker build -t nightsilvertech/clockwerk:1.0.0 .
	sudo docker tag nightsilvertech/clockwerk:1.0.0 nightsilvertech/clockwerk:1.0.0
	sudo docker push nightsilvertech/clockwerk:1.0.0