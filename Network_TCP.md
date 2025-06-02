
## Server socket BIND

When using bind(), you need to pass the desired socket-adress, which is a struct containing port and IP address (and address family).

Assuming you're programming a server that should be accessible by client applications:

The port and IP for bind() are those of the server, you basically define how the server can be contacted. The port is pretty straight forward since it's just a number. The IP address can be for example "127.0.0.1", then that server will only be available on the same computer where you're running the server program. It can also be a "real" IP address like "192.168.178.12" (in this case it's a private IP but it doesn't matter for what I'm trying to explain). If you specify the IP like that, a client can connect to the server with the port and IP. Now the part that got you confused: It's also possible that a computer/server has multiple network interfaces and thus multiple IP addresses. You can tell the server to listen on all its interfaces by setting the IP address to "INADDR_ANY".

You can also think of it like this (no real production example but easy to follow I think): If my PC is connected via VPN, it has the normal private home IP address and the VPN IP address. If I set the address for bind to one of them, the server will only be available either in my private network or in the VPN network. If I set it to INADDR_ANY, it will available to both the VPN network and my private home network.