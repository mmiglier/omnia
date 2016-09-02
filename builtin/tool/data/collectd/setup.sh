#!/bin/bash
set -e

export DEBIAN_FRONTEND=noninteractive

COLLECTD_VERSION=5.5.2
COLLECTD_FIELD_MAX_LEN=1024

apt-get update
apt-get install -y \
		wget \
		unzip \
		python-pip \
		libtool \
		python-dev \
		protobuf-c-compiler
cd /tmp
wget https://collectd.org/files/collectd-${COLLECTD_VERSION}.tar.bz2
tar jxf collectd-${COLLECTD_VERSION}.tar.bz2	
cd collectd-${COLLECTD_VERSION}

# collectd fields have a maximum length of 63 characters by default, we make it tunable
sed -i "s/DATA_MAX_NAME_LEN 64/DATA_MAX_NAME_LEN ${COLLECTD_FIELD_MAX_LEN}/g" src/daemon/plugin.h
./configure
make all install

# install docker collectd plugin if need to monitor docker
# cd /opt/collectd/lib && \
# wget https://github.com/mmiglier/docker-collectd-plugin/archive/master.zip && \
# unzip master.zip && \
# rm master.zip && \
# mv docker-collectd-plugin-master docker && \
# cd docker && \
# pip install -r requirements.txt

mv /conf/common/collectd.conf /opt/collectd/etc/collectd.conf
mv /conf/metrics /opt/collectd/etc/include