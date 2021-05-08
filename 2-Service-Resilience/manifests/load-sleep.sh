while :
do
    http get $INGRESS_HOST/api/hello  sleep==2
    sleep 5
done
