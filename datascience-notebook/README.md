# Jupyter Notebook Data Science Stack

## What it Gives You

* Jupyter Notebook 4.1.x
* Conda Python 3.x and Python 2.7.x environments
* pandas, matplotlib, scipy, seaborn, scikit-learn, scikit-image, sympy, cython, patsy, statsmodel, cloudpickle, dill, numba, bokeh pre-installed
* Conda R v3.2.x and channel
* plyr, devtools, dplyr, ggplot2, tidyr, shiny, rmarkdown, forecast, stringr, rsqlite, reshape2, nycflights13, caret, rcurl, and randomforest pre-installed
* Julia v0.3.x with Gadfly and RDatasets pre-installed
* Options for HTTPS, password auth, and passwordless `sudo`

## Basic Use

The following command starts a container with the Notebook server listening for HTTP connections on port 8888 without authentication configured.

```
docker run -d -p 8888:8888 jupyter/datascience-notebook
```

## Conda Environments

The default Python 3.x [Conda environment](http://conda.pydata.org/docs/using/envs.html) resides in `/opt/conda`. A second Python 2.x Conda environment exists in `/opt/conda/envs/python2`. You can [switch to the python2 environment](http://conda.pydata.org/docs/using/envs.html#change-environments-activate-deactivate) in a shell by entering the following:

```
source activate python2
```

You can return to the default environment with this command:

```
source deactivate
```

The commands `ipython`, `python`, `pip`, `easy_install`, and `conda` (among others) are available in both environments.