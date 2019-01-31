docker build -t luksa/fortune:args .
docker build -t luksa/fortune:env .

kubectl create configmap fortune-config --from-file=configmap-files

kubectl create secret generic fortune-https --from-file=fortune-https
