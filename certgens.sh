#!/bin/bash

cur_dir=`pwd`
ngrok_dir=${1:-$cur_dir}
cert_dir=cert

CN=${2:-tunnel.local}

if [ -d $cert_dir ]; then
    echo directory $cert_dir exists
else
    mkdir $cert_dir 
fi

echo ngrok directory: $ngrok_dir
echo CN: $CN

openssl genrsa -out ${ngrok_dir}/${cert_dir}/rootCA.key 2048
openssl req -x509 -new -nodes \
    -key ${ngrok_dir}/${cert_dir}/rootCA.key \
    -subj "/CN=${CN}" \
    -out ${ngrok_dir}/${cert_dir}/rootCA.pem \
    -days 2048
openssl genrsa -out ${ngrok_dir}/${cert_dir}/device.key 2048
openssl req -new -key ${ngrok_dir}/${cert_dir}/device.key \
    -subj "/CN=${CN}" \
    -out ${ngrok_dir}/${cert_dir}/device.csr
openssl x509 -req -in ${ngrok_dir}/${cert_dir}/device.csr \
    -CA ${ngrok_dir}/${cert_dir}/rootCA.pem \
    -CAkey ${ngrok_dir}/${cert_dir}/rootCA.key \
    -CAcreateserial \
    -out ${ngrok_dir}/${cert_dir}/device.crt \
    -days 2048

cp -vf ${ngrok_dir}/${cert_dir}/rootCA.pem $ngrok_dir/assets/client/tls/ngrokroot.crt
cp -vf ${ngrok_dir}/${cert_dir}/device.crt $ngrok_dir/assets/server/tls/snakeoil.crt
cp -vf ${ngrok_dir}/${cert_dir}/device.key $ngrok_dir/assets/server/tls/snakeoil.key
