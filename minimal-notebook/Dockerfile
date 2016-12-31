FROM 3blades/base-notebook

MAINTAINER 3Blades <contact@3blades.io>

USER root

# Install all OS dependencies for fully functional notebook server
RUN apt-get update && apt-get install -yq --no-install-recommends \
    git \
    vim \
    jed \
    emacs \
    build-essential \
    python-dev \
    python-pip \
    unzip \
    libsm6 \
    pandoc \
    libpq-dev \
    libmysqlclient-dev \
    texlive-latex-base \
    texlive-latex-extra \
    texlive-fonts-extra \
    texlive-fonts-recommended \
    texlive-generic-recommended \
    libxrender1 \
    && apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN pip install --upgrade pip
# https://github.com/ContinuumIO/anaconda-issues/issues/542
RUN pip install -q --upgrade --ignore-installed pip wheel setuptools
RUN pip install \
    psycopg2 \
    pymongo \
    cassandra-driver \
    mysqlclient

ADD serve /