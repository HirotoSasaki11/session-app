## Deployment Manager
gcloud deployment-manager deployments create ${deploymentName} --config environment.yaml --project sandbox-hirotosasaki --preview 
DMではmemcachedのリソースがサポートされていないので、手動で作成する

https://cloud.google.com/deployment-manager/docs/configuration/supported-gcp-types