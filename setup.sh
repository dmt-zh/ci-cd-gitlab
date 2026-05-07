#!/bin/bash

ENV_FILE=".env"

###############################################################################

set_env_var() {
    local name=$1
    local value=$2
    sed -i "s|^[[:space:]]*${name}[[:space:]]*=.*|${name}=${value}|" "$ENV_FILE"
}

###############################################################################

set_env_var "ROOT_PASSWORD" $(openssl rand -base64 32 | tr -d /=+ | cut -c1-30)
set_env_var "REGISTRATION_TOKEN" "$(openssl rand -hex 32)"

###############################################################################

echo -e "\033[0;32m✔\033[0m Credentials successfully updated in \".env\" file!"

###############################################################################
