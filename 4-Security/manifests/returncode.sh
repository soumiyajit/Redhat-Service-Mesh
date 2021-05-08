while :
do
    curl $INGRESS_HOST -s -o /dev/null -w "%{http_code}\n"
    sleep 5
done
