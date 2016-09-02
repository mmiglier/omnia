#!/bin/bash
set -e

export DEBIAN_FRONTEND=noninteractive

VERSION=1.0.2

apt-get update
apt-get install -y \
    wget

cd /tmp
wget -O prometheus.tar.gz \
    https://github.com/prometheus/prometheus/releases/download/v${VERSION}/prometheus-${VERSION}.linux-amd64.tar.gz
tar -zxf prometheus.tar.gz
cd prometheus-${VERSION}.linux-amd64
mv prometheus /bin/prometheus
chmod +x /bin/prometheus