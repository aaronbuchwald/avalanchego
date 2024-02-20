import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd


def graph_byz_round_to_termination(csv: str, n: int, k: int, save_path: str = None):
    sns.set_theme(style="ticks") # sets a global theme

    # Initialize the figure with a linear x axis
    f, ax = plt.subplots(figsize=(10, 8))
    ax.set_xscale("linear")

    # Load the sim output data
    simdata = pd.read_csv(csv)

    # Plot the portion of byzantine stake vs. rounds to termination
    # capped at 1k rounds
    # violin plot may be good too
    sns.violinplot(
        simdata, x="byz", y="rounds", hue="byz", width=.4, palette="vlag",
        # whis=[0, 100],
    ).set_title('Byzantine Stake vs. Rounds to Termination')

    # Add in points to show each observation
    sns.stripplot(simdata, x="byz", y="rounds", size=1, color=".4")

    # Tweak the visual presentation
    ax.xaxis.grid(True)
    ax.set(ylabel="")
    sns.despine(trim=True, left=True)

    plt.figtext(0.9, 0.05, f'n = {n}\nk = {k}', ha='center', va='center', fontsize=10, bbox=dict(facecolor='white', alpha=0.5, edgecolor='black'))


    if save_path is not None:
        plt.savefig(save_path)

    plt.show()

def run():
    n, k = 10000, 80
    base = f'snow-sim-output-n-{n}-k-{k}'
    graph_byz_round_to_termination(f'{base}.csv', n, k, f'{base}.png')

if __name__ == "__main__":
    run()
