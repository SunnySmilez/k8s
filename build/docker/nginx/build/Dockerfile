FROM nginx

MAINTAINER SunnySmilez "phpchengxuyuan@sina.com"

COPY nginx.conf /etc/nginx/nginx.conf
COPY vhosts /etc/nginx/conf.d/
COPY fastcgi_params /etc/nginx/fastcgi_params

COPY docker-entrypoint.sh /usr/local/bin/
RUN /bin/bash -c 'mkdir -p /data/logs/'

ENTRYPOINT ["docker-entrypoint.sh"]