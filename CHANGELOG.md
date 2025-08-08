# Changelog

## 0.7.0 (2025-08-08)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/prelude-so/go-sdk/compare/v0.6.0...v0.7.0)

### Features

* **api:** update via SDK Studio ([e6b3ca1](https://github.com/prelude-so/go-sdk/commit/e6b3ca1528c7726398987af37c1f079b84f31fb4))
* **client:** support optional json html escaping ([8f9d8a4](https://github.com/prelude-so/go-sdk/commit/8f9d8a472bdbd401d01f035e6b2868589cad555c))


### Bug Fixes

* **client:** process custom base url ahead of time ([3f1d80f](https://github.com/prelude-so/go-sdk/commit/3f1d80f7d1931c767188c384499e6972e00a3b78))
* don't try to deserialize as json when ResponseBodyInto is []byte ([cf385aa](https://github.com/prelude-so/go-sdk/commit/cf385aae3a726932291f4c9893153e347f26d610))


### Chores

* **ci:** only run for pushes and fork pull requests ([0da7840](https://github.com/prelude-so/go-sdk/commit/0da7840eb5ef751b927a0759b13a5a972a96a036))
* **internal:** fix lint script for tests ([5412ddc](https://github.com/prelude-so/go-sdk/commit/5412ddc451370366f3995b6247dd8d5bd4e5cbf5))
* **internal:** update comment in script ([bd5b3c3](https://github.com/prelude-so/go-sdk/commit/bd5b3c383dbeee8b28a78145b105b6092829b1c3))
* lint tests ([ebc4caa](https://github.com/prelude-so/go-sdk/commit/ebc4caa7ca3a7c4e4d8c673c25d6e83d2efe5609))
* lint tests in subpackages ([9d28dc7](https://github.com/prelude-so/go-sdk/commit/9d28dc7de841d96d18f6d3f55f471129fad9e749))
* update @stainless-api/prism-cli to v5.15.0 ([640c13b](https://github.com/prelude-so/go-sdk/commit/640c13b615f6a2956443015c707d38dab2196672))

## 0.6.0 (2025-06-17)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/prelude-so/go-sdk/compare/v0.5.0...v0.6.0)

### Features

* **client:** add debug log helper ([bdc01e5](https://github.com/prelude-so/go-sdk/commit/bdc01e53e36b3aa16aaf7c5bbcc55e8994598c88))


### Chores

* **ci:** enable for pull requests ([3b46ccd](https://github.com/prelude-so/go-sdk/commit/3b46ccdb10b8d4f25365a75ba5aed537af5a30f0))

## 0.5.0 (2025-06-02)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/prelude-so/go-sdk/compare/v0.4.0...v0.5.0)

### Features

* **api:** update via SDK Studio ([9e5bbd6](https://github.com/prelude-so/go-sdk/commit/9e5bbd6d5d95b3920ca0d331dc5ff01e4d6aaff0))
* **client:** add support for endpoint-specific base URLs in python ([8a88868](https://github.com/prelude-so/go-sdk/commit/8a88868b6e87eb44042fb84eb055563dd1b833d6))


### Chores

* **docs:** grammar improvements ([e298d78](https://github.com/prelude-so/go-sdk/commit/e298d784e4b48e344ced24d11475ef448678a037))
* improve devcontainer setup ([7e19069](https://github.com/prelude-so/go-sdk/commit/7e19069917354f694ae6f224d8e9f182a3e58a6a))
* make go mod tidy continue on error ([986e865](https://github.com/prelude-so/go-sdk/commit/986e865c53cb0aa3e93ba5e630211f4b51eac1cc))

## 0.4.0 (2025-05-13)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/prelude-so/go-sdk/compare/v0.3.0...v0.4.0)

### Features

* **api:** update via SDK Studio ([c5cb0f8](https://github.com/prelude-so/go-sdk/commit/c5cb0f8243ed85f87691c55b45208beea508809e))
* **client:** add support for reading base URL from environment variable ([6425aaa](https://github.com/prelude-so/go-sdk/commit/6425aaa98169585728a22b9dafd8e22850e5bcdb))


### Bug Fixes

* **client:** clean up reader resources ([29c742a](https://github.com/prelude-so/go-sdk/commit/29c742ab17f2d18bc502792a120e717edfd2f171))
* **client:** correctly update body in WithJSONSet ([e3a7f75](https://github.com/prelude-so/go-sdk/commit/e3a7f755580e3a15ea7839cfc77b39e908def51c))
* handle empty bodies in WithJSONSet ([960afe6](https://github.com/prelude-so/go-sdk/commit/960afe66f028812a1d8dd2d2ea627ac5044ea583))


### Chores

* **ci:** add timeout thresholds for CI jobs ([26b3f82](https://github.com/prelude-so/go-sdk/commit/26b3f827e738009e9e441d7b90d71b0f8ee7d172))
* **ci:** only use depot for staging repos ([7dd3e52](https://github.com/prelude-so/go-sdk/commit/7dd3e52a5c51b0aaf5c513764bc4fbdd9af1f795))
* **ci:** run on more branches and use depot runners ([54428fe](https://github.com/prelude-so/go-sdk/commit/54428fe0808ad1877990d8d88c4e468f3e6cca84))
* **docs:** document pre-request options ([e87f9ea](https://github.com/prelude-so/go-sdk/commit/e87f9ea11681827ee8b3e8c50f938fbd163a1fda))


### Documentation

* update documentation links to be more uniform ([d7bfab0](https://github.com/prelude-so/go-sdk/commit/d7bfab07dfe42cfca734a3becc72f232a11bca55))

## 0.3.0 (2025-04-11)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/prelude-so/go-sdk/compare/v0.2.0...v0.3.0)

### Features

* **api:** update via SDK Studio ([838fe81](https://github.com/prelude-so/go-sdk/commit/838fe81c948b2898e2bf3240e96c5638a2f2ac36))
* **api:** update via SDK Studio ([f1241b1](https://github.com/prelude-so/go-sdk/commit/f1241b1249182e053f98e05b67becabc3b4ab384))
* **client:** improve default client options support ([#42](https://github.com/prelude-so/go-sdk/issues/42)) ([0e6914b](https://github.com/prelude-so/go-sdk/commit/0e6914b94ae832cd1f4c69874713ea0d556eb39f))
* **client:** support custom http clients ([#50](https://github.com/prelude-so/go-sdk/issues/50)) ([e1c08b6](https://github.com/prelude-so/go-sdk/commit/e1c08b6e64fa387071e93c0c9bd564c862d6b131))


### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#49](https://github.com/prelude-so/go-sdk/issues/49)) ([adc11db](https://github.com/prelude-so/go-sdk/commit/adc11db706c3b32e5a9a7c505b744e85c6edcddd))
* **test:** return early after test failure ([#47](https://github.com/prelude-so/go-sdk/issues/47)) ([9a44846](https://github.com/prelude-so/go-sdk/commit/9a448469de751b96447822b0bf73c5fe0f2fc177))


### Chores

* add request options to client tests ([#46](https://github.com/prelude-so/go-sdk/issues/46)) ([de700e5](https://github.com/prelude-so/go-sdk/commit/de700e5c622cb8e869c6bc792dfc8dcf91b56d47))
* **docs:** improve security documentation ([#45](https://github.com/prelude-so/go-sdk/issues/45)) ([7e200cb](https://github.com/prelude-so/go-sdk/commit/7e200cbb0bd779490cd6911cc620240260d06329))
* fix typos ([#48](https://github.com/prelude-so/go-sdk/issues/48)) ([ae49898](https://github.com/prelude-so/go-sdk/commit/ae49898459f1ad2bfb4916edebadc81678c38cd4))
* **internal:** expand CI branch coverage ([#52](https://github.com/prelude-so/go-sdk/issues/52)) ([a814dc5](https://github.com/prelude-so/go-sdk/commit/a814dc5d46d99350187dd30117b22659c15e0073))
* **internal:** reduce CI branch coverage ([b723f9d](https://github.com/prelude-so/go-sdk/commit/b723f9d7ab07502b19187a58e56f57794abe2288))
* **internal:** remove extra empty newlines ([#44](https://github.com/prelude-so/go-sdk/issues/44)) ([9d04921](https://github.com/prelude-so/go-sdk/commit/9d04921cd42e9df3416c8a79831c973eaeb026b8))
* **tests:** improve enum examples ([#51](https://github.com/prelude-so/go-sdk/issues/51)) ([8a3b337](https://github.com/prelude-so/go-sdk/commit/8a3b3370c962930777f9612b9c71be7a4bbd1d11))

## 0.2.0 (2025-03-11)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/prelude-so/go-sdk/compare/v0.1.0...v0.2.0)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#37](https://github.com/prelude-so/go-sdk/issues/37)) ([c3f0d41](https://github.com/prelude-so/go-sdk/commit/c3f0d41329dae63f04ae339e99eed1cec9615a3f))
* **api:** update via SDK Studio ([#40](https://github.com/prelude-so/go-sdk/issues/40)) ([33ad2ad](https://github.com/prelude-so/go-sdk/commit/33ad2ad2fbde7526730303e0a6e5c4fa14af9af8))
* **client:** accept RFC6838 JSON content types ([#38](https://github.com/prelude-so/go-sdk/issues/38)) ([a3b19b1](https://github.com/prelude-so/go-sdk/commit/a3b19b13d6ac2eae424574e7260c635804eaf67b))
* **client:** allow custom baseurls without trailing slash ([#36](https://github.com/prelude-so/go-sdk/issues/36)) ([2830cfb](https://github.com/prelude-so/go-sdk/commit/2830cfbaedcf5468e87c412e7a4a4842beaa1b5a))
* **client:** send `X-Stainless-Timeout` header ([#27](https://github.com/prelude-so/go-sdk/issues/27)) ([7d81306](https://github.com/prelude-so/go-sdk/commit/7d8130689b1ff542a81405fc75b54bcc8ed59b7c))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#31](https://github.com/prelude-so/go-sdk/issues/31)) ([96fdb15](https://github.com/prelude-so/go-sdk/commit/96fdb150cc6e186c5a1341632bbf7a9f66870700))
* do not call path.Base on ContentType ([#30](https://github.com/prelude-so/go-sdk/issues/30)) ([bc311ee](https://github.com/prelude-so/go-sdk/commit/bc311ee19fd8fd5d98794451d54566dcf508f5ba))
* fix early cancel when RequestTimeout is provided for streaming requests ([#29](https://github.com/prelude-so/go-sdk/issues/29)) ([f690d25](https://github.com/prelude-so/go-sdk/commit/f690d25deaf8add9133537ad4ad7c350ade0c8af))
* fix unicode encoding for json ([#25](https://github.com/prelude-so/go-sdk/issues/25)) ([1dbf55b](https://github.com/prelude-so/go-sdk/commit/1dbf55b8eca9c3f86a61fd35b68a5f1e17fc5b92))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#28](https://github.com/prelude-so/go-sdk/issues/28)) ([32c01d0](https://github.com/prelude-so/go-sdk/commit/32c01d07dd4d9d7cdef988572e523f979cd8c544))
* **internal:** codegen related update ([#23](https://github.com/prelude-so/go-sdk/issues/23)) ([bd068d9](https://github.com/prelude-so/go-sdk/commit/bd068d94991e2a5fa38f20bd48012d0d6a8578d0))
* **internal:** codegen related update ([#33](https://github.com/prelude-so/go-sdk/issues/33)) ([b120ab9](https://github.com/prelude-so/go-sdk/commit/b120ab94db58004c6d135fe3fc5b7558fa1caaa5))
* **internal:** codegen related update ([#34](https://github.com/prelude-so/go-sdk/issues/34)) ([991c6d2](https://github.com/prelude-so/go-sdk/commit/991c6d23eaed8597a2d0f5f3594f2c4b3bbc0c29))
* **internal:** fix devcontainers setup ([#32](https://github.com/prelude-so/go-sdk/issues/32)) ([56a3b27](https://github.com/prelude-so/go-sdk/commit/56a3b27fc822f3085b33d9f816016c507d55c92f))


### Documentation

* document raw responses ([#26](https://github.com/prelude-so/go-sdk/issues/26)) ([17fe9f9](https://github.com/prelude-so/go-sdk/commit/17fe9f9d4f4a68cf23068bd104f4dd5dada53704))
* update URLs from stainlessapi.com to stainless.com ([#35](https://github.com/prelude-so/go-sdk/issues/35)) ([fea942e](https://github.com/prelude-so/go-sdk/commit/fea942e0da9c3c8f0d766a3fe1f7e195512ea597))


### Refactors

* tidy up dependencies ([#39](https://github.com/prelude-so/go-sdk/issues/39)) ([5079ac1](https://github.com/prelude-so/go-sdk/commit/5079ac1f780f30fc27a18da83530604603b0c3b8))

## 0.1.0 (2025-02-05)

Full Changelog: [v0.1.0-beta.1...v0.1.0](https://github.com/prelude-so/go-sdk/compare/v0.1.0-beta.1...v0.1.0)

### Features

* **api:** update via SDK Studio ([#20](https://github.com/prelude-so/go-sdk/issues/20)) ([4753580](https://github.com/prelude-so/go-sdk/commit/47535809952acd1228a7a83d541845b6ee68009a))

## 0.1.0-beta.1 (2025-01-14)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-beta.1](https://github.com/prelude-so/go-sdk/compare/v0.1.0-alpha.4...v0.1.0-beta.1)

### Features

* **api:** update via SDK Studio ([#18](https://github.com/prelude-so/go-sdk/issues/18)) ([3264fcd](https://github.com/prelude-so/go-sdk/commit/3264fcdb8104274178ed78342b01a9512d1fa1ca))


### Chores

* **internal:** codegen related update ([#16](https://github.com/prelude-so/go-sdk/issues/16)) ([a3f1fef](https://github.com/prelude-so/go-sdk/commit/a3f1fefac1137c42aa92e7c3d4b31b1f3160f75a))

## 0.1.0-alpha.4 (2025-01-08)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/prelude-so/go-sdk/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** update via SDK Studio ([#14](https://github.com/prelude-so/go-sdk/issues/14)) ([2b1929e](https://github.com/prelude-so/go-sdk/commit/2b1929e7a5941b1217d706bcd48049989584c80b))


### Chores

* **internal:** codegen related update ([#11](https://github.com/prelude-so/go-sdk/issues/11)) ([f65ed56](https://github.com/prelude-so/go-sdk/commit/f65ed56b7eafc300fa9b7877cd8b814ac9c74478))
* **internal:** codegen related update ([#13](https://github.com/prelude-so/go-sdk/issues/13)) ([bd4db42](https://github.com/prelude-so/go-sdk/commit/bd4db42a9bbf3ea5f8cc0f14931ce0bab51e6b8d))

## 0.1.0-alpha.3 (2024-12-11)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/prelude-so/go-sdk/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([#8](https://github.com/prelude-so/go-sdk/issues/8)) ([6e56add](https://github.com/prelude-so/go-sdk/commit/6e56add5fe0c793a6fa9295236a6fc65cdfcc9e2))

## 0.1.0-alpha.2 (2024-11-27)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/prelude-so/go-sdk/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([#4](https://github.com/prelude-so/go-sdk/issues/4)) ([32e594a](https://github.com/prelude-so/go-sdk/commit/32e594ab0917eadf6760b6f820afe753d7f5eb2a))
* **api:** update via SDK Studio ([#6](https://github.com/prelude-so/go-sdk/issues/6)) ([5298212](https://github.com/prelude-so/go-sdk/commit/5298212a02fab3335bb035bc7a9fb95922dbd160))

## 0.1.0-alpha.1 (2024-11-13)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/prelude-so/go-sdk/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** api update ([1dfd23d](https://github.com/prelude-so/go-sdk/commit/1dfd23dd5f0b35b636a9cabd9769271a47da1c92))
* **api:** update via SDK Studio ([96bf346](https://github.com/prelude-so/go-sdk/commit/96bf346a64334a6e1132aab00119f4c5f2df6695))
* **api:** update via SDK Studio ([fcf2e15](https://github.com/prelude-so/go-sdk/commit/fcf2e154e74dd08214ca6cb61c86abdf69c2bccd))
* **api:** update via SDK Studio ([e25a9f7](https://github.com/prelude-so/go-sdk/commit/e25a9f7d68dd335cd7c92b2a19693bd88fe5f63a))
* **api:** update via SDK Studio ([b1ce3a9](https://github.com/prelude-so/go-sdk/commit/b1ce3a92fd3f6219db46096c07e9e6f855a803c0))
* **api:** update via SDK Studio ([707b20d](https://github.com/prelude-so/go-sdk/commit/707b20d83ef46a6acf6d0028e1e163ce6f8cec52))
* **api:** update via SDK Studio ([4b2b93e](https://github.com/prelude-so/go-sdk/commit/4b2b93e66aaf28a5301cf1f63723e189a0deceed))
* **api:** update via SDK Studio ([3284a41](https://github.com/prelude-so/go-sdk/commit/3284a41c9c8cc1652ec74f43c7803915d2780db2))
* **api:** update via SDK Studio ([eccf0b8](https://github.com/prelude-so/go-sdk/commit/eccf0b8775446e19b8bd7d65f3fcae0df5e46aa9))
* **api:** update via SDK Studio ([c7b0129](https://github.com/prelude-so/go-sdk/commit/c7b012927950a89dec95416045e0a6871de78728))


### Chores

* configure new SDK language ([08c784d](https://github.com/prelude-so/go-sdk/commit/08c784d7986ce6dba9018638e1a6ff7de0fead00))
* go live ([#1](https://github.com/prelude-so/go-sdk/issues/1)) ([1a6b1b6](https://github.com/prelude-so/go-sdk/commit/1a6b1b6d566663a9a700107b5d1c17f1464e3285))
* rebuild project due to codegen change ([2c875a2](https://github.com/prelude-so/go-sdk/commit/2c875a244f28e7ee22c33f777ef2b003241a97d4))
* rebuild project due to codegen change ([9eebdee](https://github.com/prelude-so/go-sdk/commit/9eebdee594ed853ac3a12713caa153e64bd8ed06))
