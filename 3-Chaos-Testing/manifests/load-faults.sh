while :
do
    http get $INGRESS_HOST/api/hello 
    sleep 5
done
