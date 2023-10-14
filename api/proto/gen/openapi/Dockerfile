FROM swaggerapi/swagger-ui

COPY *.yaml /usr/share/nginx/html/

ARG API_BASE_URL
RUN sed -i "s|__API_BASE_URL__|${API_BASE_URL}|g" /usr/share/nginx/html/*.yaml

ENV URLS \
    "[ \
        { url: \"/auth.swagger.yaml\", name: \"auth\" }, \
        { url: \"/speech.swagger.yaml\", name: \"speech\" }, \
        { url: \"/video.swagger.yaml\", name: \"video\" } \
    ]"
ENV URLS_PRIMARY_NAME auth