# Notification Service

This service is a RESTful API intended for sending notification particularly payment notification to the certail url specified by customer. Mainly this service has APIs such as :

- Generate API Key
- Register Notification URL
- Update Notification URL
- Activate Notification URL
- Send Notification
- Retry Send Notification


## Getting Started

- Create Go projects path 

```
$ cd <YOUR_WORK_PATH>
$ mkdir src bin pkg
$ cd src
$ mkdir github.com
$ cd github.com
$ mkdir liemeilla
$ cd liemeilla
```

- Clone the repository
```
git clone ...
```

- Running the API
```
$ make run_api
```

- Running the tests
```
$ make run_test
```

