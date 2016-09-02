#!/bin/bash
set -e
# From official grafana docker image https://github.com/grafana/grafana-docker

export DEBIAN_FRONTEND=noninteractive
GRAFANA_VERSION=3.1.1-1470047149

apt-get update
apt-get -y --no-install-recommends install \
    libfontconfig \
    curl \
    ca-certificates \
    jq # needed by omnia
apt-get clean
curl https://grafanarel.s3.amazonaws.com/builds/grafana_${GRAFANA_VERSION}_amd64.deb > /tmp/grafana.deb
dpkg -i /tmp/grafana.deb
rm /tmp/grafana.deb
curl -L https://github.com/tianon/gosu/releases/download/1.7/gosu-amd64 > /usr/sbin/gosu
chmod +x /usr/sbin/gosu

# [omnia] Merging dashboard panels
PANELS=$(find /conf/metrics -type f -name 'panel.json')
DASHBOARD=$(cat /conf/common/dashboard.json)
for panel in $PANELS
do
    DASHBOARD=$(echo $DASHBOARD | jq --slurpfile panel $panel '(.dashboard.rows[] | select(.title=="Default Row") | .panels) += $panel')
done
echo $DASHBOARD > /conf/merged_dashboard.json