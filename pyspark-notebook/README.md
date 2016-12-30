# Jupyter Notebook Python, Spark, Mesos Stack

## What it Gives You

* Jupyter Notebook 4.2.x
* Conda Python 3.x and Python 2.7.x environments
* pyspark, pandas, matplotlib, scipy, seaborn, scikit-learn pre-installed
* Spark 1.6.0 for use in local mode or to connect to a cluster of Spark workers
* Mesos client 0.22 binary that can communicate with a Mesos master
* [tini](https://github.com/krallin/tini) as the container entrypoint and [start-notebook.sh](../minimal-notebook/start-notebook.sh) as the default command

## Basic Use

The following command starts a container with the Notebook server listening for HTTP connections on port 8888 without authentication configured.

```
docker run -d -p 8888:8888 jupyter/pyspark-notebook
```

## Using Spark

> To use Python 2 in the notebook and on the workers, change the `PYSPARK_PYTHON` environment variable to point to the location of the Python 2.x interpreter binary. If you leave this environment variable unset, it defaults to `python`.

### Spark in Local Mode

This configuration is nice for using Spark on small, local data.

0. Run the container as shown above.
2. Open a Python 2 or 3 notebook.
3. Create a `SparkContext` configured for local mode.

For example, the first few cells in the notebook might read:

```python
import pyspark
sc = pyspark.SparkContext('local[*]')

# do something to prove it works
rdd = sc.parallelize(range(1000))
rdd.takeSample(False, 5)
```

### Connecting to a Spark Cluster on Mesos

This configuration allows your compute cluster to scale with your data.

0. [Deploy Spark on Mesos](http://spark.apache.org/docs/latest/running-on-mesos.html).
1. Configure each slave with [the `--no-switch_user` flag](https://open.mesosphere.com/reference/mesos-slave/) or create the `jovyan` user on every slave node.
2. Ensure Python 2.x and/or 3.x and any Python libraries you wish to use in your Spark lambda functions are installed on your Spark workers.
3. Run the Docker container with `--net=host` in a location that is network addressable by all of your Spark workers. (This is a [Spark networking requirement](http://spark.apache.org/docs/latest/cluster-overview.html#components).)
    * NOTE: When using `--net=host`, you must also use the flags `--pid=host -e TINI_SUBREAPER=true`. See https://github.com/jupyter/docker-stacks/issues/64 for details.
4. Open a Python 2 or 3 notebook.
5. Create a `SparkConf` instance in a new notebook pointing to your Mesos master node (or Zookeeper instance) and Spark binary package location.
6. Create a `SparkContext` using this configuration.

For example, the first few cells in a Python 3 notebook might read:

```python
import os
# make sure pyspark tells workers to use python3 not 2 if both are installed
os.environ['PYSPARK_PYTHON'] = '/usr/bin/python3'

import pyspark
conf = pyspark.SparkConf()

# point to mesos master or zookeeper entry (e.g., zk://10.10.10.10:2181/mesos)
conf.setMaster("mesos://10.10.10.10:5050")
# point to spark binary package in HDFS or on local filesystem on all slave
# nodes (e.g., file:///opt/spark/spark-1.6.0-bin-hadoop2.6.tgz)
conf.set("spark.executor.uri", "hdfs://10.122.193.209/spark/spark-1.6.0-bin-hadoop2.6.tgz")
# set other options as desired
conf.set("spark.executor.memory", "8g")
conf.set("spark.core.connection.ack.wait.timeout", "1200")

# create the context
sc = pyspark.SparkContext(conf=conf)

# do something to prove it works
rdd = sc.parallelize(range(100000000))
rdd.sumApprox(3)
```

### Connecting to a Spark Cluster on Standalone Mode

Connection to Spark Cluster on Standalone Mode requires the following set of steps:

0. Verify that the docker image (check the Dockerfile) and the Spark Cluster which is being deployed, run the same version of Spark.
1. [Deploy Spark on Standalone Mode](http://spark.apache.org/docs/latest/spark-standalone.html).
2. Run the Docker container with `--net=host` in a location that is network addressable by all of your Spark workers. (This is a [Spark networking requirement](http://spark.apache.org/docs/latest/cluster-overview.html#components).)
    * NOTE: When using `--net=host`, you must also use the flags `--pid=host -e TINI_SUBREAPER=true`. See https://github.com/jupyter/docker-stacks/issues/64 for details.
3. The language specific instructions are almost same as mentioned above for Mesos, only the master url would now be something like spark://10.10.10.10:7077

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