#!/bin/sh -eux

appcfg.py -A $PROJECT_ID --version 1 --oauth2_access_token=$ACCESS_TOKEN --oauth2_refresh_token=$REFRESH_TOKEN update app
