# Set the base image to centos6.8
FROM sunnysmilez/php

ENV APP_NAME=test

#copy项目代码
COPY ./code /data/htdocs/test/
EXPOSE 9000
ENTRYPOINT ["docker-entrypoint.sh"]