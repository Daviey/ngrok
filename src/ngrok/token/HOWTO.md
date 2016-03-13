# How to generate auth_token

## Generate RSA key
Generate private key
```
openssl genrsa -out private.pem 256
```

Generate public key
```
openssl rsa -in private.pem -pubout -out public.pem
```

## Update server key
Edit rsa.go, update the private key and pubic key,
then rebuild server

## Update generator
Edit gentoken.go, update the private key and pubic key

## Generate auth_token
Run
```
go run gentoken.go -u new_user_name
```

## Update client config
Edit config

```
auth_token: xxxxxx
```
then, you can use the new token and new server!
