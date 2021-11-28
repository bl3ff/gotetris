TO RUN FROM DOCKER (linux -env) 
docker run -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=:0 -h HeLios -v /home/davide/.Xautority:/home/bl3ff/.Xautority <image-name>
