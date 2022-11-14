ARG BASE

# gcr.io/distroless/cc
# gcr.io/distroless/static

FROM ${BASE}

# # Any non-zero number will do, and unfortunately a named user will not, as k8s
# # pod securityContext runAsNonRoot can't resolve the user ID:
# # https://github.com/kubernetes/kubernetes/issues/40958. Make root (uid 0) when
# # not specified.
ARG USER=0

USER ${USER}

WORKDIR /app
COPY ./opa_envoy_redis_linux_amd64 /app

ENTRYPOINT ["/app/opa_envoy_redis_linux_amd64"]

CMD ["run"]
