push:
	sudo docker build -t nightsilvertech/clockwerk:2.0.0 .
	sudo docker tag nightsilvertech/clockwerk:2.0.0 nightsilvertech/clockwerk:2.0.0
	sudo docker push nightsilvertech/clockwerk:2.0.0