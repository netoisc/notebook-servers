# Minimal Jupyter Notebook Stack

Small image for working in the notebook and installing your own libraries

## What it Gives You

* Fully-functional Jupyter Notebook 4.2.x
* Miniconda Python 3.x
* No preinstalled scientific computing packages
* [tini](https://github.com/krallin/tini) as the container entrypoint and [start-notebook.sh](./start-notebook.sh) as the default command

## Basic Use

The following command starts a container with the Notebook server listening for HTTP connections on port 8888 without authentication configured.

```
docker run -d -p 8888:8888 jupyter/minimal-notebook
```

## Conda Environment

The default Python 3.x [Conda environment](http://conda.pydata.org/docs/using/envs.html) resides in `/opt/conda`. The commands `ipython`, `python`, `pip`, `easy_install`, and `conda` (among others) are available in this environment.