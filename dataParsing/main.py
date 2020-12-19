import matplotlib.pyplot as plt
import numpy as np

instances = ['chr12c', 'nug30']
algorithms = ['S', 'G', 'H', 'R', 'RW', 'TS']


def read_file(file_name):
    time = []
    cost = []
    steps = []
    explored_solutions = []

    f = open(file_name, "r")
    f1 = f.readlines()
    for x in f1:
        cost.append(int(x.split()[0]))
        steps.append(int(x.split()[1]))
        explored_solutions.append(int(x.split()[2]))
        time.append(int(x.split()[3]))

    return cost, steps, explored_solutions, time


def plot_score(data_mean, data_std):
    width = 0.35

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(len(algorithms)):
        ax.bar(x + i * (width / len(algorithms)), data_mean[i], width / len(algorithms), yerr=data_std[i],
               label=algorithms[i])

    ax.set_xticks(x + (len(algorithms) // 2) * width / len(algorithms))
    ax.set_xticklabels(instances)

    ax.set_ylabel('Koszt')
    ax.set_title('Koszty dla każdego z algorytmów')
    ax.legend()

    plt.show()


def plot_time(data_mean, data_std):
    width = 0.35

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(3):
        ax.bar(x + i * (width / 3), data_mean[i], width / 3, yerr=data_std[i], label=algorithms[i])

    ax.set_xticks(x + (3 // 2) * width / 3)
    ax.set_xticklabels(instances)

    ax.set_ylabel('Średni czas w mikrosekundach')
    ax.set_title('Czasy dla algorytmów S, G oraz H')
    ax.legend()

    plt.show()


def plot_steps(data_mean, data_std):
    width = 0.35

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(3):
        ax.bar(x + i * (width / 3), data_mean[i], width / 3, yerr=data_std[i], label=algorithms[i])

    ax.set_xticks(x + (3 // 2) * width / 3)
    ax.set_xticklabels(instances)

    ax.set_ylabel('Średnia liczba kroków')
    ax.set_title('Liczba kroków dla algorytmów S, G oraz H')
    ax.legend()

    plt.show()


def plot_explored_solutions(data_mean, data_std):
    width = 0.35

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(3):
        ax.bar(x + i * (width / 3), data_mean[i], width / 3, yerr=data_std[i], label=algorithms[i])

    ax.set_xticks(x + (3 // 2) * width / 3)
    ax.set_xticklabels(instances)

    ax.set_ylabel('Średnia sprawdzonych rozwiązań')
    ax.set_title('Liczba sprawdzonych rozwiązań dla algorytmów S, G oraz H')
    ax.legend()

    plt.show()


if __name__ == "__main__":
    alg_mean_cost = []
    alg_min_cost = []
    alg_std_cost = []

    alg_mean_time = []
    alg_std_time = []

    alg_mean_steps = []
    alg_std_steps = []

    alg_mean_explored_solutions = []
    alg_std_explored_solutions = []

    for name in algorithms:
        mean_cost = []
        min_cost = []
        std_cost = []

        mean_time = []
        std_time = []

        mean_steps = []
        std_steps = []

        mean_explored_solutions = []
        std_explored_solutions = []

        for instance in instances:
            cost, steps, explored_solutions, time = read_file("../results/" + name + "_" + instance + ".txt")
            mean_cost.append(np.mean(cost))
            min_cost.append(np.min(cost))
            std_cost.append(np.std(cost))

            if name != 'R' and name != 'RW':
                mean_steps.append(np.mean(steps))
                std_steps.append(np.std(steps))

            mean_explored_solutions.append(np.mean(explored_solutions))
            std_explored_solutions.append(np.std(explored_solutions))

            mean_time.append(np.mean(time))
            std_time.append(np.std(time))

        alg_mean_cost.append(mean_cost)
        alg_min_cost.append(min_cost)
        alg_std_cost.append(std_cost)

        alg_mean_time.append(mean_time)
        alg_std_time.append(std_time)

        alg_mean_steps.append(mean_steps)
        alg_std_steps.append(std_steps)

        alg_mean_explored_solutions.append(mean_explored_solutions)
        alg_std_explored_solutions.append(std_explored_solutions)

    plot_score(alg_mean_cost, alg_std_cost)
    plot_time(alg_mean_time, alg_std_time)
    plot_steps(alg_mean_steps, alg_std_steps)
    plot_explored_solutions(alg_mean_explored_solutions, alg_std_explored_solutions)
