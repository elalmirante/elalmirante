# elalmirante

elalmirante is a tool to help you server deployment when you have a bunch of servers running the same application with different configuration.

## Why?
Imagine you have an application that runs on hundreds on servers
not on a cluster, but actually the same application but with different configuration and a different database connection.

How do you handle this scenario for deployment?

a simple solution is that you could make a list of every server you own and then deploy like you deploy a single server but doing it for every server you have on the list, however in the real world we have found that on fast moving software you can not always have the rigidity of a simple "deploy to all" scenario, sometimes you need to have some flexibility on which servers to deploy.

## How?
elalmirante proposes a query language, that when paired with a configuration schema (in yaml) aided by a tag system  
can make for a very flexible system for selecting which servers to deploy, and what to exclude from the pool of servers you have configured.

elalmirante gives you the flexibility to include or exclude servers based on a tag system.

## Simple example
Imagine you have 5 servers, all running the same application but for 5 different clients:

server1 -> client1  
server2 -> client2  

and so for, now server1, server2 and server3 are running on europe so you tag them with "europe", server 4 is running on america and server 5 on asia, so you tag them accordingly:

you will end with something like this:

server|tags
-|-
1|europe,germany
2|europe,finland
3|europe,uk
4|america,usa
5|asia,japan

after you also tag them with their country.

Now imagine this scenario: the european union pases a new regulation, you move fast to make that code change to comply with the regulation, however you only want to deploy to europe and monitor for a bit, if everything is good then you deploy to the rest of your services, with elalmirante you would only have to do:

$ `elalmirante deploy europe <version>`

however, you remember that the uk left the european union, so the new regulation is not yet in place for them, you can deploy to europe except uk with the following command:

$ `elalimirante deploy europe,!uk <version>`

and this would target european servers, except uk.

## Whats up with the name?

elalmirante means "The Admiral" in spanish, complying with the marine/sea theme of everyting deployment related, its also the name of a dodgy bar with exotic dancers.
