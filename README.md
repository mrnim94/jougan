<h1 align="center" style="border-bottom: none">
    <a href="https://nimtechnology.com/2023/07/02/jougan-project/" target="_blank"><img alt="JouGan" width="120px" src="https://cdn131.picsart.com/329673960061211.png"></a><br>JouGan
</h1>

<p align="center">Monitoring the speed and volume of Kubernetes in a real-life scenario.</p>

## Introduction
JouGan can download the file you request. I calculate and report:   
- Time taken to download the file (seconds)
- Download speed (MB/s)
- Time taken to save the file (seconds)
- Save speed (MB/s)
- Time taken to delete the file (seconds)
- Delete speed (MB/s)

## Install Jougan on K8s   

Add repository:   
```shell
helm repo add jougan https://mrnim94.github.io/jougan
```

Install chart:
```shell
helm install my-jougan jougan/jougan --version x.x.x
```

### value file   

Example:   
#### Measure Disk Speed any file from Download URL
```yaml
envVars:
  DOWNLOAD_URL: "https://files.testfile.org/PDF/10MB-TESTFILE.ORG.pdf"
  SAVE_TO_LOCATION: /app/downloaded/dynamicSize.bin
nodeSelector:
  kubernetes.io/os: linux
podAnnotations:
  prometheus.io/scrape: 'true'
  prometheus.io/path: '/metrics'
  prometheus.io/port: '1994'
volumes:
  - name: file-service
    persistentVolumeClaim:
    claimName: pvc-file-service-smb-1
volumeMounts:
  - mountPath: /app/downloaded
    name: file-service
```

#### Measure Disk Speed any file on S3 (AWS)   
```yaml
envVars:
  AWS_REGION: us-west-2
  DOWNLOAD_FROM_S3_BUCKET: "ahihi-09262023"
  DOWNLOAD_FROM_S3_KEY: "10MB-TESTFILE.ORG.pdf"
  SAVE_TO_LOCATION: "/app/downloaded/dynamicSize.bin"
  AWS_ACCESS_KEY_ID: "XXXXXGKBQ65KXXXXXX"
  AWS_SECRET_ACCESS_KEY: "xxxxxxx//1dSxxxxxxxJ7nkIrxxxxxxx"
nodeSelector:
  kubernetes.io/os: linux
podAnnotations:
  prometheus.io/scrape: 'true'
  prometheus.io/path: '/metrics'
  prometheus.io/port: '1994'
volumes:
  - name: file-service
    persistentVolumeClaim:
    claimName: pvc-file-service-smb-1
volumeMounts:
  - mountPath: /app/downloaded
    name: file-service
```

### Yaml deployment:   
You install quickly Jougan

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app.kubernetes.io/instance: jougan
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: jougan
    app.kubernetes.io/version: v0.0.2
    helm.sh/chart: jougan-0.1.1
  name: jougan
  namespace: mdaas-engines-dev
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/instance: jougan
      app.kubernetes.io/name: jougan
  template:
    metadata:
      annotations:
        prometheus.io/path: /metrics
        prometheus.io/port: '1994'
        prometheus.io/scrape: 'true'
      labels:
        app.kubernetes.io/instance: jougan
        app.kubernetes.io/managed-by: Helm
        app.kubernetes.io/name: jougan
        app.kubernetes.io/version: v0.0.2
        helm.sh/chart: jougan-0.1.1
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            - labelSelector:
                matchExpressions:
                  - key: app.kubernetes.io/name
                    operator: In
                    values:
                      - jougan
              topologyKey: kubernetes.io/hostname
      containers:
        - env:
            - name: AWS_ACCESS_KEY_ID
              value: XXXXXGKBQ65KXXXXXX
            - name: AWS_REGION
              value: us-west-2
            - name: AWS_SECRET_ACCESS_KEY
              value: xxxxxxx//1dSxxxxxxxJ7nkIrxxxxxxx
            - name: DEBUG_LOG
              value: 'false'
            - name: DOWNLOAD_FROM_S3_BUCKET
              value: ahihi-09262023
            - name: DOWNLOAD_FROM_S3_KEY
              value: 200MB-TESTFILE.ORG.pdf
            - name: SAVE_TO_LOCATION
              value: /app/downloaded/dynamicSize.bin
          image: 'mrnim94/jougan:v0.0.3'
          imagePullPolicy: IfNotPresent
          livenessProbe:
            httpGet:
              path: /
              port: http
          name: jougan
          ports:
            - containerPort: 1994
              name: http
              protocol: TCP
          readinessProbe:
            httpGet:
              path: /
              port: http
          resources: {}
          securityContext: {}
          volumeMounts:
            - mountPath: /app/downloaded
              name: file-service
      nodeSelector:
        kubernetes.io/os: linux
      securityContext: {}
      serviceAccountName: jougan
      volumes:
        - name: file-service
          persistentVolumeClaim:
            claimName: pvc-file-service-smb-1
```
## Grafana

Links: https://grafana.com/grafana/dashboards/20013-jougan-measure-disk-speed/

<a href="https://nimtechnology.com/2023/07/02/jougan-project/" target="_blank"><img alt="JouGan" src="https://grafana.com/api/dashboards/20013/images/15212/image"></a>

## Create Helm Chart
helm package ./helm-chart/jougan --destination ./helm-chart/   
helm repo index . --url https://mrnim94.github.io/jougan   