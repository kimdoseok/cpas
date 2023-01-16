FROM rabbitmq:management-alpine
ADD etc/rabbitmq.config /etc/rabbitmq/
ADD etc/definitions.json /etc/rabbitmq/
ADD etc/enabled_plugins /etc/rabbitmq/
RUN chown rabbitmq:rabbitmq /etc/rabbitmq/rabbitmq.config /etc/rabbitmq/definitions.json /etc/rabbitmq/enabled_plugins
#CMD ["rabbitmq-server"]