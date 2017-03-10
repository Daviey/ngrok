# How to generate auth_token

## Generate RSA key
### Generate private key
```
openssl genrsa -out private.pem 384
```

### Generate public key
```
openssl rsa -in private.pem -pubout -out public.pem
```

## Update server key
Edit **rsa.go**, update the private key and pubic key, then **rebuild server**


## Compile ngrokg binary
In top ngrok directory, run
```
make token-gen
```
or
```
make release-token-gen
```

## Generate auth_token
Run
```
bin/ngrokg -u new_user_name
```
You will receive a new auth_token, like ```auth_token: xxxxxx```

## Update client config
Edit config

```
auth_token: xxxxxx
```
then, you can use the new token and new server!

PS: Once you have generated rsa keys, you don't have to generate them again, just generate auth_token with new user name

## TODO
Currently, this token wouldn't expire, so we need add expired time into token
