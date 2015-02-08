# envconf

> Hi! I'm `envconf`. Please improve me.

## Build

    $ sudo docker run --cidfile=cid -v $(pwd):/go/src/app golang:1.4-onbuild \
      bash -c "go-wrapper download && go-wrapper install"
      
    $ docker cp $(cat cid):/go/bin/app .
    $ mv app envconf
    $ chmod +x envconf
    
## Just install

    $ sudo curl -L -o /usr/local/bin/envconf https://github.com/webwurst/envconf/releases/download/v0.1-alpha/envconf
    $ sudo chmod +x /usr/local/bin/envconf
    
## Use it

..like this:

    $ envconf some_file.txt
    
Envconf reads in `some_file.txt.env` and replaces all environment variables. Result is written to `some_file.txt`.
