<h1 align="center">Welcome to Clockwerk 🤖</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-2.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-GPLv3-blue.svg" />
  </a>
  <img alt="documentation: yes" src="https://img.shields.io/badge/Documentation-Yes-green.svg" />
  <img alt="maintained: yes" src="https://img.shields.io/badge/Maintained-Yes-green.svg" />
</p>

> A distributed cron service with API for execute HTTP request, with simple lovely dashboard

----

## Installation & Configuration

Before you run clockwerk you need to set up redis database for store scheduler data, after your 
redis ready to use you can continue to execute this command for running the clockwerk
```bash
docker run -d --name clockwerk -p 1929:1929 -e SCHEDULER_USERNAME=clockwerk -e SCHEDULER_PASSWORD=password -e REDIS_HOST=localhost -e REDIS_PORT=6379 -e REDIS_PASS=redis123  nightsilvertech/clockwerk:2.0.0
```

Below is only needed if you don't have redis database running, here is for spinning up redis database with docker
```bash
docker run -d --name redis-clockwerk -p 6379:6379 redis redis-server --save 60 1 --loglevel warning --requirepass redis123
```

Below for spinning up the simple web dashboard for manage the scheduler (please change the env parameter because this is just example)
```bash
docker run -d --name clockwerk-ui -p 1930:1930 -e REACT_APP_SCHEDULER_HOST=localhost -e REACT_APP_SCHEDULER_PORT=1929 -e REACT_APP_SCHEDULER_USERNAME=clockwerk -e REACT_APP_SCHEDULER_PASSWORD=123 nightsilvertech/clockwerk-ui:1.0.0
```

## Usage

### Create Client
```go
// Create clockwerk client
client, err := clockwerk.NewClockwerk("localhost", "1929", "clockwerk", "password")
if err != nil {
    panic(err)
}
```
### SchedulerHTTP & HTTPHeader struct definition
```go
[]clockwerk.HTTPHeader{
    {
        K: "Content-Type",     // http header key
        V: "application/json", // http header value
    },
}
```
```go
clockwerk.SchedulerHTTP{
    Name:           "call get dummy every one minute", // scheduler name
    URL:            "http://localhost:1929/v1/dummy",  // url that scheduler calls
    ReferenceId:    id.String(),                       // reference id is generated as you need (uuid recommended)
    Executor:       clockwerk.HTTP,                    // scheduler executor type (for now just HTTP)
    Method:         clockwerk.GET,                     // http method (available GET POST PUT DELETE)
    Body:           "",                                // body json (prettied with tabulation and new line)
    Spec:           "* * * * *",                       // cron spec (go to https://crontab.guru/ for detail)
    Persist:        true,                              // scheduler should not be deleted if job executed
    Disabled:       false,                             // scheduler should not be running if disabled
    Retry:          15,                                // retry when scheduler unable to execute job
    RetryThreshold: 1,                                 // retry threshold (defined in second unit)
    HTTPHeader:     httpHeader,                        // http header that jobs needed
}
```
### Add scheduler
```go
// Add Scheduler Example
createdScheduler, err = client.Add(clockwerk.SchedulerHTTP{
    Name:           "call get dummy every one minute",
    URL:            "http://localhost:1929/v1/dummy",
    ReferenceId:    id.String(),
    Executor:       clockwerk.HTTP,
    Method:         clockwerk.GET,
    Body:           "",
    Spec:           "* * * * *",
    Persist:        true,
    Disabled:       false,
    Retry:          15,
    RetryThreshold: 1,
    HTTPHeader: []clockwerk.HTTPHeader{
        {K: "Content-Type", V: "application/json"},
    },
})
if err != nil {
    log.Println(err)
}
```

### Toggle scheduler
```go
// Toggle Scheduler Example

// enable scheduler
err := client.Toggle(clockwerk.SchedulerToggle{Id: createdScheduler.Id, ReferenceId: createdScheduler.ReferenceId, Disabled: false})
if err != nil {
    log.Println(err)
}

// disable scheduler
err := client.Toggle(clockwerk.SchedulerToggle{Id: createdScheduler.Id, ReferenceId: createdScheduler.ReferenceId, Disabled: true})
if err != nil {
    log.Println(err)
}
```

### Delete scheduler
```go
// Delete Scheduler Example
err := client.Del(clockwerk.SchedulerSelect{Id: createdScheduler.Id, ReferenceId: createdScheduler.ReferenceId})
if err != nil {
    log.Println(err)
}
```

### Non Golang Usage


----

## Open source licensing info
1. [LICENSE](LICENSE)


----

## Credits and references
1. https://dkron.io/