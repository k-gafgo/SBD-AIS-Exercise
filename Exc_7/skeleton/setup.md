# todo note commands

## Creating a Docker Swarm 

On the master node:
```docker swarm init --advertise-addr 192.168.178.29```

On the other nodes:
```docker swarm join --token SWMTKN-1-1c4odyydrviqze0yzfa9efxs2nvn4gsthmzh9pqlstttfrw6o6-11qdg32lxjhptrymk81sfj3r2 192.168.178.29:2377```

To gather more information about the swarm one can run: 
* ```docker info```: gives a long list of information, look for ```Swarm: active```
* ```docker node ls```: lists all nodes connected to the swarm, ```*``` indicates the node you are connected to