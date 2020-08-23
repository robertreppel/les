les-node -b Eventsourcing.eml.yaml \
&& cd api \
&& npm install \
&& docker-compose up -d \
&& docker-compose restart api

