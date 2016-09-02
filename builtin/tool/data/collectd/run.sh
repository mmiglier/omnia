#!/bin/bash

set -e

# if collectd is launched in a container proc and hostname
# can be mounted on /mnt/proc and /mnt/hostname
if [ -d /mnt/proc ]; then
  umount /proc
  mount -o bind /mnt/proc /proc
fi
if [ -e /mnt/hostname ]; then
    HOST_HOSTNAME=$(cat /mnt/hostname)
    if [ -n "$HOST_HOSTNAME" ]; then
        printf "\nHostname \"$HOST_HOSTNAME\"\n" >> /opt/collectd/etc/collectd.conf
    fi
fi

if [ -z "$@" ]; then
  exec /opt/collectd/sbin/collectd -C /opt/collectd/etc/collectd.conf -f
else
  exec $@
fi