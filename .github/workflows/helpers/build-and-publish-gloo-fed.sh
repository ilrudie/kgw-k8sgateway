# !/bin/bash
# --------------------------------------
# This script builds and publishes a gloo-fed release...
#   * images:           https://quay.io/organization/solo-io
#   * helm chart:       https://console.cloud.google.com/storage/browser/gloo-ee-test-helm
# ...based on the current branch
# --------------------------------------
BRANCH=$(echo $(git rev-parse --abbrev-ref HEAD) | tr -d '0123456789/.')
VERSION=$(go run .github/workflows/helpers/find-latest-local-version.go)-b$BRANCH-$(git rev-parse --short HEAD)
# for example: VERSION=$(1.14.0-beta2)-b$(-one-click)-$(7c8df00ef)
HELM_BUCKET="gs://gloo-ee-test-helm"
HELM_SYNC_DIR="_output/helm_gloo_fed"
HELM_DIR="install/helm"

# build and push images
VERSION=$VERSION make gloofed-docker -B
VERSION=$VERSION TAGGED_VERSION=$VERSION make docker-push-fed -B

# create appropriate Values.yaml and Chart.yaml files
VERSION=$VERSION make init-helm

# Complicated block ripped from gloo-ee Makefile.  Roughly, this block...
#   1. Grabs GENERATION id of helm repo
#   2. Downloads helm index.yaml
#   3. Packages local helm chart and merges it into the local index.yaml
#   4. Uploads the local helm chart and index.yaml to the helm repo
#   5. If the helm repo has been updated since the last download, the upload will fail
until $(GENERATION=$(gsutil ls -a $HELM_BUCKET/index.yaml | tail -1 | cut -f2 -d '#') && \
                gsutil cp -v $HELM_BUCKET/index.yaml $HELM_SYNC_DIR/index.yaml && \
                helm package --destination $HELM_SYNC_DIR/charts $HELM_DIR/gloo-fed >> /dev/null && \
                helm repo index $HELM_SYNC_DIR --merge $HELM_SYNC_DIR/index.yaml && \
                gsutil -m rsync $HELM_SYNC_DIR/charts $HELM_BUCKET/charts && \
                gsutil -h x-goog-if-generation-match:"$GENERATION" cp $HELM_SYNC_DIR/index.yaml $HELM_BUCKET/index.yaml); do \
    echo "Failed to upload new helm index (updated helm index since last download?). Trying again"; \
    sleep 2; \
done

# provide (hopefully) useful output
echo "Successfully published a test build of gloo-fed!" > published-gloo-fed.txt
echo "  Version:    $VERSION" >> published-gloo-fed.txt
echo "  Helm Repo:  https://console.cloud.google.com/storage/browser/gloo-ee-test-helm" >> published-gloo-fed.txt
echo "  Image Repo: https://quay.io/organization/solo-io" >> published-gloo-fed.txt
echo "Can Install Via:" >> published-gloo-fed.txt
echo "❯ helm repo add gloo-fed-test https://storage.googleapis.com/gloo-ee-test-helm" >> published-gloo-fed.txt
echo "❯ helm repo update" >> published-gloo-fed.txt
echo "❯ helm install -n gloo-system gloo-fed-test gloo-fed-test/gloo-fed --create-namespace --version $VERSION --set-string license_key=\$GLOO_LICENSE_KEY" >> published-gloo-fed.txt
cat published-gloo-fed.txt
