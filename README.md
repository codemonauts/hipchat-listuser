# hipchat-listuser

Connect to your HipChat account and print all useraccounts sorted descending by the "Last active" field.
This helps to identify inactive user and therefore to reduce your costs.


## Installation
```
dep ensure
go build
./hipchat-listuser -apikey "yourapikey"
```