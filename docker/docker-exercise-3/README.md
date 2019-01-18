# Solution

You do not need to modify either the go source code, or the Dockerfile.

```
# build the go executable for linux
GOOS=linux go build

# build your docker container. Make sure to replace "brendankellogg" with your Dockerhub name!
docker build -t brendankellogg/docker-exercise-3 .

# clean up the built Go executable
go clean

# run your newly built container. Make sure to replace "brendankellogg" with your Dockerhub name!
docker run -d \
-p 4000:4000 \
-e PORT=4000 \
-e FILEPATH=/secret/secret-message.txt \
-v $(pwd)/secret/:/secret/:ro \
brendankellogg/docker-exercise-3

# You should now be able to see the container on http://localhost:4000/
```

## Things To Note

`FILEPATH` expects an absolute filepath (including the file name) to the location of the file within the context of the container. That is, the filepath provided should be where that file lives inside of the container and NOT where it is on your host machine.

`-v` takes an absolute path and mounts it to an absolute path on the docker container. A quick way to get an absolute path from where you are is the `$(pwd)` which will evaluate to the current directory.

**WINDOWS Users:** `$(pwd)` can be unreliable in this case. Docker seems to not like the drive letter format that Windows uses (`C:/some/path/`) and instead wants something like `//c/some/path/`. Replace the `$(pwd)` with `//c/rest/of/path/`.

# Docker Excercise Three

This is a simple challenge that involves publishing ports and setting environment variables, as well as mounting volumes when running a Docker container.

You will **NOT** need to modify `main.go`

You will **NOT** need to modify the `Dockerfile`

## What do I do?

Build this directory into a docker container. Remember to build the go binary for Linux first!

When running your docker container you will need to:

- Set an environment variable for the PORT
- Set an environment variavle for the FILEPATH
- Publish the same port as the environment variable PORT
- Mount a volume containing `secret-message.txt` into your container corresponding to `FILEPATH`

Take another look at [Dr. Stearn's Docker tutorial](https://drstearns.github.io/tutorials/docker/) if you need hints!

## Test it out

Open a browser window and go to `localhost:<port>`. If everything went well you should see a success message. If you aren't able to connect then you there is an issue with your ports!

If your container crashed, use `docker logs <container-id>` to see what the issue is and fix it!

*If you named your container, you will need to remove it before you can start another with the same name.*

## Stop and remove the container

You likely don't need or want this container to be running any more so let's stop and remove it.

Run `docker ps` to list all currently running containers and copy the `CONTAINER ID` of the docker exercise container.

Stop the conainer:

`docker stop <container-id>` or `docker stop <container-name>`

Remove the container:

`docker rm <container-id>` `docker rm <container-name>`

*When you run `docker stop` and `docker rm` you will see the container ID printed to your terminal window. You can treat this as a success message.*

Remember that you can also stop and remove a running container in one command:

`docker rm -f <container-id>` or `docker rm -f <container-name>`

This forces the removal of the container, even if it is currently running.
