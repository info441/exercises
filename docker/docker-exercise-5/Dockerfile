# Put your docker build instructions here!

# use the scratch image since we do not need any linux utilities
FROM scratch

# add in the files that the go web server will be serving
# we can safely copy these into the container because
# there are no sensitive information inside of these files
# so we don't care if someone decompiles our container
# and gets access to them.
#
# This line will copy the contents of the "static" folder
# and put them in the "/static" directory within the container.
COPY static /static

# copy the executable that we'll be running into the container
COPY docker-exercise-5 docker-exercise-5

# tell docker to use the executable as the main entrypoint when the
# container starts. That is, what should the container do when it starts?
ENTRYPOINT [ "/docker-exercise-5" ]
