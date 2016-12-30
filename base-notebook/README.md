# Base Jupyter Notebook Stack

Small base image for defining your own stack

## What it Gives You

* Minimally-functional Jupyter Notebook 4.2.x (e.g., no pandoc for document conversion)
* Miniconda Python 3.x
* No preinstalled scientific computing packages
* [tini](https://github.com/krallin/tini) as the container entrypoint and [start-notebook.sh](./start-notebook.sh) as the default command

## Basic Use

The following command starts a container with the Notebook server listening for HTTP connections on port 8888 without authentication configured.

```
docker run -d -p 8888:8888 jupyter/base-notebook
```

## Conda Environment

The default Python 3.x [Conda environment](http://conda.pydata.org/docs/using/envs.html) resides in `/opt/conda`. The commands `ipython`, `python`, `pip`, `easy_install`, and `conda` (among others) are available in this environment.