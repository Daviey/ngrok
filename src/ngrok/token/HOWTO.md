# How to generate auth_token

## Generate RSA key
### Generate private key
```
openssl genrsa -out private.pem 256
```
PS: the key length determine the user name length,
```key length - 11 = user name length```(in bytes)

### Generate public key
```
openssl rsa -in private.pem -pubout -out public.pem
```
PS: in this case, you shouldn't make the public key opened to others.

## Update server key
Edit **rsa.go**, update the private key and pubic key, then **rebuild server**

## Update generator
Edit gentoken.go, update the private key and pubic key

## Generate auth_token
Run
```
go run gentoken.go -u new_user_name
```
You will receive a new auth_token, like ```auth_token: xxxxxx```

## Update client config
Edit config

```
auth_token: xxxxxx
```
then, you can use the new token and new server!

PS: Once you have generated rsa keys, you don't have to generate them again, just generate auth_token with new user name
