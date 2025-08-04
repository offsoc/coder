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

# Map openapi-generator-cli filenames to original filenames
declare -A filename_map=(
    ["AgentsApi"]="Agents"
    ["ApplicationsApi"]="Applications"
    ["AuditApi"]="Audit"
    ["AuthorizationApi"]="Authorization"
    ["BuildsApi"]="Builds"
    ["DebugApi"]="Debug"
    ["EnterpriseApi"]="Enterprise"
    ["FilesApi"]="Files"
    ["GeneralApi"]="General"
    ["GitApi"]="Git"
    ["InsightsApi"]="Insights"
    ["MembersApi"]="Members"
    ["NotificationsApi"]="Notifications"
    ["OrganizationsApi"]="Organizations"
    ["PortSharingApi"]="PortSharing"
    ["PrebuildsApi"]="Prebuilds"
    ["ProvisioningApi"]="Provisioning"
    ["TemplatesApi"]="Templates"
    ["UsersApi"]="Users"
    ["WorkspaceProxiesApi"]="WorkspaceProxies"
    ["WorkspacesApi"]="Workspaces"
)

for api_file in "${TMP_OUTPUT_DIR}/Apis/"*.md; do
    if [ -f "$api_file" ]; then
        openapi_name=$(basename "$api_file" .md)
        # Use mapped name if available, otherwise use original name
        section_name=${filename_map[$openapi_name]:-$openapi_name}
        echo "<!-- APIDOCGEN: BEGIN SECTION -->" >> "${API_MD_TMP_FILE}"
        echo "# ${section_name}" >> "${API_MD_TMP_FILE}"
        echo "" >> "${API_MD_TMP_FILE}"
        # Skip the first line (which is the title) and add the rest
        tail -n +2 "$api_file" >> "${API_MD_TMP_FILE}"
        echo "" >> "${API_MD_TMP_FILE}"
    fi
done

# Add static authentication section
echo "<!-- APIDOCGEN: BEGIN SECTION -->" >> "${API_MD_TMP_FILE}"
echo "# Authentication" >> "${API_MD_TMP_FILE}"
echo "" >> "${API_MD_TMP_FILE}"
echo "Long-lived tokens can be generated to perform actions on behalf of your user account:" >> "${API_MD_TMP_FILE}"
echo "" >> "${API_MD_TMP_FILE}"
echo '```shell' >> "${API_MD_TMP_FILE}"
echo 'coder tokens create' >> "${API_MD_TMP_FILE}"
echo '```' >> "${API_MD_TMP_FILE}"
echo "" >> "${API_MD_TMP_FILE}"
echo "You can use tokens with the Coder's REST API using the \`Coder-Session-Token\` HTTP header." >> "${API_MD_TMP_FILE}"
echo "" >> "${API_MD_TMP_FILE}"

# Add schemas section from Models if it exists
if [ -d "${TMP_OUTPUT_DIR}/Models" ] && [ "$(ls -A "${TMP_OUTPUT_DIR}/Models")" ]; then
    echo "<!-- APIDOCGEN: BEGIN SECTION -->" >> "${API_MD_TMP_FILE}"
    echo "# Schemas" >> "${API_MD_TMP_FILE}"
    echo "" >> "${API_MD_TMP_FILE}"
    echo "## Models" >> "${API_MD_TMP_FILE}"
    echo "" >> "${API_MD_TMP_FILE}"
    for model_file in "${TMP_OUTPUT_DIR}/Models/"*.md; do
        if [ -f "$model_file" ]; then
            model_name=$(basename "$model_file" .md)
            echo "### ${model_name}" >> "${API_MD_TMP_FILE}"
            echo "" >> "${API_MD_TMP_FILE}"
            # Skip the first line (which is the title) and add the rest
            tail -n +2 "$model_file" >> "${API_MD_TMP_FILE}"
            echo "" >> "${API_MD_TMP_FILE}"
        fi
    done
fi

# Clean up temporary directory
rm -rf "${TMP_OUTPUT_DIR}"
# Perform the postprocessing
go run postprocess/main.go -in-md-file-single "${API_MD_TMP_FILE}"
popd
