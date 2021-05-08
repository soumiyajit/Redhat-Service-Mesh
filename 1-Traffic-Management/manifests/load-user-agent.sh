while :
do
    http get $INGRESS_HOST/api/hello User-Agent:Chrome | grep "Version"
    sleep 1
done
