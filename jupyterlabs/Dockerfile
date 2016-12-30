FROM ubuntu:16.04

MAINTAINER 3Blades <contact@3blades.io>

USER root

RUN apt-get update && apt-get install -yq --no-install-recommends \
    bzip2 \
    ca-certificates \
    python-dev \
    python3-dev \
    python-pip \
    python3 \
    python3-pip \
    libpq-dev \
    libmysqlclient-dev \
    wget \
    && apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN echo "en_US.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen

# Install Tini
RUN wget --quiet https://github.com/krallin/tini/releases/download/v0.10.0/tini && \
    echo "1361527f39190a7338a0b434bd8c88ff7233ce7b9a4876f3315c22fce7eca1b0 *tini" | sha256sum -c - && \
    mv tini /usr/local/bin/tini && \
    chmod +x /usr/local/bin/tini

# Install jupyter notebook and jupyterlabs
RUN pip3 install -U setuptools pip wheel
RUN pip3 install jupyter
RUN pip3 install jupyterlab
RUN jupyter serverextension enable --py jupyterlab --sys-prefix

RUN pip3 install \
    psycopg2 \
    pymongo \
    cassandra-driver \
    mysqlclient

# Add a notebook profile.
RUN mkdir -p -m 700 /root/.jupyter/
COPY jupyter_notebook_config.py /root/.jupyter/jupyter_notebook_config.py

RUN mkdir /home/work

# Configure container startup as root
EXPOSE 8888
ENTRYPOINT ["tini", "--"]
CMD ["start-notebook.sh"]

# Add local files as late as possible to avoid cache busting
COPY start-notebook.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/start-notebook.sh

ADD serve /
