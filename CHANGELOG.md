# Changelog

## 0.13.0 (2026-05-12)

Full Changelog: [v0.12.1...v0.13.0](https://github.com/prelude-so/go-sdk/compare/v0.12.1...v0.13.0)

### Features

* **api:** api update ([3205cd8](https://github.com/prelude-so/go-sdk/commit/3205cd8ee44e7327e2e6d751c28f564845306a3b))
* **api:** api update ([b10173f](https://github.com/prelude-so/go-sdk/commit/b10173f411659971e5e932d9c20579258bc48643))

## 0.12.1 (2026-05-07)

Full Changelog: [v0.12.0...v0.12.1](https://github.com/prelude-so/go-sdk/compare/v0.12.0...v0.12.1)

### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([4b5c5ea](https://github.com/prelude-so/go-sdk/commit/4b5c5ea548f01a636d31bbcb0783823a083edd4a))


### Chores

* redact api-key headers in debug logs ([b609c12](https://github.com/prelude-so/go-sdk/commit/b609c1280466af6f8cafe3784db1f10925bbde5d))

## 0.12.0 (2026-05-07)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/prelude-so/go-sdk/compare/v0.11.0...v0.12.0)

### Features

* **api:** api update ([38ea2f9](https://github.com/prelude-so/go-sdk/commit/38ea2f9ec3f87f98e3d4a3f979daaa2ce14c8bce))
* **api:** api update ([bded950](https://github.com/prelude-so/go-sdk/commit/bded95008bb42313fc1b17cadee8abdfcad6184e))
* **api:** api update ([6a5a290](https://github.com/prelude-so/go-sdk/commit/6a5a290682f800bf2984a3110f325f5b1bf5e073))
* **api:** api update ([b92e164](https://github.com/prelude-so/go-sdk/commit/b92e164cff63f94d7d28c41e65f129e214e1064c))
* **api:** api update ([07ef577](https://github.com/prelude-so/go-sdk/commit/07ef5773c75408c5664688b936ce20652879ee88))
* **api:** api update ([89f3e7b](https://github.com/prelude-so/go-sdk/commit/89f3e7bbb2b5e9e4ec4b14b9bbcedd46aa0faf7a))
* **api:** api update ([5fcb192](https://github.com/prelude-so/go-sdk/commit/5fcb1921c9e4796e3064713c20e5382edb189089))
* **go:** add default http client with timeout ([c128022](https://github.com/prelude-so/go-sdk/commit/c12802247c72a6150d8716fd9c78a6daed4004d8))
* **internal:** support comma format in multipart form encoding ([17b3d4f](https://github.com/prelude-so/go-sdk/commit/17b3d4fe2df4a284fad0e153fd346ca26df07afc))
* support setting headers via env ([8d7b878](https://github.com/prelude-so/go-sdk/commit/8d7b8783e456a32d3721ab80233efab8d5291e32))


### Bug Fixes

* prevent duplicate ? in query params ([3d7ec7c](https://github.com/prelude-so/go-sdk/commit/3d7ec7cbeeeb2eb278faa3cd311712c379857a06))


### Chores

* avoid embedding reflect.Type for dead code elimination ([8317544](https://github.com/prelude-so/go-sdk/commit/8317544004a608839ad363a1158b51bc0f25898a))
* **ci:** skip lint on metadata-only changes ([25d23e6](https://github.com/prelude-so/go-sdk/commit/25d23e635ca329f1dd2d9dc8476e4eeda83355b3))
* **ci:** support opting out of skipping builds on metadata-only commits ([edc443c](https://github.com/prelude-so/go-sdk/commit/edc443c000bde1eac192756fb4bbb103dad48847))
* **internal:** more robust bootstrap script ([84aa58c](https://github.com/prelude-so/go-sdk/commit/84aa58c7006a44c244b67dd6ea60f1230f0fe5e4))
* remove unnecessary error check for url parsing ([8e23ec3](https://github.com/prelude-so/go-sdk/commit/8e23ec37d2f5d269c4bc3842662f972959d90917))
* **tests:** bump steady to v0.19.6 ([031ee47](https://github.com/prelude-so/go-sdk/commit/031ee471dc68bcc7e9e25971c1ae48f1261aa9d8))
* **tests:** bump steady to v0.19.7 ([6566496](https://github.com/prelude-so/go-sdk/commit/6566496ff718e2d1fbc50b9ec8ecdd8c34819753))
* **tests:** bump steady to v0.20.1 ([2af98ac](https://github.com/prelude-so/go-sdk/commit/2af98acfdd3ca5165bfe5770882abdbac8f3341e))
* **tests:** bump steady to v0.20.2 ([12e4a63](https://github.com/prelude-so/go-sdk/commit/12e4a6394c78ca8b3997f1e1c4b4b56761395760))
* **tests:** bump steady to v0.22.1 ([7ee8d5e](https://github.com/prelude-so/go-sdk/commit/7ee8d5e0352760fcacfbe92f35c976ae4b9e1c81))

## 0.11.0 (2026-03-23)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/prelude-so/go-sdk/compare/v0.10.0...v0.11.0)

### Features

* **api:** api update ([279a0e1](https://github.com/prelude-so/go-sdk/commit/279a0e1e8f2d28be802a7849d43244f584ff157a))
* **api:** api update ([c8e8724](https://github.com/prelude-so/go-sdk/commit/c8e872480406eb4ac6cea8cca2ee12ee13712de2))
* **api:** api update ([b49f7b3](https://github.com/prelude-so/go-sdk/commit/b49f7b38ce3895c435e4fd235bb049cc25679800))
* **api:** api update ([bbef138](https://github.com/prelude-so/go-sdk/commit/bbef1382ec8ef57bd38415e6dc8c20a8ff40bfa1))
* **api:** api update ([073ef8a](https://github.com/prelude-so/go-sdk/commit/073ef8afe12bbb5af4f486b675bb156bb67ea0b2))
* **api:** api update ([e64ba8a](https://github.com/prelude-so/go-sdk/commit/e64ba8a46d56388f41e6c44eeca52f566814a217))
* **api:** api update ([bd6ddf8](https://github.com/prelude-so/go-sdk/commit/bd6ddf8ce85df31f236bc6141224542d6a550736))
* **api:** api update ([1991e31](https://github.com/prelude-so/go-sdk/commit/1991e31e5e2c51d3a70dbab8e5d8d5bdf50a61e1))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([80ba5ab](https://github.com/prelude-so/go-sdk/commit/80ba5abad729fd74f1dd46176a58466e734d1c4f))
* **docs:** add missing pointer prefix to api.md return types ([3a21fda](https://github.com/prelude-so/go-sdk/commit/3a21fdaf80af5e7a58ba64016dd3278fe2f4d460))
* fix request delays for retrying to be more respectful of high requested delays ([92b0da0](https://github.com/prelude-so/go-sdk/commit/92b0da0e660650f78ddd95befb3fbb2a3b71c979))
* **mcp:** correct code tool API endpoint ([398430f](https://github.com/prelude-so/go-sdk/commit/398430faae2f7e1c8d14bacaef4680eb96d4ef49))
* rename param to avoid collision ([8c9b1c5](https://github.com/prelude-so/go-sdk/commit/8c9b1c50519fa0484ec368f3b9a3e619eae09e96))
* skip usage tests that don't work with Prism ([1b0e2e4](https://github.com/prelude-so/go-sdk/commit/1b0e2e450db0a9381ea70c4a4d2a35df1f55e989))


### Chores

* **ci:** add build step ([79da442](https://github.com/prelude-so/go-sdk/commit/79da442ba43029b976c9aff41d74ebe99656926a))
* **ci:** skip uploading artifacts on stainless-internal branches ([d285e3c](https://github.com/prelude-so/go-sdk/commit/d285e3c2f70c0c9eadd2a82731e67c3ef7af56a9))
* **docs:** add missing descriptions ([3955dce](https://github.com/prelude-so/go-sdk/commit/3955dce2d4a478fe5b76d4f6be567da437b7fbc4))
* elide duplicate aliases ([afd20b9](https://github.com/prelude-so/go-sdk/commit/afd20b9caa24f94b27421f0e5247cf526ce0848f))
* **internal:** codegen related update ([8a3ff52](https://github.com/prelude-so/go-sdk/commit/8a3ff526fc9e712c7f493dc4aac1feffc0ba5f4e))
* **internal:** codegen related update ([7a807da](https://github.com/prelude-so/go-sdk/commit/7a807da3d84166c83d7dc12675db3e2f732cd796))
* **internal:** minor cleanup ([17dfaa1](https://github.com/prelude-so/go-sdk/commit/17dfaa1224bd69cf60f0f4394d88934732d24dfb))
* **internal:** move custom custom `json` tags to `api` ([23d3987](https://github.com/prelude-so/go-sdk/commit/23d3987cee165fff5742a97fe40717e0802ad1e7))
* **internal:** tweak CI branches ([18ee14b](https://github.com/prelude-so/go-sdk/commit/18ee14b632b2154eb81d17e520dacb409a72d252))
* **internal:** update `actions/checkout` version ([3449b16](https://github.com/prelude-so/go-sdk/commit/3449b1648b8ce7677871c04323afbd91dccab63d))
* **internal:** update gitignore ([e0c9c8d](https://github.com/prelude-so/go-sdk/commit/e0c9c8ddcafe2275c6c3c87147619e530ab8e9c9))
* **internal:** use explicit returns ([aae93e9](https://github.com/prelude-so/go-sdk/commit/aae93e9e98ae4d8eb2d411e3a6a81fc0f2e5e099))
* **internal:** use explicit returns in more places ([47486c3](https://github.com/prelude-so/go-sdk/commit/47486c3b31e2c6be9a38445aed0a034cbb007361))
* **test:** do not count install time for mock server timeout ([f958f24](https://github.com/prelude-so/go-sdk/commit/f958f24ae1092c7ce9ea18a166915f620bc72923))
* **tests:** bump steady to v0.19.4 ([4c1123b](https://github.com/prelude-so/go-sdk/commit/4c1123b9e6f25a1bcd52dba6a491736836e713cf))
* **tests:** bump steady to v0.19.5 ([7cd2942](https://github.com/prelude-so/go-sdk/commit/7cd294298d851ee9123d91c2757c3249eb87a51c))
* update mock server docs ([1ac4fa8](https://github.com/prelude-so/go-sdk/commit/1ac4fa810469d6f02915b31c72e701994d6a2910))


### Refactors

* **tests:** switch from prism to steady ([592846f](https://github.com/prelude-so/go-sdk/commit/592846f262f5287ee5172f2c2a6cbb41b61fbb6a))

## 0.10.0 (2025-12-05)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/prelude-so/go-sdk/compare/v0.9.0...v0.10.0)

### Features

* **api:** add Notify API methods ([cb029f5](https://github.com/prelude-so/go-sdk/commit/cb029f5e07ebaab61a8b9bbac992ec2de3e5016d))

## 0.9.0 (2025-11-17)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/prelude-so/go-sdk/compare/v0.8.0...v0.9.0)

### Features

* **api:** api update ([2a464ed](https://github.com/prelude-so/go-sdk/commit/2a464ed2c8c0a98c0507ce3aed9a111e38053493))
* **api:** api update ([db62c87](https://github.com/prelude-so/go-sdk/commit/db62c879b223b0c306eb3af8c5e8170c4b3b0246))
* **api:** api update ([d178ee0](https://github.com/prelude-so/go-sdk/commit/d178ee0da7f1d4b7035803728cba9580f72a2fb5))
* **api:** api update ([7d03714](https://github.com/prelude-so/go-sdk/commit/7d03714a9b6d9e1b0bb763f6feca1a783e45fe78))
* **api:** api update ([ab44c73](https://github.com/prelude-so/go-sdk/commit/ab44c732d54e56ce52d2b9310fd8c53781084394))
* **api:** api update ([ae74d88](https://github.com/prelude-so/go-sdk/commit/ae74d88df68cccf5c6303db498834278de15f011))
* **api:** expose phone numbers management methods ([780308c](https://github.com/prelude-so/go-sdk/commit/780308c11dda1bdf7a4dc46929837d4ce38f140f))
* **api:** expose verification management methods ([852ff3b](https://github.com/prelude-so/go-sdk/commit/852ff3b4b849dab157bca92a08494e9190f3eaab))

## 0.8.0 (2025-09-25)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/prelude-so/go-sdk/compare/v0.7.0...v0.8.0)

### Features

* **api:** api update ([26a6130](https://github.com/prelude-so/go-sdk/commit/26a6130bbc2dd0c1c61f3ff39eb8d3a8fce43b1e))
* **api:** api update ([0d1e889](https://github.com/prelude-so/go-sdk/commit/0d1e88966226302c1dcec3e8dbb0a7a7b4b9aa59))
* **api:** api update ([9ad4c14](https://github.com/prelude-so/go-sdk/commit/9ad4c14447cc16828d2a2433e6984f441625cdda))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([b116360](https://github.com/prelude-so/go-sdk/commit/b11636067c7611807ddcae1405e5a31caaa3eaba))
* use slices.Concat instead of sometimes modifying r.Options ([063e46f](https://github.com/prelude-so/go-sdk/commit/063e46f0117ac38e95f39525e4d4c06fcdee122c))


### Chores

* bump minimum go version to 1.22 ([7ae34c5](https://github.com/prelude-so/go-sdk/commit/7ae34c522207fec7ad24750066c11999a99bab50))
* do not install brew dependencies in ./scripts/bootstrap by default ([0ea92dd](https://github.com/prelude-so/go-sdk/commit/0ea92ddfaad7dc28b5ace9ee6d7b44fa8899fc45))
* update more docs for 1.22 ([e4c4dd1](https://github.com/prelude-so/go-sdk/commit/e4c4dd17256f134a7e3ec5b2311b627ff57e4d58))

## 0.7.0 (2025-09-03)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/prelude-so/go-sdk/compare/v0.6.0...v0.7.0)

### Features

* **api:** api update ([63d9f36](https://github.com/prelude-so/go-sdk/commit/63d9f36a681073587b1e963e1d889d2b7fb96751))
* **api:** update via SDK Studio ([e6b3ca1](https://github.com/prelude-so/go-sdk/commit/e6b3ca1528c7726398987af37c1f079b84f31fb4))
* **client:** support optional json html escaping ([8f9d8a4](https://github.com/prelude-so/go-sdk/commit/8f9d8a472bdbd401d01f035e6b2868589cad555c))


### Bug Fixes

* **client:** process custom base url ahead of time ([3f1d80f](https://github.com/prelude-so/go-sdk/commit/3f1d80f7d1931c767188c384499e6972e00a3b78))
* close body before retrying ([c503f08](https://github.com/prelude-so/go-sdk/commit/c503f08458a701e16ac84ec7a10d09ead3acff2d))
* don't try to deserialize as json when ResponseBodyInto is []byte ([cf385aa](https://github.com/prelude-so/go-sdk/commit/cf385aae3a726932291f4c9893153e347f26d610))
* remove null from release please manifest ([6fb6926](https://github.com/prelude-so/go-sdk/commit/6fb6926185c0175d582ab4ccb289cca574c63dc7))
* use release please annotations on more places ([2b1481e](https://github.com/prelude-so/go-sdk/commit/2b1481e4f3c7684ec6e509ffae64f39805ea13fe))


### Chores

* **ci:** only run for pushes and fork pull requests ([0da7840](https://github.com/prelude-so/go-sdk/commit/0da7840eb5ef751b927a0759b13a5a972a96a036))
* **internal:** fix lint script for tests ([5412ddc](https://github.com/prelude-so/go-sdk/commit/5412ddc451370366f3995b6247dd8d5bd4e5cbf5))
* **internal:** update comment in script ([bd5b3c3](https://github.com/prelude-so/go-sdk/commit/bd5b3c383dbeee8b28a78145b105b6092829b1c3))
* **internal:** update test skipping reason ([71d160b](https://github.com/prelude-so/go-sdk/commit/71d160b5cf297b1ba48f0d5a610fd906d53f2105))
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
