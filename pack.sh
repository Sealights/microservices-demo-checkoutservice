docker build -t checkoutservice .
docker tag checkoutservice:latest 159616352881.dkr.ecr.eu-west-1.amazonaws.com/microservices-demo-checkoutservice:latest
docker push 159616352881.dkr.ecr.eu-west-1.amazonaws.com/microservices-demo-checkoutservice:latest
