while :
do
    http get $INGRESS_HOST/api/hello v2-header:canary | grep "Version"
    sleep 1
done
