# gestalt

a `gestalt` is something that is made of many parts and yet is somehow more than or different from the combination of its parts - a perfect name for an installation / setup script!

This repo is setup currently to support Ubuntu 22+ as the primary operating system, and additional operating systems may be added in the future.

## Download and Setup
You can copy and paste this entire codeblock and run in your terminal.

```
# Install Taskfile
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

# Download accoutrements tarball
curl -LkSs https://api.github.com/repos/theopenlane/accoutrements/tarball -o accoutrements.tar.gz

# Create directory and extract files
sudo mkdir -p accoutrements
tar -xvf accoutrements.tar.gz --strip-components=1 -C accoutrements

# Navigate to correct directory
cd accoutrements/buildkite-gestalt
```

```
task fullinstall
```


## Add Agent Token

1. Go to `Buildkite` -> `Agents` -> `Agent Tokens`
2. Create a new agent token
3. Update the config with the new token:

```
vim /etc/buildkite-agent/buildkite-agent.cfg 
```

```
# The token from your Buildkite "Agents" page
token="xxxx"
...
```