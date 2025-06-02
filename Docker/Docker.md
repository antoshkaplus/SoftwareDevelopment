## Commands (may require `sudo`)
* `docker run {container_name / image}`
    Pulls from the docker hub if does not exist locally.
    Add `-d` after `run` to run in detach mode, in background.
* `docker ps`
    List all running containers. Use `-a` also shows stopped or existed containers 
    since they still take disk space
    .
* `docker stop {name / id}`
* `docker rm {name}` 
    Remove stopped or exited container. Reclaims disk space from ran containers.
* `docker images`
    List available images.
* `docker rmi {name}`
    Removes image.
* `docker pull {image name}`
    Downloads image.
* `docker attach {id}`
    Pull container output from background. `id` input allows to specify just few first characters
    that differenciate it from other ids.
* `docker inspect {name}`
    Shows container configuration.
* `docker logs {name/id}`


`docker run -it centos bash` - Run container `centos` and execute `bash` command in it. 
    `-it` gets into container command line?

By default `docker` requires `sudo` to use Unix sockets. 
The following talks about remedies: 
https://askubuntu.com/questions/477551/how-can-i-use-docker-without-sudo


`docker run` in non-interactive mode, does not read any input.
Can do interactive mode with `-i` flag: `docker run -i {container-name}`
Also add `-t` to attach terminal, so in the end you end up with `-it`.


Every docker container gets IP asigned by default. 
Internal IP available only under docker host.


There is also `-v` option.


## Create my own image.
* `docker build` 
    use `-t` option to provide image tag.