FROM 3blades/minimal-notebook

MAINTAINER 3Blades <contact@3blades.io>

USER root

# libav-tools for matplotlib anim
RUN apt-get update && \
    apt-get install -y --no-install-recommends libav-tools && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install Python 3 packages
RUN conda install --quiet --yes \
    'conda-build' \
    'ipywidgets=5.1*' \
    'pandas=0.17*' \
    'numexpr=2.5*' \
    'matplotlib=1.5*' \
    'scipy=0.17*' \
    'seaborn=0.7*' \
    'scikit-learn=0.17*' \
    'scikit-image=0.11*' \
    'sympy=1.0*' \
    'cython=0.23*' \
    'patsy=0.4*' \
    'statsmodels=0.6*' \
    'cloudpickle=0.1*' \
    'dill=0.2*' \
    'numba=0.23*' \
    'bokeh=0.11*' \
    'h5py=2.5*' \
    'mpld3=0.2*' \
    && conda clean -tipsy
# Activate ipywidgets extension in the environment that runs the notebook server
RUN jupyter nbextension enable --py widgetsnbextension --sys-prefix

# Install Python 2 packages
RUN conda create --quiet --yes -p $CONDA_DIR/envs/python2 python=2.7 \
    'ipython=4.*' \
    'ipywidgets=5.1*' \
    'pandas=0.17*' \
    'numexpr=2.5*' \
    'matplotlib=1.5*' \
    'scipy=0.17*' \
    'seaborn=0.7*' \
    'scikit-learn=0.17*' \
    'scikit-image=0.11*' \
    'sympy=1.0*' \
    'cython=0.23*' \
    'patsy=0.4*' \
    'statsmodels=0.6*' \
    'cloudpickle=0.1*' \
    'dill=0.2*' \
    'numba=0.23*' \
    'bokeh=0.11*' \
    'h5py=2.5*' \
    'pyzmq' \
    'mpld3=0.2*' \
    && conda clean -tipsy

# Add shortcuts to distinguish pip for python2. python3 is enabled
# as default with pip already
RUN ln -s $CONDA_DIR/envs/python2/bin/pip $CONDA_DIR/bin/pip2

# Install Python 2 kernel spec globally
RUN $CONDA_DIR/envs/python2/bin/python -m ipykernel install

ADD serve /
