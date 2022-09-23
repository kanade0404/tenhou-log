FROM hashicorp/terraform:1.3.0 AS terraform

FROM ubuntu:20.04 AS base
RUN apt update && \
    apt install -y unzip curl wget && \
    curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash && \
    curl -Lso tfsec https://github.com/tfsec/tfsec/releases/download/v1.15.4/tfsec-linux-amd64 && \
    chmod +x tfsec && \
    mv tfsec /usr/local/bin/tfsec && \
    wget https://github.com/minamijoyo/tfmigrate/releases/download/v0.3.7/tfmigrate_0.3.7_linux_amd64.tar.gz && \
    tar -zxvf tfmigrate_0.3.7_linux_amd64.tar.gz tfmigrate && \
    mv tfmigrate /usr/local/bin/tfmigrate

FROM ubuntu:20.04 AS gcloud
RUN apt update && \
    apt install -y curl apt-transport-https ca-certificates gnupg python && \
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - && \
    apt update && \
    apt install google-cloud-sdk

COPY --from=terraform /bin/terraform /bin/terraform
COPY --from=base /usr/local/bin/tflint /usr/local/bin/tflint
COPY --from=base /usr/local/bin/tfsec /usr/local/bin/tfsec
COPY --from=base /usr/local/bin/tfmigrate /usr/local/bin/tfmigrate
CMD ["bash"]
