while :
do
    http get $INGRESS_HOST/api/v1/hello | grep "Version"
    sleep 1
done
