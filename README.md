# hipchat-listuser

Connect to your HipChat account and print all useraccounts sorted descending by the "Last active" field.
This helps to identify inactive user and therefore to reduce your costs.

## Get an API Key
To generate an API key for your account, go to your 'Account Settings' under `https://<instance-name>.hipchat.com/account/api`.

## Installation
```
dep ensure
go build
./hipchat-listuser -apikey "yourapikey"
```