### Network Graphs

Using the `neighbors` files in the folder above this one, this Python script (thanks in large part to nollied#6773) genates network graphs, images of the same, and adjacency lists.

This data is very much a work in progress.

### Notes

This data is alpha status and may be (and probably is) incomplete.

### How this Data is Created

1. Change into this folder
2. Run the following code:

```
python -m venv env
source env/bin/activate
pip install -r requirements.txt
python neighbor_networks.py <address>
```

or, to generate the data for all the addresses, run this:

```
source run_all
```
Three files will be produced for each address:

```
../images/<address>.svg               # an SVG image of the network
../images/pngs/<address>.png          # a PNG image of the network
../adjacencies/<address>.txt          # an adjacency list of the network
```
