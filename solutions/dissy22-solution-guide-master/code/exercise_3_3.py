n = 15

edges = [(1, 3),(1, 12),(2, 1),(2, 14),(3, 2),(3, 15),
         (4, 2),(4, 13),(5, 12),(5, 14),(6, 5),(6, 11),
         (7, 6),(7, 11),(8, 6),(8, 14),(9, 10),(9, 15),
         (10, 4),(10, 5),(11, 10),(11, 13),(12, 4),(12, 7),
         (13, 9),(13, 14),(14, 1),(14, 11),(15, 7),(15, 8)]

def all_pairs_shortest_path(n, edges):
    cost = [[-1 for _ in range(n)] for _ in range(n)]
    for s, t in edges:
        cost[s-1][t-1] = 1

    for i in range(n):
        cost[i][i] = 0

    for k in range(n):
        for i in range(n):
            for j in range(n):
                if not cost[i][k] == -1 and \
                   not cost[k][j] == -1 and \
                   (cost[i][k] + cost[k][j] < cost[i][j] or cost[i][j] == -1):
                    cost[i][j] = cost[i][k] + cost[k][j]

    return cost


from collections import deque

def relabel_to_front(n, s, t, edge):
    # max-flow min-cut / ford-fulkerson , edmund-karp, Relabel-to-front O(V^3)

    # n,m,s,t = map(int, input().split())

    # initialize-preflow
    adjecency = [[] for _ in range(n)]
    h = [0 for _ in range(n)]
    e = [0 for _ in range(n)]
    c = [{} for _ in range(n)]
    f = [{} for _ in range(n)]
    current = [0 for _ in range(n)]

    # O(m)
    for u, v in edge:
        w = 1

        adjecency[u].append(v)
        adjecency[v].append(u)

        c[u][v] = w
        c[v][u] = 0

        f[u][v] = 0
        f[v][u] = 0

    # O(1)
    h[s] = n

    # O(n)
    for v in adjecency[s]:
        csv = c[s][v]

        f[s][v] += csv
        f[v][s] -= csv

        e[v] += csv
        e[s] -= csv

    # O(1)
    def cf(e):
        (u,v) = e
        return c[u][v] - f[u][v]

    def bfs_height(u):
        S = []
        discovered = [False for _ in range(n)]

        uindex = 0

        S.append(u)
        while not len(S) == 0:
            u = S.pop()

            if discovered[u]:
                continue

            h[u] = uindex
            uindex += 1

            discovered[u] = True

            for v in adjecency[u]:
                if cf((v,u)) > 0:
                    S.append(v)

    # O(1)
    def push(uv):
        (u,v) = uv

        df_u_v = min(e[u], cf((u,v)))

        f[u][v] += df_u_v
        f[v][u] -= df_u_v

        e[u] -= df_u_v
        e[v] += df_u_v


    # O(n)
    def relabel(u):
        h[u] = 1 + min(list(map(lambda x: h[x], filter(lambda v: cf((u,v)) > 0, adjecency[u]))))

    # O(n^2)
    def discharge(u):
        while e[u] > 0:
            if len(adjecency[u]) == current[u]:
                relabel(u)
                current[u] = 0
            else:
                v = adjecency[u][current[u]]

                if cf((u,v)) > 0 and h[u] == h[v] + 1:
                    push((u,v))
                else:
                    current[u] += 1

    def relabel_to_front_algo():
        # O(n)
        Lback = deque([])
        for i in range(n):
            if i == s or i == t:
                continue
            Lback.append(i)
        Lfront = deque([])

        iters = 0

        # O(n^3)
        while not len(Lback) == 0:
            # Exact node labels heuristic
            iters += 1
            if iters % 25000 == 0:
                bfs_height(t)

            u = Lback.popleft()

            old_height = h[u]
            discharge(u)

            if h[u] > old_height:
                Lback = Lfront + Lback
                Lfront = deque([])
            Lfront.append(u)


    relabel_to_front_algo()

    res = []
    for i, fi in enumerate(f):
        for j in fi.keys():
            if fi[j] > 0:
                res.append((i,j,fi[j]))

    return (e[t])

    # print (n, e[t], len(res))
    # for (i,j,v) in res:
    #     print (i,j,v)

edges_ = [(a-1, b-1) for (a,b) in edges]

min_val = -1
min_disjoint_connections = []
for i in range(n):
    for j in range(n):
        if not i is j:
            val = relabel_to_front(n, i, j, edges_)
            if min_val == -1 or val < min_val:
                min_disjoint_connections = [(i+1, j+1)]
                min_val = val
            elif val == min_val:
                min_disjoint_connections.append((i+1, j+1))
print ("Connectivity:", min_val)

print ("   ", end="")
for i in range(n):
    print(str(i+1).rjust(3), end="")
print()
for i in range(n+1):
    print(str("---"), end="")
print()
max_val = 0
index_list = []
for j, line in enumerate(all_pairs_shortest_path(n, edges)):
    print((str(j+1)+"|").rjust(3), end="")
    for i, v in enumerate(line):
        print(str(v).rjust(3), end="")
        if v > max_val:
            index_list = [(j+1, i+1)]
            max_val = v
        elif v == max_val:
            index_list.append((j+1, i+1))
    print()
print ()
for i in index_list:
    print ("length", i, "=", max_val)
