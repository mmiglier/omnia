#!/bin/bash
set -e

export DEBIAN_FRONTEND=noninteractive

apt-get update
apt-get install -y \
	default-jre \
	wget

wget -O /tmp/riemann.deb https://aphyr.com/riemann/riemann_0.2.11_all.deb
dpkg -i /tmp/riemann.deb

mv /conf/common/riemann.config /etc/riemann/riemann.config
mkdir -p /etc/riemann/omnia/etc/
mv /conf/common/*.clj /etc/riemann/omnia/etc/