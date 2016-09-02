#!/bin/bash
set -e

export DEBIAN_FRONTEND=noninteractive
VERSION=0.3.1

apt-get update
apt-get install -y \
    wget

cd /tmp
wget -O collectd_exporter.tar.gz \
    https://github.com/prometheus/collectd_exporter/releases/download/${VERSION}/collectd_exporter-${VERSION}.linux-amd64.tar.gz
tar -zxf collectd_exporter.tar.gz
cd collectd_exporter-${VERSION}.linux-amd64
mv collectd_exporter /bin/collectd_exporter
chmod +x /bin/collectd_exporter