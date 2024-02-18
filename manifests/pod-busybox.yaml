apiVersion: v1
kind: Pod
metadata:
  name: busyboxtest
  labels:
    app: busyboxtest
spec:
  containers:
  - image: busybox
    command:
      - sleep
      - "3600"
    imagePullPolicy: IfNotPresent
    name: busybox
  restartPolicy: Always