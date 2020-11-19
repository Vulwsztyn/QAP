import matplotlib
import matplotlib.pyplot as plt
import numpy as np

instances = ['tai256c', 'tho150', 'wil50', 'sko100c', 'lipa80a', 'nug30', 'rou20', 'kra32', 'chr12c', 'bur26e']
algorithms = ['S', 'G', 'R', 'RW']

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

    return cost,steps,explored_solutions, time


def plot_score(labels, data_mean, data_min, data_std, filename):
    width = 0.35  # the width of the bars: can also be len(x) sequence

    fig, ax = plt.subplots()

    ax.bar(labels, data_min, width, label='Najlepszy wynik')
    ax.bar(labels, np.subtract(data_mean, data_min), width/2, yerr=data_std, bottom=data_min,
           label='Średni wynik')

    ax.set_ylabel('Koszt')
    ax.set_title('Koszty dla instancji ' +filename+ ' dla każdego z algorytmów')
    ax.legend()

    plt.show()

def plot_time(labels, data_mean, data_std, filename):
    width = 0.35  # the width of the bars: can also be len(x) sequence

    fig, ax = plt.subplots()

    ax.bar(labels, data_mean, width,yerr=data_std)

    ax.set_ylabel('Średni czas w mikrosekundach')
    ax.set_title('Czasy dla instancji ' +filename+ ' dla każdego z algorytmów')

    plt.show()

def plot_steps(labels, data_mean, data_std, filename):
    width = 0.35  # the width of the bars: can also be len(x) sequence

    fig, ax = plt.subplots()

    ax.bar(labels, data_mean, width,yerr=data_std)

    ax.set_ylabel('Średnia liczba kroków')
    ax.set_title('Liczba kroków dla instancji ' +filename+ ' dla S oraz G')

    plt.show()

def plot_explored_solutions(labels, data_mean, data_std, filename):
    width = 0.35  # the width of the bars: can also be len(x) sequence

    fig, ax = plt.subplots()

    ax.bar(labels, data_mean, width,yerr=data_std)

    ax.set_ylabel('Średnia sprawdzonych rozwiązań')
    ax.set_title('Liczba sprawdzonych rozwiązań dla instancji ' +filename+ ' dla S oraz G')

    plt.show()

if __name__ == "__main__":
    mean_cost = []
    min_cost = []
    std_cost = []

    mean_time = []
    std_time = []

    mean_steps = []
    std_steps = []

    mean_explored_solutions = []
    std_explored_solutions = []

    for name in algorithms:
        cost,steps,explored_solutions, time = read_file("../results/"+name+"_chr12c.txt")
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

    plot_score(algorithms,mean_cost, min_cost, std_cost, 'chr12c' )
    plot_time(algorithms, mean_time, std_time, 'chr12c')
    plot_steps(['S', 'G'], mean_steps, std_steps, 'chr12c')
    plot_explored_solutions(algorithms, mean_explored_solutions, std_explored_solutions, 'chr12c')
