Reference: https://www.reddit.com/r/selfhosted/comments/mfwfym/nginx_why_should_i_use_it/

NGINX is a web server with a ton of functionality, but how you would use it is most likely as a reverse-proxy.

One of the main uses of a reverse-proxy in relation to self hosting is to provide your self-hosted services with domain names (can be internal / LAN only) so that they are more easily accessible. So instead of running Jellyfin at 10.1.1.2:8008 and Nextcloud at 10.1.1.2:9009, yoy could create an internal domain name i.e. homeserver.lan and have Jellyfin accessible at jellyfin.homeserver.lan and Nextcloud accessible at nextcloud.homeserver.lan, no more remembering port numbers and IP addresses! This will require a little tinkering in your home router (adding a DNS entry so all requests to *.homeserver.lan redirect to 10.1.1.2) but once it's set it it makes it much easier to manage lots of services.

Additionally this sets you up for later when you might want to use SSL (https), with a reverse-proxy you can set up your TLS/SSL in a single place rather than in each service you are running.

Lastly, NGINX is extremely stable and well documented as it has been around for a long time so you should be able to find plenty of posts about it (and SO questions if you have issues). But there are some other reverse-proxies out there which are newer and have some very nice integrations with Docker, if you are already using docker for everything then it might be worth checking out one of these; Traefik and Caddy. I've used both of them and they are both great! Traefik has a little bit of a learning curve behind it when managing with docker and I found it a pain when I wanted to do something outside of docker, while Caddy is much simpler and I found it a breeze to use both within and outside a docker environment.

Hope that helps, and there are no stupid questions! One day you'll be giving these answers to someone else (i hope), we all start somewhere :)



