#!/bin/bash
# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Script to deploy Recon service to a GKE cluster.
#
# Usage:
#
# ./deploy_key.sh <"prod"|"autopush"> <commit_hash>
#
# First argument is either "prod" or "autopush"
# (Optional) second argument is the git commit hash of the reconciliation repo.
#
# !!! WARNING: Run this script in a clean Git checkout at the desired commit.
#
# This retrives the docker images and gRPC descriptor based on git commit hash,
# so these binaries should have been pushed to container registry and Cloud
# Storage by the continous deployment flow (../build/ci/cloudbuild.push.yaml).

set -e

ENV=$1

if [[ $ENV != "prod" && $ENV != "autopush" ]]; then
  echo "First argument should be 'prod' or 'autopush'"
  exit
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT="$(dirname "$DIR")"

TAG=$(git rev-parse --short HEAD)
if [[ $2 != "" ]]; then
  TAG=$2
  cd "$ROOT"
  # This is important to get the correct Bigtable version
  git checkout "$TAG"
fi

cd "$ROOT/deploy/git"
echo -n "$TAG" > recon_hash.txt

cd $ROOT

if [[ $ENV == "autopush" ]]; then
  # Update bigtable version
  gsutil cp gs://datcom-control/latest_base_cache_version.txt deploy/storage/bigtable.version
fi

PROJECT_ID=$(yq read deploy/gke/$ENV.yaml project)
REGION=$(yq read deploy/gke/$ENV.yaml region)
IP=$(yq read deploy/gke/$ENV.yaml ip)
DOMAIN=$(yq read deploy/gke/$ENV.yaml domain)
API_TITLE=$(yq read deploy/gke/$ENV.yaml api_title)
CLUSTER_NAME=recon-$REGION

cd $ROOT/deploy/overlays/$ENV

# Deploy to GKE
kustomize edit set image gcr.io/datcom-ci/datacommons-recon=gcr.io/datcom-ci/datacommons-recon:$TAG
kustomize build > $ENV.yaml
gcloud config set project $PROJECT_ID
gcloud container clusters get-credentials $CLUSTER_NAME --region $REGION
kubectl apply -f $ENV.yaml

# Deploy Cloud Endpoints
yq w --style=double $ROOT/esp/endpoints.yaml.tmpl name $DOMAIN > endpoints.yaml
yq w -i endpoints.yaml title "$API_TITLE"
yq w -i endpoints.yaml endpoints[0].target "$IP"
yq w -i endpoints.yaml endpoints[0].name "$DOMAIN"
gsutil cp gs://datcom-public/recon-grpc/recon-grpc.$TAG.pb .
gcloud endpoints services deploy recon-grpc.$TAG.pb endpoints.yaml --project $PROJECT_ID
