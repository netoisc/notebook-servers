# Jupyter Notebook Scientific Python Stack

## What it Gives You

* Jupyter Notebook 4.2.x
* Conda Python 3.x and Python 2.7.x environments
* pandas, matplotlib, scipy, seaborn, scikit-learn, scikit-image, sympy, cython, patsy, statsmodel, cloudpickle, dill, numba, bokeh pre-installed
* [tini](https://github.com/krallin/tini) as the container entrypoint and [start-notebook.sh](../minimal-notebook/start-notebook.sh) as the default command

## Basic Use

The following command starts a container with the Notebook server listening for HTTP connections on port 8888 without authentication configured.

```
docker run -d -p 8888:8888 jupyter/scipy-notebook
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

The commands `jupyter`, `ipython`, `python`, `pip`, `easy_install`, and `conda` (among others) are available in both environments. For convenience, you can install packages into either environment regardless of what environment is currently active using commands like the following:

```
# install a package into the python2 environment
pip2 install some-package
conda install -n python2 some-package

# install a package into the default (python 3.x) environment
pip3 install some-package
conda install -n python3 some-package
```