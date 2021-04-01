# Changelog

## 1.1.0 (2021/02/18)

* Refactor CI and dev workflow with buildx bake (#54)
    * Add `image-local` target
    * Single job for artifacts and image
    * Add `armv5`, `ppc64le` and `s390x` artifacts
* Remove `linux/s390x` Docker platform support for now
* Bump github.com/jfrog/jfrog-client-go from 0.16.0 to 0.19.1 (#36 #38 #39 #43 #49)
* MkDocs Materials 6.2.8
* Bump github.com/alecthomas/kong from 0.2.12 to 0.2.15 (#52)
* Bump github.com/stretchr/testify from 1.6.1 to 1.7.0 (#45)
* Bump github.com/go-resty/resty/v2 from 2.3.0 to 2.4.0 (#44)

## 1.0.1 (2020/12/15)

* Use `tag-semver`
* Bump github.com/alecthomas/kong from 0.2.11 to 0.2.12 (#34)

## 1.0.0 (2020/11/16)

* Use embedded tzdata package
* Remove `--timezone` flag
* Docker image also available on [GitHub Container Registry](https://github.com/users/crazy-max/packages/container/package/artifactory-cleanup)
* Switch to Docker actions
* Update deps

## 0.3.1 (2020/09/24)

* Fix logger

## 0.3.0 (2020/09/24)

* Allow to disable scheduling to execute policies right away
* Update deps

## 0.2.0 (2020/09/23)

* Add docs
* Rename `generic` policy type to `common`
* Fix logging settings

## 0.1.0 (2020/09/16)

* Initial version
