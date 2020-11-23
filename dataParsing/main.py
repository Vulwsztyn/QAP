import matplotlib.pyplot as plt
import numpy as np

instances = ["bur26d", "kra30a", "tho40", "wil50", "lipa60a", "lipa70a", "tai80a", "sko90", "sko100a", "esc128"]
algorithms = ['S', 'G', 'H', 'R', 'RW']


def read_instance(name, instance, size):
    time = []
    cost = []
    steps = []
    explored_solutions = []
    init_result = []
    dist = []

    f = open("../results/" + name + "_" + instance + "_" + size + ".txt", "r")
    f1 = f.readlines()
    for x in f1:
        cost.append(int(x.split()[0]))
        steps.append(int(x.split()[1]))
        explored_solutions.append(int(x.split()[2]))
        time.append(int(x.split()[3]))
        init_result.append(int(x.split()[4]))
        dist.append(float(x.split()[5]))

    return cost, steps, explored_solutions, time, init_result, dist

def read_solution(instance):
    solution = []
    best_score = 0

    f = open("../solutions/" + instance + ".sln", "r")
    f1 = f.readlines()
    i=0
    for x in f1:
        if i==0:
            best_score = int(x.split()[1])
        else:
            solution += [int(x)-1 for x in x.split()]
        i+=1

    return solution, best_score


def plot_quality(data_mean, data_std):
    width = 0.6

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(len(algorithms)):
        # ax.plot(x + i * (width / len(algorithms)), data_mean[i], '--o', label=algorithms[i])

        ax.bar(x + i * (width / len(algorithms)), data_mean[i], width / len(algorithms), yerr=data_std[i],
               label=algorithms[i])

        # plt.plot(x + i * (width / len(algorithms)), data_mean[i], '_')

    ax.set_xticks(x + (len(algorithms) // 2) * width / len(algorithms))
    ax.set_xticklabels(instances)

    ax.set_ylabel('Jakość')
    plt.xticks(rotation=-45)

    ax.legend()


    #plt.show()
    plt.savefig('2.1.1.png')

def plot_quality_min(data_mean):
    width = 0.6

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(len(algorithms)):
        ax.bar(x + i * (width / len(algorithms)), data_mean[i], width / len(algorithms),
               label=algorithms[i])

        # plt.plot(x + i * (width / len(algorithms)), data_mean[i], '_')

    ax.set_xticks(x + (len(algorithms) // 2) * width / len(algorithms))
    ax.set_xticklabels(instances)

    ax.set_ylabel('Jakość')
    ax.legend()

    plt.xticks(rotation=-45)

    #plt.show()
    plt.savefig('2.1.2.png')


def plot_time(data_mean, data_std):
    width = 0.6

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(3):
        ax.bar(x + i * (width / 3), data_mean[i], width / 3, yerr=data_std[i], label=algorithms[i])

    ax.set_xticks(x + (3 // 2) * width / 3)
    ax.set_xticklabels(instances)

    ax.set_ylabel('Średni czas w mikrosekundach')
    ax.legend()
    plt.xticks(rotation=-45)


    #plt.show()
    plt.savefig('2.2.png')


def plot_effectiveness(data_mean, data_std):
    width = 0.6

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(len(algorithms)):
        ax.bar(x + i * (width / len(algorithms)), data_mean[i], width / len(algorithms),  yerr=data_std[i],
               label=algorithms[i])

        # plt.plot(x + i * (width / len(algorithms)), data_mean[i], '_')

    ax.set_xticks(x + (len(algorithms) // 2) * width / len(algorithms))
    ax.set_xticklabels(instances)

    ax.set_ylabel('Efektywność')
    ax.legend()
    plt.xticks(rotation=-45)

    ax.set_yscale('log')

    #plt.show()
    plt.savefig('2.3.png')


def plot_steps(data_mean, data_std):
    width = 0.6

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(2):
        ax.bar(x + i * (width / 2), data_mean[i], width / 2, yerr=data_std[i], label=algorithms[i])

    ax.set_xticks(x +  width / 2)
    ax.set_xticklabels(instances)

    ax.set_ylabel('Średnia liczba kroków')
    ax.legend()
    plt.xticks(rotation=-45)

    #plt.show()
    plt.savefig('2.4.png')


def plot_explored_solutions(data_mean, data_std):
    width = 0.6

    fig, ax = plt.subplots()
    x = np.arange(len(instances))

    for i in range(4):
        ax.bar(x + i * (width / 4), data_mean[i], width / 4, yerr=data_std[i], label=['S', 'G', 'R', 'RW'][i])

    ax.set_xticks(x + (4 // 2) * width / 4)
    ax.set_xticklabels(instances)

    ax.set_ylabel('Średnia sprawdzonych rozwiązań')
    ax.legend()
    plt.xticks(rotation=-45)

    #plt.show()
    plt.savefig('2.5.png')

def plot_init_result(alg_init_costs, alg_costs, label):

    fig, ax = plt.subplots()

    ax.scatter(alg_costs[0], alg_init_costs[0], label='S')
    ax.scatter(alg_costs[1], alg_init_costs[1], label='G')


    ax.set_ylabel('Końcowa jakość')
    ax.set_xlabel('Początkowa jakość')

    ax.legend()

    plt.savefig('3.'+label+'.png')

def plot_costs(costs, label):

    fig, ax = plt.subplots()

    x = np.arange(len(costs[0]))


    n_of_starts_G = []
    n_of_starts_S = []
    for i in range(1,len(costs[0])):
        n_of_starts_G.append(np.mean(costs[0][:i]))
        n_of_starts_S.append(np.mean(costs[1][:i]))


    ax.plot(x[1:], n_of_starts_G,label='S mean', color='blue')
    ax.plot(x[1:], n_of_starts_S,label='G mean', color='orange')

    n_of_starts_G = []
    n_of_starts_S = []
    for i in range(1,len(costs[0])):
        n_of_starts_G.append(np.max(costs[0][:i]))
        n_of_starts_S.append(np.max(costs[1][:i]))


    ax.plot(x[1:], n_of_starts_G, '--', label='S best', color='blue')
    ax.plot(x[1:], n_of_starts_S, '--', label='G best', color='orange')


    ax.set_ylabel('Jakość')
    ax.set_xlabel('Liczba startów')
    ax.legend()

    #plt.show()
    plt.savefig('4.'+label+'.png')

def plot_dists_optimal(qualities, dists, label):
    fig, ax = plt.subplots()

    for i in range(len(qualities)):
        ax.scatter(qualities[i], dists[i], label=algorithms[i])


    ax.set_ylabel('Dystans')
    ax.set_xlabel('Jakość')

    ax.legend()

    plt.savefig('5.'+label+'.png')

if __name__ == "__main__":
    alg_mean_quality = []
    alg_min_quality = []
    alg_std_quality = []

    alg_mean_time = []
    alg_std_time = []

    alg_mean_effectiveness = []
    alg_std_effectiveness = []

    alg_mean_steps = []
    alg_std_steps = []

    alg_mean_explored_solutions = []
    alg_std_explored_solutions = []

    best_solutions = []

    for instance in instances:
        _, score = read_solution(instance)
        best_solutions.append(score)

    for name in algorithms:
        mean_quality = []
        min_quality = []
        std_quality = []

        mean_time = []
        std_time = []

        mean_effectiveness = []
        std_effectiveness = []

        mean_steps = []
        std_steps = []

        mean_explored_solutions = []
        std_explored_solutions = []

        i = 0
        for instance in instances:
            cost, steps, explored_solutions, time, init_cost, _ = read_instance(name, instance, "10")
            quality = [best_solutions[i]/x for x in cost]
            init_quality = [best_solutions[i]/x for x in init_cost]
            effectiveness = np.true_divide(quality,time)
            mean_quality.append(np.mean(quality))
            min_quality.append(np.min(quality))
            std_quality.append(np.std(quality))

            if name != 'R' and name != 'RW' and name != 'H':
                mean_steps.append(np.mean(steps))
                std_steps.append(np.std(steps))

            if name != 'H':
                mean_explored_solutions.append(np.mean(explored_solutions))
                std_explored_solutions.append(np.std(explored_solutions))


            mean_time.append(np.mean(time))
            std_time.append(np.std(time))

            mean_effectiveness.append(np.mean(effectiveness))
            std_effectiveness.append(np.std(effectiveness))

            i += 1

        alg_mean_quality.append(mean_quality)
        alg_min_quality.append(min_quality)
        alg_std_quality.append(std_quality)

        alg_mean_time.append(mean_time)
        alg_std_time.append(std_time)

        alg_mean_effectiveness.append(mean_effectiveness)
        alg_std_effectiveness.append(std_effectiveness)

        if name != 'R' and name != 'RW' and name != 'H':
            alg_mean_steps.append(mean_steps)
            alg_std_steps.append(std_steps)

        if name != 'H':
            alg_mean_explored_solutions.append(mean_explored_solutions)
            alg_std_explored_solutions.append(std_explored_solutions)


    plot_quality(alg_mean_quality, alg_std_quality)
    plot_quality_min(alg_min_quality)
    plot_time(alg_mean_time, alg_std_time)
    plot_effectiveness(alg_mean_effectiveness, alg_std_effectiveness)
    plot_steps(alg_mean_steps, alg_std_steps)
    plot_explored_solutions(alg_mean_explored_solutions, alg_std_explored_solutions)

    i=0
    for instance in instances[:5]:
        i+=1
        S_cost, _, _, _, S_init_cost, _ = read_instance(algorithms[0], instance, "300")
        G_cost, _, _, _, G_init_cost, _ = read_instance(algorithms[1], instance, "300")
        _, score = read_solution(instance)
        S_quality = [score / x for x in S_cost]
        S_init_quality = [score / x for x in S_init_cost]
        G_quality = [score / x for x in G_cost]
        G_init_quality = [score / x for x in G_init_cost]

        plot_init_result([S_quality,G_quality], [S_init_quality,G_init_quality], str(i))
        plot_costs([S_quality, G_quality], str(i))

    i = 0
    for instance in instances[:5]:
        i += 1
        _, score = read_solution(instance)
        qualities = []
        dists = []
        for alg in algorithms:
            cost, _, _, _, _, dist = read_instance(alg, instance, "300")
            qualities.append([score / x for x in cost])
            dists.append(dist)
        plot_dists_optimal(qualities, dists, str(i))


    plt.show()
