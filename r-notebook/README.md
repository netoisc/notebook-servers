# Jupyter Notebook R Stack

## What it Gives You

* Jupyter Notebook 4.2.x
* Conda R v3.2.x and channel
* plyr, devtools, dplyr, ggplot2, tidyr, shiny, rmarkdown, forecast, stringr, rsqlite, reshape2, nycflights13, caret, rcurl, and randomforest pre-installed
* [tini](https://github.com/krallin/tini) as the container entrypoint and [start-notebook.sh](../minimal-notebook/start-notebook.sh) as the default command

## Basic Use

The following command starts a container with the Notebook server listening for HTTP connections on port 8888 without authentication configured.

```
docker run -d -p 8888:8888 3blades/r-notebook
```