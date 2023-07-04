package createContOrch

var kuberneteslb = `
apiVersion: v1
kind: Service
metadata:
  name: my-load-balancer
spec:
  selector:
    app: my-app
  ports:
  - name: http
    port: 80
    targetPort: 8080
  type: LoadBalancer
`
