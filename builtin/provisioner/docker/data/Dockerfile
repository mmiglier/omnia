FROM ubuntu:16.04
MAINTAINER Marco Miglierina <marco.miglierina@icloud.com>

COPY conf /conf

COPY setup.sh /tmp/setup.sh
RUN chmod +x /tmp/setup.sh
RUN /tmp/setup.sh

{% if .Ports.Expose %}{% range .Ports.Expose -%}
EXPOSE {% . %}
{% end %}{% end -%}
{% if .Ports.Publish %}{% range .Ports.Publish -%}
EXPOSE {% . %}
{% end %}{% end -%}

COPY run.sh /usr/local/bin/run.sh
RUN chmod +x /usr/local/bin/run.sh

ENTRYPOINT ["run.sh"]