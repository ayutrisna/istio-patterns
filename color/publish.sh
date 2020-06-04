docker build . -t itwonderlab/color
docker push itwonderlab/color
kubectl apply -f color-blue.yaml
kubectl apply -f color-green.yaml
kubectl apply -f color-nodeport.yaml
