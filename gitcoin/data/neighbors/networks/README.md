### Network Graphs

Using the `neighbors` files in the folder above this one, this Python script (thanks in large part to nollied#6773) genates network graphs.

This is very much a work in progress.

### Using

Change into this folder:

```
python -m venv env
source env/bin/activate
pip install -r requirements.txt
python neighbor_networks.py <address>
```

or

```
source run_all
```
Three files will be produced for each address:

```
../images/<address>.svg               # an SVG image of the network
../images/pngs/<address>.png          # a PNG image of the network
../adjacencies/<address>.txt          # an adjacency list of the network
```
