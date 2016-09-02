#!/bin/bash
set -e

export DEBIAN_FRONTEND=noninteractive

apt-get update
apt-get install -y \
    graphite-carbon \
    supervisor \
    build-essential \
    python-dev \
    libffi-dev \
    libcairo2-dev \
    python-pip

pip install --upgrade pip
pip install gunicorn graphite-api[sentry,cyanite]

cd /conf/common

# Graphite API
mv graphite-api.yaml /etc/graphite-api.yaml

# Graphite
mv carbon.conf /etc/carbon/carbon.conf
mv storage-schemas.conf /etc/carbon/storage-schemas.conf
mv storage-aggregation.conf /etc/carbon/storage-aggregation.conf

# Supervisord
mv supervisord.conf /etc/supervisor/conf.d/supervisord.conf