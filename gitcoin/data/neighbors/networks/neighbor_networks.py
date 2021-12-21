import networkx as nx
import pandas as pd
from collections import defaultdict
import matplotlib.pyplot as plt
import sys
import os
from colorhash import ColorHash

image_dest = "../images/"
adj_dest = "../adjacencies/"


def get_transaction_map(
    df: pd.DataFrame, traces: bool = True, logs: bool = False
) -> dict:
    """

    Incoming dataframe stucture:

    idx, blockNumber, transactionIndex, neighor, reason


    Output structure should be a dict with the following spec:

    transaction_index: {
        from: addr,
        to: addr,
        traces: [],
        logs: [],
        block: blockNumber,
    }
    """

    transactions = defaultdict(dict)

    # populate
    for row in df.iterrows():
        row = row[1]  # pandas is weird
        block_n = row["blockNumber"]
        txn_idx = row["transactionIndex"]
        neighbor_addr = row["neighbor"]
        reason = row["reason"]

        cur_dir = transactions[f"{block_n}, {txn_idx}"]

        # process main transaction info
        if reason in ("from", "to", "input", "creation", "miner"):
            cur_dir[reason] = neighbor_addr

        # process logs
        elif reason.startswith("log"):
            # skip if caller doesn't want logs
            if not logs:
                continue

            if "logs" not in cur_dir:
                cur_dir["logs"] = []
            cur_dir["logs"].append(neighbor_addr)

        # process traces
        elif reason.startswith("trace"):
            # skip if caller doesn't want traces
            if not traces:
                continue

            if "traces" not in cur_dir:
                cur_dir["traces"] = []
            cur_dir["traces"].append(neighbor_addr)

        else:
            raise ValueError(
                f'Reason "{reason}" was not parsable. Full row: \n\n{str(row)}\n\n'
            )

    return transactions


def get_unique_addresses(transaction_map: dict) -> set:
    """The returned set will contain all unique addresses the subject address has interacted with (including itself)."""

    addrs = set()

    for k, txn in transaction_map.items():
        if "from" in txn:
            addrs.add(txn["from"])
        if "to" in txn:
            addrs.add(txn["to"])
        if "input" in txn:
            addrs.add(txn["input"])
        if "logs" in txn:
            for elem in txn["logs"]:
                addrs.add(elem)
        if "traces" in txn:
            for elem in txn["traces"]:
                addrs.add(elem)
    return addrs


def get_transaction_graph(
    subject_address, transaction_map, unique_addresses
) -> nx.Graph:
    # use a multi directional graph to allow bi-directional/parallel edges
    G = nx.MultiDiGraph()

    color_map = []
    for addr in unique_addresses:
        G.add_node(addr)

        if addr == subject_addr:
            color_map.append("black")
        else:
            c = ColorHash(addr)
            color_map.append(c.hex)

    for _, txn in transaction_map.items():
        if "from" in txn and "to" in txn:
            G.add_edge(txn["from"], txn["to"])
        if "to" in txn and "traces" in txn:
            for elem in txn["traces"]:
                G.add_edge(txn["to"], elem)
        if "to" in txn and "logs" in txn:
            for elem in txn["logs"]:
                if elem == "0x7d655c57f71464b6f83811c55d84009cd9f5221c":
                    G.add_edge(txn["to"], elem)
                    G.add_edge(txn["from"], elem)

    print(f"nEdges:        {G.number_of_edges()}")

    graph_f = nx.draw_spring

    graph_f(
        G,
        node_color=color_map,
        node_size=20,
        edge_color="lightgrey",
        linewidths=0.01,
        arrowsize=6,
    )
    plt.savefig(image_dest + "pngs/" + subject_address + ".png")
    plt.savefig(image_dest + subject_address + ".svg", format="svg")
    nx.write_adjlist(G, adj_dest + subject_address + ".txt")


if __name__ == "__main__":
    subject_addr = sys.argv[1]

    file_size = os.path.getsize(f"../{subject_addr}.csv")
    if file_size == 0:
        print(f"File ../{subject_addr}.csv is empty.")
    else:
        df = pd.read_csv(f"../{subject_addr}.csv")
        print(df.head(5))

        txns = get_transaction_map(df)

        c = 0
        for k, v in txns.items():
            print(k, v)
            if c > 10:
                break
            c += 1

        unique_addrs = get_unique_addresses(txns)

        print(f" ")
        print(f"Address:       {subject_addr}")
        print(f"nRows:         {len(df.index)}")
        print(f"nTransactions: {len(txns)}")
        print(f"nAddresses:    {len(unique_addrs)}")

        get_transaction_graph(subject_addr, txns, unique_addrs)
