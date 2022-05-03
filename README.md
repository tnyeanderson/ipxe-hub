# Pixie

Web interface for managing PXE booted clients

PROJECT IS IN ALPHA. MORE DOCUMENTATION AND FUNCTIONALITY TO COME

## Features
- [x] TFTP server (no docker support yet)
- [x] Web interface & HTTP API
- [x] Chainload iPXE boot image generator
- [x] Upload/manage scripts and boot images
- [x] Devices which try to network boot are added to the database
- [x] Assign a default boot script
- [x] Assign scripts to devices
- [x] Logging (audit trail)
- [x] Upload/manage config files (ignition, cloud-init, etc)
- [ ] Config file (ignition, cloud-init, etc) handlebars templating
- [ ] Build a boot script (select images, config) with a GUI
- [ ] Authentication (RBAC?)
- [ ] Groups (assign a script/config to a group of devices)
- [ ] Boot script handlebars templating
- [ ] Library of boots (ex: ubuntu, debian, coreos, etc) that can be "imported" and modified
- [ ] Temporarily boot another script for testing/debugging (for x number of boots, or y seconds, etc)
- [ ] Disable TFTP server? (use your own)

## Getting started
First, generate a `pixie.kpxe` file that will be the default boot image sent by the DHCP server. The URL used below will be the [chainloaded](https://ipxe.org/howto/chainloading) host.

This requires building iPXE which requires a fair amount of dependencies. A docker image is defined here for convenience:
```bash
# Build the docker image used to generate the kpxe file
docker build -f Dockerfile.generate-kpxe -t pixie-kpxe-generator .
```

The generated file will be placed inside the container at `/output/pixie.kpxe`, and it needs to be placed in the TFTP root directory using volume mounts.

If using the default configuration, `Paths.FileServer` (the TFTP root) is `data/files`, so this will be used as an example.
```bash
# Create the local directory that the generated file will be copied into
mkdir -p data/files
# Change the localhost address below to a host that will resolve to Pixie!
docker run -it -v "$(pwd)/data/files:/output" pixie-kpxe-generator 'http://localhost:8880'
```

Then, run the app!

## Run using docker
This is the recommended method:
```bash
docker-compose up -d
```

Then navigate to `localhost:8880` in a browser.

>NOTE: The image can be rebuilt/updated with `docker-compose build`

## Run locally
To run the app without docker, first build the components:
```bash
# Build the web interface
(cd web/ && ng build)
# Build the app
(cd src/ && go build -o ../pixie)
```

Then run the app:
```bash
# TFTP server (running on a privileged port) is included
# Therefore sudo must be used.
sudo ./pixie
```

Then navigate to `localhost:8880` in a browser.

## Configuration

**The default configuration should be sufficient for most users.**

To see configuration options and their explanations, see the [default config](config/default.yaml)

To tweak the config, use the default config as a template:
```bash
cp config/default.yaml data/pixie.yaml
```

Then edit `data/pixie.yaml` as needed.

Or, load a different config file:
```bash
./pixie --config-file /path/to/config
```
