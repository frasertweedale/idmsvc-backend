#!/bin/bash

function error {
    local err=$?
    printf "%s\n" "$*" >&2
    exit $err
}

INVENTORY_ID=$"$1"
FQDN="$2"
[ "${INVENTORY_ID}" != "" ] || error "INVENTORY_ID is empty"
[ "${FQDN}" != "" ] || error "FQDN is empty"

export X_RH_IDENTITY="$( ./tools/bin/xrhidgen -org-id 12345 system -cn "3f35fc7f-079c-4940-92ed-9fdc8694a0f3" -cert-type system | base64 -w0 )"
export X_RH_IDM_VERSION='{"ipa-hcc": "0.9", "ipa": "4.10.0-8.el9_1", "os-release-id": "rhel", "os-release-version-id": "9.1"}'
unset X_RH_FAKE_IDENTITY
unset CREDS
BASE_URL="http://localhost:8000/api/idmsvc/v1"
./scripts/curl.sh -i -X POST -d '{}' "${BASE_URL}/host-conf/${INVENTORY_ID}/${FQDN}"
