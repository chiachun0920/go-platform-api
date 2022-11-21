# PlatformAPI

## Pre-requites installation
1. docker
2. docker-compose

## Start Dev Server
```bash
# start mongo in docker firstly
make run
```

## Test Endpoint
```bash
curl localhost:8080/health

# if you get response {"status":"ok"}, represent it work.
```

## Function
### Receive message from line
1. register web hook url, web hook endpoint will be `http://{endpoint}/webhook/line`
![](Screenshot%202022-11-21%20at%201.35.29%20PM.png)
2. copy **channel secret** and **access token** to project/config.json
![](Screenshot%202022-11-21%20at%201.37.24%20PM.png)
3. restart server again
```bash
make run
```
4. Send message from your line

### Send back message to customer
1. you can get customerId from
	* collection: `message`
	* field: `sender`

```bash
# curl http://localhost:8080/messaging/line -X POST -d '{"To": "Ue289fb7b4854f46458a934d286cb178d", "Text": "hello"}'

curl http://localhost:8080/messaging/line -X POST -d '{"To": "$customerId", "Text": "hello"}'
```

### List messages of customer

```bash
curl http://localhost:8080/message/customer/{$customerId}

# curl http://localhost:8080/message/customer/Ue289fb7b4854f46458a934d286cb178d
```