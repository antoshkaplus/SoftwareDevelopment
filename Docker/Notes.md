
---

Docker images (including the ubiquitous ubuntu and debian images) don’t contain kernels, and containers based on them don’t run kernels; they always share the host kernel. You need to go through special setup to be able to use host devices, regardless of what /dev special files exist in a container.

You can make an argument that the ubuntu image has unnecessary components that aren’t used in typical Docker setups (an init system, a DHCP client) that make it closer to a “full OS image”; or you can make an argument that, since ubuntu doesn’t have its own kernel, it’s “not an OS”. I’m not really sure how “operating system” is defined these days.

---

The operating system of the host is not the operating system running in the container. That is the beauty of using containers. For example: the OS of the host could be RedHat and the OS of the container could be Ubuntu. The only thing that the host and container share is the kernel of the host. The filesystems and everything else are of their respective operating systems. Applications running in an Ubuntu container are running on Ubuntu. Applications running in an Alpine container are running on Alpine. It doesn’t matter that the host might be CentOS or any other Linux distro. Every container is running some linux distro that does not have to be the host distro. Hopefully that helps to clarifiy things for you. If not, ask more questions.