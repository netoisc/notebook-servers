FROM 3blades/pyspark-notebook:latest

MAINTAINER 3Blades <contact@g3blades.io>

USER root

# RSpark config
ENV R_LIBS_USER $SPARK_HOME/R/lib

# R pre-requisites
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    fonts-dejavu \
    gfortran \
    gcc && apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# R packages
RUN conda config --add channels r && \
    conda install --quiet --yes \
    'r-base=3.2*' \
    'r-irkernel=0.5*' \
    'r-ggplot2=1.0*' \
    'r-rcurl=1.95*' && conda clean -tipsy

ADD serve /
