while :
do
    curl -kv $INGRESS_HOST/api/hello
    sleep 1
done
