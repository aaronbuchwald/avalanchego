import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd


def graph_byz_round_to_termination(csv: str):
    sns.set_theme(style="ticks") # sets a global theme

    # Initialize the figure with a linear x axis
    f, ax = plt.subplots(figsize=(7, 6))
    ax.set_xscale("linear")

    # Load the sim output data
    simdata = pd.read_csv(csv)

    # Plot the portion of byzantine stake vs. rounds to termination
    # capped at 1k rounds
    # violin plot may be good too
    sns.violinplot(
        simdata, x="byz", y="rounds", hue="byz", width=.4, palette="vlag",
        # whis=[0, 100],
    )

    # Add in points to show each observation
    sns.stripplot(simdata, x="byz", y="rounds", size=1, color=".4")

    # Tweak the visual presentation
    ax.xaxis.grid(True)
    ax.set(ylabel="")
    sns.despine(trim=True, left=True)

    plt.show()

def run():
    graph_byz_round_to_termination('snow-sim-output-k-80.csv')

if __name__ == "__main__":
    run()
