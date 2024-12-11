# gestalt

a `gestalt` is something that is made of many parts and yet is somehow more than or different from the combination of its parts - a perfect name for an installation / setup script!

This repo is setup currently to support Ubuntu 22+ as the primary operating system, and additional operating systems may be added in the future.

## Download and Setup

```
sudo apt install age
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
curl -LkSs https://api.github.com/repos/theopenlane/accoutrements/tarball -o accoutrements.tar.gz
sudo mkdir accoutrements
tar -xvf gestalt.tar.gz --strip-components=1 -C accoutrements
cd accoutrements/buildkite-gestalt
age -d .env.age >> .env
```

Then enter the passphrase. Once this is done, run:

```
task fullinstall
```

## Secrets

The `.env.age` file in this repository is encrypted using [age](https://github.com/FiloSottile/age); you can setup `task` as outlined below and run `task install:age` or simply run `apt install age` on the server, or decrypt the file locally vs. moving to the server. In either case, you'll need to run `age -d .env.age` and enter the passphrase which will print out the secret values required for the scripts / installation.

### Encrypting an .env file

Create a new file or update the existing by first decrypting it per the steps above. Whenever you're ready to encrypt, simply run the `age` command and if the desire is to use a passphrase like is currently setup for (since the repo is intended to be cloned onto a server which may or may not be able to have / use different keys or otherwise) you can run:


```
age -p .env > .env.age
age -d .env.age >> .env-mitb
```

