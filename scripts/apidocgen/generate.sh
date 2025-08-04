#!/usr/bin/env bash

# This script generates swagger description file and required Go docs files
# from the coderd API.

set -euo pipefail
# shellcheck source=scripts/lib.sh
source "$(dirname "$(dirname "${BASH_SOURCE[0]}")")/lib.sh"

APIDOCGEN_DIR=$(dirname "${BASH_SOURCE[0]}")
API_MD_TMP_FILE=$(mktemp /tmp/coder-apidocgen.XXXXXX)

cleanup() {
        rm -f "${API_MD_TMP_FILE}"
}
trap cleanup EXIT

log "Use temporary file: ${API_MD_TMP_FILE}"

pushd "${PROJECT_ROOT}"
go run github.com/swaggo/swag/cmd/swag@v1.8.9 init \
        --generalInfo="coderd.go" \
        --dir="./coderd,./codersdk,./enterprise/coderd,./enterprise/wsproxy/wsproxysdk" \
        --output="./coderd/apidoc" \
        --outputTypes="go,json" \
        --parseDependency=true
popd

pushd "${APIDOCGEN_DIR}"

# Make sure that openapi-generator-cli is installed correctly.
pnpm exec -- openapi-generator-cli version
# Render the Markdown file.
TMP_OUTPUT_DIR=$(mktemp -d /tmp/coder-openapi-gen.XXXXXX)
pnpm exec -- openapi-generator-cli generate \
        -i "../../coderd/apidoc/swagger.json" \
        -g markdown \
        -o "${TMP_OUTPUT_DIR}" \
        --additional-properties=skipOperationExample=true

# Combine all API files into a single file with section markers
echo "" > "${API_MD_TMP_FILE}"
for api_file in "${TMP_OUTPUT_DIR}/Apis/"*.md; do
    if [ -f "$api_file" ]; then
        api_name=$(basename "$api_file" .md)
        echo "<!-- APIDOCGEN: BEGIN SECTION -->" >> "${API_MD_TMP_FILE}"
        echo "# ${api_name}" >> "${API_MD_TMP_FILE}"
        echo "" >> "${API_MD_TMP_FILE}"
        # Skip the first line (which is the title) and add the rest
        tail -n +2 "$api_file" >> "${API_MD_TMP_FILE}"
        echo "" >> "${API_MD_TMP_FILE}"
    fi
done

# Clean up temporary directory
rm -rf "${TMP_OUTPUT_DIR}"
# Perform the postprocessing
go run postprocess/main.go -in-md-file-single "${API_MD_TMP_FILE}"
popd
