FROM uselagoon/nginx:25.5.0

COPY index.html /app/.
COPY api.html /app/.
COPY cms.html /app/.
COPY .lagoon.yml /app/.
COPY app.conf /etc/nginx/conf.d/app.conf