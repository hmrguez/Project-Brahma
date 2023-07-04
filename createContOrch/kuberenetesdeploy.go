package createContOrch

var kubernetesdeploy =
`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-microservice
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-microservice
  template:
    metadata:
      labels:
        app: my-microservice
    spec:
      containers:
        - name: my-microservice
          image: my-registry/my-microservice:1.0.0
          ports:
            - containerPort: 8080
`
