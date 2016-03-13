# How to generate auth_token

## generate RSA key
generate private key
```openssl genrsa -out private.pem 256```

generate public key
```openssl rsa -in private.pem -pubout -out public.pem```

## update server key
edit rsa.go, update the private key and pubic key,
then rebuild server

## update generator
edit gentoken.go, update the private key and pubic key

## generate auth_token
run
```go run gentoken.go -u new_user_name```
then, you can use the new token and new server!
