import pandas as pd
import numpy as np
from matplotlib import pyplot as plt, ticker

from lib.lib_agent import agent_name, kubo_version
from lib.lib_fmt import fmt_thousands
from lib.lib_plot import cdf


def configure_axis(ax):
    ax.set_xlim(0, 48)
    ax.set_xticks(np.arange(0, 50, step=2))
    ax.set_xlabel("Inter-Arrival Time in Hours")


def plot_inter_arrival_time(df: pd.DataFrame) -> plt.Figure:
    df = df.assign(inter_arrival_time_in_h=df.inter_arrival_time.apply(lambda x: x.total_seconds() / 3600))

    fig, (ax1, ax2, ax3) = plt.subplots(1, 3, figsize=(15, 5), sharey=True)

    # First plot
    data = cdf(df["inter_arrival_time_in_h"])
    ax1.plot(data["inter_arrival_time_in_h"], data["cdf"])
    ax1.legend(loc='lower right', labels=[f"all ({fmt_thousands(len(df))})"])
    ax1.set_title("Overall")
    ax1.set_ylabel("Online Peers in %")
    ax2.get_yaxis().set_major_formatter(ticker.FuncFormatter(lambda x, p: "%d" % int(x * 100)))
    configure_axis(ax1)

    df = df.dropna().assign(
        agent_name=lambda data_frame: data_frame.agent_version.apply(agent_name),
        kubo_version=lambda data_frame: data_frame.agent_version.apply(kubo_version),
    )

    # Second plot
    for agent in list(df["agent_name"].unique()) + ['other']:
        data = df[df['agent_name'] == agent]
        data = cdf(data["inter_arrival_time_in_h"])
        ax2.plot(data["inter_arrival_time_in_h"], data["cdf"], label=f"{agent} ({fmt_thousands(len(data))})")

    ax2.legend(loc='lower right')
    ax2.set_title("By Agent")
    configure_axis(ax2)

    # Third plot
    df = df.dropna().assign(
        minor=lambda data_frame: data_frame.kubo_version.apply(lambda row: int(row.split(".")[1])),
    )

    filter_group = df \
        .groupby(by="minor", as_index=False) \
        .count() \
        .sort_values('agent_name', ascending=False)
    filter_group = filter_group.head(10)

    for minor in reversed(sorted(df["minor"].unique())):
        if minor not in set(filter_group["minor"]):
            continue

        data = df[df['minor'] == minor]
        data = cdf(data["inter_arrival_time_in_h"])
        ax3.plot(data["inter_arrival_time_in_h"], data["cdf"], label=f"0.{minor}.x ({fmt_thousands(len(data))})")

    ax3.legend(loc='lower right')
    ax3.set_title("By Kubo Version")
    configure_axis(ax3)

    fig.set_tight_layout(True)

    return fig