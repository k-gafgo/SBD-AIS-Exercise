# todo note commands

## Creating a Docker Swarm 

On the master node:
```docker swarm init --advertise-addr <ipv4-address>```

On the other nodes:
```docker swarm join --token <token>```

To gather more information about the swarm one can run: 
* ```docker info```: gives a long list of information, look for ```Swarm: active```
* ```docker node ls```: lists all nodes connected to the swarm, ```*``` indicates the node you are connected to