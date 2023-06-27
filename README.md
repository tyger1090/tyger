DOCKER:
docker build . -t joska99/tyger:go_app
docker run -p 3000:3000 -it -t joska99/tyger:go_app
docker push joska99/tyger:go_app
localhost:3000


GO_WEB_APP:
go build
./news-demo-starter-files
localhost:3000



I used Prometheus Helm chart:
https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack

apply hel chart:
sudo kubectl create ns prometheus
sudo helm --kubeconfig ~/.kube/config  install prom  prometheus-community/kube-prometheus-stack -n prometheus 