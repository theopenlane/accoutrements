# accoutrements

This repository is not intended to be totally comprehensive of every need, or used as a source for a git submodule but rather a place to store common files that are used in many projects that can be copy / pasted or serve as a starting point for a new project if you may be unsure as to the "right way" to set something up.

As a rule of thumb, nearly all the files and configurations in this repository should be in most repos containing code or which interact with Openlane's environment. However, given the variability or different needs of projects at any given time, instead of having a template repo setup that requires you to update the repository name in multiple places and/or be required to scrub out files we've elected to have a very basic `base` template repository that should effectively never have to have files removed from it and this repo can then be used to add anything additional on top.

For convenience and coding fun, there is a small script in this repo which will take a repo name as input and spit out a full "copypasta" set of configurations for a new repository, hopefully alleviating the amount of button clicking needed.

## Layout overview

- actions-pipelines: commonly used github actions pipelines and their respective configurations
- buildkite-pipelines: commonly used buildkite pipelines and their respective configurations with the file names indicating the use (and a "full" one you could trim back if needed)
- config-generation: small golang files leveraging the koanf project + jsonschema to generate local `.config` files, API docs, deployment (kuberentes configmap mainly) artifacts based on go struct tags
- github-repo: things specific to github repositories such as `CONTRIBUTING.md`, `LICENSE`, `.gitignore` and similar
- golang: configuration files and scripts for golang projects, e.g. `.goreleaser`, `golanglint`
- pre-commit: pre-commit configurations and pre-commit hook configuration files (e.g. typos config, yamlfmt config, etc.)
- scanning-utilities: configuration files for utilities performing scanning (and reporting) on a repository such as sonarlint config, renovate config, etc.
- taskfiles: probably the most useful pieces inside of this repo, Taskfiles have been broken out and named according to their purpose (ex: `Taskfile_localinstallation.yaml` for local installation, `Taskfile_ci.yaml` for CI/CD, `Taskfile_local.yaml` for local development, etc.). If you pull files from any of the above listed directories you will likely want the corresponding taskfile as well. There is a `full` taskfile with just about everything you could want, and the copypasta script will substitute the repository name and pathing according to the structure which is spit out.

## Notes

There is no panacea to consistent configurations across repositories. There are many ways to approach this problem, and a simple approach (that's slightly manual) was chosen intentionally. We do a tremendous amount of code generation and automation across repositories in the Openlane github organization which means while we could apply fancy code towards a "solution" to this problem, we are intentionally and knowingly not doing so due to experience with overrotating on automation that ends up being more work and a headache for everyone. 

Want to automate stuff using these files any way? Feel free! Just do it in a different repository.
