while :
do
    http get $INGRESS_HOST/api/hello | grep "Version"
    sleep 5
done
