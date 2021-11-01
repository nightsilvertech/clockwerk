deploy:
	podman build -t 35.219.77.34:8082/nbdg-promo/nobita-promo-scheduler:1.0.0-$(shell git rev-parse --short HEAD) .
	podman push 35.219.77.34:8082/nbdg-promo/nobita-promo-scheduler:1.0.0-$(shell git rev-parse --short HEAD)
	podman rm nobita-promo-scheduler-dev -f
	podman run --pod promo-engine --rm --name nobita-promo-scheduler-dev -dt 35.219.77.34:8082/nbdg-promo/nobita-promo-scheduler:1.0.0-$(shell git rev-parse --short HEAD)
	podman image prune -a