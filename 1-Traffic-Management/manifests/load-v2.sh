while :
do
    http get $INGRESS_HOST/api/v2/hello | grep "Version"
    sleep 1
done
