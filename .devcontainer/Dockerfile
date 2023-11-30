FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye


#installing nats


# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install github.com/nats-io/natscli/nats@latest
USER root

# now setting inner nats server as default in container
# so now can 
#     nat pub com.my.topic
#     nat sub com.my.*
RUN nats context add nats --server nats --description "NATS Service In This Container" --select



# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
