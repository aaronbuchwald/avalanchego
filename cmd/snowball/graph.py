import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd


def graph_byz_round_to_termination(csv: str, n: int, k: int, save_path: str = None):
    sns.set_theme(style="ticks") # sets a global theme

    fontsize = 24

    # Initialize the figure with a linear x axis
    f, ax = plt.subplots(figsize=(10, 8))
    ax.set_xscale("linear")

    # Load the sim output data
    simdata = pd.read_csv(csv)

    # Plot the portion of byzantine weight vs. rounds to termination
    # capped at 1k rounds
    # violin plot may be good too
    sns.violinplot(
        simdata, x="byz", y="rounds", hue="byz", width=.2, palette=None, legend=False,
        # whis=[0.25, 0.75],
    ).set_title(f"n = {'{:,}'.format(n)}", fontsize=fontsize)

    # Add in points to show each observation
    sns.stripplot(simdata, x="byz", y="rounds", size=1, color=".4")

    # Tweak the visual presentation
    ax.xaxis.grid(True)
    ax.set_xlabel("Byzantine Percentage", fontsize=fontsize)
    ax.set_ylabel("Rounds to Termination", fontsize=fontsize)
    sns.despine(trim=True, left=True)

    plt.xticks(fontsize=fontsize)
    plt.yticks(fontsize=fontsize)

    if save_path is not None:
        plt.savefig(save_path)
    else:
        plt.show()

def show():
    n, k = 500, 80
    base = f'snow-sim/snow-sim-output-n-{n}-k-{k}'
    graph_byz_round_to_termination(f'{base}.csv', n, k)

def save_each():
    k = 80
    for n in [500, 1000, 10000]:
        base = f'snow-sim/snow-sim-output-n-{n}-k-{k}'
        graph_byz_round_to_termination(f'{base}.csv', n, k, f'{base}.png')

if __name__ == "__main__":
    save_each()
