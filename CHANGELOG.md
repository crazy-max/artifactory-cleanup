# Changelog

## 1.4.0 (2021/07/05)

* Add `linux/riscv64` artifact
* Alpine Linux 3.14
* Bump github.com/jfrog/jfrog-client-go from 0.23.1 to 1.0.1
* Update module github.com/rs/zerolog to v1.23.0 (#86)
* Update golangci/golangci-lint Docker tag to v1.41 (#84 #87)
* MkDocs Materials 7.1.9
* Update module github.com/alecthomas/kong to v0.2.17 (#82)
* Update github.com/hako/durafmt commit hash to 5c1018a (#81)

## 1.3.0 (2021/06/06)

* MkDocs Materials 7.1.5
* Allow disabling log color output
* Add `NO_COLOR` support
* Update build workflow
* Fix artifacts download links
* Set `cacheonly` output for validators
* Bump github.com/rs/zerolog from 1.21.0 to 1.22.0 (#75)
* Remove vendor workflow
* Bump github.com/go-playground/validator/v10 from 10.5.0 to 10.6.1 (#72 #74)
* Bump github.com/jfrog/jfrog-client-go from 0.21.3 to 0.23.1 (#73 #78)
* Move to `docker/metadata-action`
* Add `darwin/arm64` artifact

## 1.2.0 (2021/04/25)

* MkDocs Materials 7.1.3
* Bump github.com/jfrog/jfrog-client-go from 0.19.1 to 0.21.3 (#57 #58 #69 #70)
* Bump github.com/go-resty/resty/v2 from 2.4.0 to 2.6.0 (#67)
* Bump github.com/go-playground/validator/v10 from 10.4.1 to 10.5.0 (#66)
* Bump github.com/rs/zerolog from 1.20.0 to 1.21.0 (#59)
* Go 1.16 (#60)
* Fix CodeQL workflow
* Bump github.com/alecthomas/kong from 0.2.15 to 0.2.16 (#56)
* Switch to goreleaser-xx (#55)

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

* Allow disabling scheduling to execute policies right away
* Update deps

## 0.2.0 (2020/09/23)

* Add docs
* Rename `generic` policy type to `common`
* Fix logging settings

## 0.1.0 (2020/09/16)

* Initial version
