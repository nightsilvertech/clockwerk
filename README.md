## How to run clockwerk
Before you run clockwerk image you need to set up redis database for store scheduler data, after your 
redis ready to use you can continue to execute this command for running the clockwerk

```bash
docker run --name redis-clockwerk -p 6379:6379 -d redis redis-server --save 60 1 --loglevel warning --requirepass redis123
```

```bash
docker run -d --name clockwerk -p 1929:1929 -e SCHEDULER_USERNAME=clockwerk -e SCHEDULER_PASSWORD=1234 -e REDIS_HOST=localhost -e REDIS_PORT=6379 -e REDIS_PASS=redis123  nightsilvertech/clockwerk:2.0.0
```