instances = ['tai256c', 'tho150', 'wil50', 'sko100c', 'lipa80a', 'nug30', 'rou20', 'kra32', 'chr12c', 'bur26e']

def read_file(file_name):
    time = []
    cost = []

    f = open(file_name, "r")
    f1 = f.readlines()
    for x in f1:
        time.append(int(x.split()[0]))
        cost.append(int(x.split()[1]))

    return time, cost


if __name__ == "__main__":
    read_file("../results/G_chr12a.txt")
