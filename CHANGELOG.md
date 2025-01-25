# Changelog

## [1.13.1](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.13.0...v1.13.1) (2025-01-25)


### Bug Fixes

* **deps:** update hotio/lidarr docker tag to release-2.8.2.4493 ([b2cf930](https://github.com/devopsarr/terraform-provider-lidarr/commit/b2cf9308105ddccd407b5f2e7553b8843c8bbf0a))
* **deps:** update module github.com/devopsarr/lidarr-go to v1.2.0 ([660e04f](https://github.com/devopsarr/terraform-provider-lidarr/commit/660e04f8b26266db0b7478d52ebe3299aa6181a2))
* **deps:** update module github.com/stretchr/testify to v1.10.0 ([79c8c23](https://github.com/devopsarr/terraform-provider-lidarr/commit/79c8c234c3d439344bdecdef4c73c34e633e31c8))

## [1.13.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.12.0...v1.13.0) (2024-10-20)


### Features

* add support for log size limit ([ca54021](https://github.com/devopsarr/terraform-provider-lidarr/commit/ca54021472f06e17ccbc5fa94e10b5cd7cae1a70))


### Bug Fixes

* artist test ([833ee11](https://github.com/devopsarr/terraform-provider-lidarr/commit/833ee118acb5bb7e1ad6867f8bc6c48af232698c))
* bump golangci-lint version ([67e1cb8](https://github.com/devopsarr/terraform-provider-lidarr/commit/67e1cb8084f9f9378a32857c25c33fa032439e4d))
* **deps:** update hotio/lidarr docker tag to release-2.6.4.4402 ([e38c767](https://github.com/devopsarr/terraform-provider-lidarr/commit/e38c76763349d024c79e2ff0803a13df6df4bba5))
* **deps:** update module github.com/devopsarr/lidarr-go to v1.1.1 ([a8de68c](https://github.com/devopsarr/terraform-provider-lidarr/commit/a8de68cb1797b5fc289c5c9768da99d2fa356a71))
* **deps:** update terraform-framework ([545446f](https://github.com/devopsarr/terraform-provider-lidarr/commit/545446fcb3c93146f07a9312cb4cea276f01a564))
* gorreleaser syntax ([5a49878](https://github.com/devopsarr/terraform-provider-lidarr/commit/5a4987895084ad95d4c014dff51fb9dbbc62f297))
* use int instead of str to update res ([d427d34](https://github.com/devopsarr/terraform-provider-lidarr/commit/d427d3409926debaf41655369dc47788792de367))

## [1.12.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.11.0...v1.12.0) (2024-05-05)


### Features

* align with sdk v1.x and correct lint ([735ad86](https://github.com/devopsarr/terraform-provider-lidarr/commit/735ad86a4de63837a2afe54e452d1f0ce3b0b331))
* move to context based authentication and add extra headers ([c72c921](https://github.com/devopsarr/terraform-provider-lidarr/commit/c72c921f7703753260c541fac900dc64da43ad15))

## [1.11.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.10.0...v1.11.0) (2023-12-07)


### Features

* **#202:** manage v2 sensitive values ([3a326f9](https://github.com/devopsarr/terraform-provider-lidarr/commit/3a326f91eacd785082cb4b9c4192a096c074b9fb))
* add host required password confirmation ([28067fa](https://github.com/devopsarr/terraform-provider-lidarr/commit/28067faf6f721df0734ec692a783ae6faa07cdd4))
* remove obsolete boxcar notification ([b78ff69](https://github.com/devopsarr/terraform-provider-lidarr/commit/b78ff691f4ec83f1d861618c2797b7267a8f58be))
* update api methods ([a3bd437](https://github.com/devopsarr/terraform-provider-lidarr/commit/a3bd4374c0d98f2289862b199d0d6e99bbd7d4c0))
* update to go 1.21 ([c8c7cc8](https://github.com/devopsarr/terraform-provider-lidarr/commit/c8c7cc8cf88ad8fdc24c090c0763f18c540a3937))

## [1.10.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.9.0...v1.10.0) (2023-10-12)


### Features

* add missing use only ID for delete ([b9ff5d6](https://github.com/devopsarr/terraform-provider-lidarr/commit/b9ff5d6a8adf4c1087511016531ae1d072142f86))
* **devopsarr/terraform-provider-radarr#203:** add host data source ([5421eeb](https://github.com/devopsarr/terraform-provider-lidarr/commit/5421eeb949f16afd3019e268d5e259428a912822))
* **devopsarr/terraform-provider-radarr#203:** add host resource ([a8abc70](https://github.com/devopsarr/terraform-provider-lidarr/commit/a8abc70d55934b1d8d6eb719ff4f5f586cf5684d))
* get back system status data source ([6b65de0](https://github.com/devopsarr/terraform-provider-lidarr/commit/6b65de0c630da4583edca0caa35588998ef31a0c))
* improfe diagnostics and remove useless logic for data list and delete methods ([44c16ab](https://github.com/devopsarr/terraform-provider-lidarr/commit/44c16ab58ff8b228cb1d94aee073d0505e78cf99))
* remove obsolete indexer rarbg ([3bd6db0](https://github.com/devopsarr/terraform-provider-lidarr/commit/3bd6db08e6180435ba48aabe0e9f6413425a3487))


### Bug Fixes

* quality profile use all formats and ordered quality groups ([b4d4a94](https://github.com/devopsarr/terraform-provider-lidarr/commit/b4d4a947e3dfa84c2feccb0e7ef05bee9fee2987))
* remove useless type conversion ([ff23fec](https://github.com/devopsarr/terraform-provider-lidarr/commit/ff23fec5ce1a58e6209345cb6ae14c1ca499e0b9))

## [1.9.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.8.0...v1.9.0) (2023-05-29)


### Features

* add notification signal resource ([b892260](https://github.com/devopsarr/terraform-provider-lidarr/commit/b892260ab5c629c6d3ad951701837c21ef76a027))
* align music fields with lidarr naming convention ([0bfe5b8](https://github.com/devopsarr/terraform-provider-lidarr/commit/0bfe5b854066761b5b357f854b9f2b27091fa138))

## [1.8.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.7.0...v1.8.0) (2023-05-23)


### Features

* add apprise notification ([1342e8d](https://github.com/devopsarr/terraform-provider-lidarr/commit/1342e8d8bc2776c9bea00f51dd7abec74d218601))
* add notification ntfy resource ([3597862](https://github.com/devopsarr/terraform-provider-lidarr/commit/35978629e9af27a419437339deea61af907faec4))
* add notification on delete flags ([4e886ad](https://github.com/devopsarr/terraform-provider-lidarr/commit/4e886ad05d97f97a7527b9cde9537984ccd11b3e))
* add notification simplepush ([0902f45](https://github.com/devopsarr/terraform-provider-lidarr/commit/0902f459c954aefcb6ab1f4aa2d057c7068df7c5))
* add on health restored notification flag ([fbc398a](https://github.com/devopsarr/terraform-provider-lidarr/commit/fbc398a623bc3c0909418cb82a963a0a2db9be6d))
* make notification flags optional ([aa3af7c](https://github.com/devopsarr/terraform-provider-lidarr/commit/aa3af7c146f2a926b4dbf91cdfc25dff950691fc))


### Bug Fixes

* remove passkey parameter from redacted indexer ([fc50e57](https://github.com/devopsarr/terraform-provider-lidarr/commit/fc50e57b5dbdcc61021d4921bb11c4e479b668dc))

## [1.7.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.6.1...v1.7.0) (2023-03-28)


### Features

* add custom format condition data source ([536d115](https://github.com/devopsarr/terraform-provider-lidarr/commit/536d115e8dc238921cb784e24bba3fe39a8aa195))
* add custom format condition release group data source ([adc2c54](https://github.com/devopsarr/terraform-provider-lidarr/commit/adc2c54b688b82f2f784f0505581a392990ee8d4))
* add custom format condition release title data source ([5d893a1](https://github.com/devopsarr/terraform-provider-lidarr/commit/5d893a13600b8d570d2c2831855d256908c558b1))
* add custom format condition size data source ([5496e5d](https://github.com/devopsarr/terraform-provider-lidarr/commit/5496e5d24f02604cc78b744bef2f68615af25f73))
* add custom format data source ([08fdf84](https://github.com/devopsarr/terraform-provider-lidarr/commit/08fdf84597607e7430664c48d734237dc1feaf76))
* add custom format parameters to quality profile ([021cc3f](https://github.com/devopsarr/terraform-provider-lidarr/commit/021cc3fad2ae1d3e97944cb7c572b7e2a137c6ee))
* add custom format resource ([25fce23](https://github.com/devopsarr/terraform-provider-lidarr/commit/25fce230424504758677466287d255764156be8d))
* add custom formats data source ([8a08049](https://github.com/devopsarr/terraform-provider-lidarr/commit/8a0804971b414ce6e05fba9b21c6a315523722c3))
* align release profile to v1.1.3 field changes ([5ed8049](https://github.com/devopsarr/terraform-provider-lidarr/commit/5ed80491bdbec716853f30715653f1d56f726f7e))

## [1.6.1](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.6.0...v1.6.1) (2023-03-10)


### Bug Fixes

* download client write implementation ([7db6c95](https://github.com/devopsarr/terraform-provider-lidarr/commit/7db6c9541b15a142ca9fa5a1a4ccbfb9c225115b))

## [1.6.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.5.1...v1.6.0) (2023-03-07)


### Features

* add artist data source ([f7d5b5b](https://github.com/devopsarr/terraform-provider-lidarr/commit/f7d5b5b9adb12e7cb047fe218374f07c942f0719))
* add artist resource ([5c788e9](https://github.com/devopsarr/terraform-provider-lidarr/commit/5c788e975047d27410f91ebafacce575295bfca9))
* add artists data source ([3203a9d](https://github.com/devopsarr/terraform-provider-lidarr/commit/3203a9d8639b55f11b9f24898d2fde2fb8610ddf))
* add delay profile data source ([47c9dfe](https://github.com/devopsarr/terraform-provider-lidarr/commit/47c9dfe4c3008348c39c11c7500eb93bf641d3fd))
* add delay profile resource ([125a2d2](https://github.com/devopsarr/terraform-provider-lidarr/commit/125a2d26a5cc0dc80ce7c07a2669aea530ceefc3))
* add delay profiles data source ([d9f1a24](https://github.com/devopsarr/terraform-provider-lidarr/commit/d9f1a24983713298478a2b6fca7e66baa78e1a39))
* add metadata config data source ([f4303ed](https://github.com/devopsarr/terraform-provider-lidarr/commit/f4303ed8215f437281f002485ee5e276bcb8c004))
* add metadata config resource ([396108c](https://github.com/devopsarr/terraform-provider-lidarr/commit/396108c4688b84f4913cd6aaba62cb799051eeca))
* add metadata consumers data source ([80bc78f](https://github.com/devopsarr/terraform-provider-lidarr/commit/80bc78fca839e4da847acee6c4ed5d8292d81c44))
* add metadata data source ([2f08de8](https://github.com/devopsarr/terraform-provider-lidarr/commit/2f08de8c6a78fe006efc3e68f7c9ea9de022c952))
* add metadata kodi resource ([359940a](https://github.com/devopsarr/terraform-provider-lidarr/commit/359940a79f0e14b9371933a6104fad01c1aa271c))
* add metadata profile data source ([e38c826](https://github.com/devopsarr/terraform-provider-lidarr/commit/e38c826054a87140721ef5363002932543a95cb2))
* add metadata profile resource ([e71abfa](https://github.com/devopsarr/terraform-provider-lidarr/commit/e71abfa5cb90dfef394799d5f2e50f319afa8404))
* add metadata profiles data source ([3a17665](https://github.com/devopsarr/terraform-provider-lidarr/commit/3a1766519b996f8687f047dc5317fdbd61725f98))
* add metadata resource ([a13d275](https://github.com/devopsarr/terraform-provider-lidarr/commit/a13d27520965edd0ca12cb05400830a6313440d9))
* add metadata roksbox resource ([c1dbeaf](https://github.com/devopsarr/terraform-provider-lidarr/commit/c1dbeaffb8e0ad63093fd5b59a121481270a65f4))
* add metadata wdtv resource ([3626105](https://github.com/devopsarr/terraform-provider-lidarr/commit/3626105e30cd113282bfe2c3e27766c27955a552))
* add primary album type data source ([71369e7](https://github.com/devopsarr/terraform-provider-lidarr/commit/71369e7214e70ae30a4c899fff627759ac9a3095))
* add primary album types data source ([8e8234d](https://github.com/devopsarr/terraform-provider-lidarr/commit/8e8234d60ee670162e13bfabdf9ecd508d56333b))
* add quality data source ([602c369](https://github.com/devopsarr/terraform-provider-lidarr/commit/602c369cdd8c8cd81a54a0f997d90953be2e7c01))
* add quality definition data source ([1bfd2b8](https://github.com/devopsarr/terraform-provider-lidarr/commit/1bfd2b848a0440404f9ed56f1af6f8b5c077640d))
* add quality definition resource ([3b7f8cb](https://github.com/devopsarr/terraform-provider-lidarr/commit/3b7f8cb60d531dea6abdbec38cff042e25279ef3))
* add quality definitions data source ([f4a5cef](https://github.com/devopsarr/terraform-provider-lidarr/commit/f4a5cef716e30b109db9d9ef97e93d7ab8a98972))
* add quality profile data source ([9e5db8d](https://github.com/devopsarr/terraform-provider-lidarr/commit/9e5db8d5ddbba6e14090cade08db7ab3a2988f83))
* add quality profile resource ([2453acf](https://github.com/devopsarr/terraform-provider-lidarr/commit/2453acf0fbc3303edc37a89904872ac1997b043a))
* add quality profiles data source ([3607d01](https://github.com/devopsarr/terraform-provider-lidarr/commit/3607d01e161de1939f9ce7ad308805b445370f85))
* add release profile data source ([b725103](https://github.com/devopsarr/terraform-provider-lidarr/commit/b725103ebf2ed8e4a2600cbe84bfc4ad56fc51da))
* add release profile resource ([47eeb9e](https://github.com/devopsarr/terraform-provider-lidarr/commit/47eeb9e473a67febc6d78f852cfc9620fa162899))
* add release profiles data source ([c939bce](https://github.com/devopsarr/terraform-provider-lidarr/commit/c939bcecc018334c2bed80178d02d76b24c0aa16))
* add release status data source ([b2c75a3](https://github.com/devopsarr/terraform-provider-lidarr/commit/b2c75a30b53dd0faf38c5f64ee600fda4e799819))
* add release statuses data source ([813f67a](https://github.com/devopsarr/terraform-provider-lidarr/commit/813f67a395d6b9c37c0e58d83c80442469b625f2))
* add secondary album type data source ([19fd445](https://github.com/devopsarr/terraform-provider-lidarr/commit/19fd44549e3d0887f4d98b01dbd4ff71b2f83643))
* add secondary album types data source ([5657ce6](https://github.com/devopsarr/terraform-provider-lidarr/commit/5657ce69dc0b32a6de2d11f2f6f29fe93b02b7d1))

## [1.5.1](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.5.0...v1.5.1) (2023-02-28)


### Bug Fixes

* data source read from config instead of state ([95b3fd7](https://github.com/devopsarr/terraform-provider-lidarr/commit/95b3fd7463094db838baaf9515f4aefc1447735b))
* manage not working tv related api fields on download clients ([2a32dc6](https://github.com/devopsarr/terraform-provider-lidarr/commit/2a32dc6d76c07c759bd2375aad9912a64b91550f))
* notification discografy seed time ([62232c2](https://github.com/devopsarr/terraform-provider-lidarr/commit/62232c2a4dba1ac4372c51119ae01c87fb7e1efd))
* notification ssl flag ([0d41d42](https://github.com/devopsarr/terraform-provider-lidarr/commit/0d41d429f793bf7e48ead611b0acc9227cd64877))
* notification subsonic wrong implementation ([c8c9e91](https://github.com/devopsarr/terraform-provider-lidarr/commit/c8c9e914c8bc10877512db8c90a1d86301dfc9c7))
* pass import list type ([bc208ea](https://github.com/devopsarr/terraform-provider-lidarr/commit/bc208eabddfb8b7061a7ad222e4ccce278e8d6f2))

## [1.5.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.4.0...v1.5.0) (2023-01-25)


### Features

* add import list data source ([09a1d7d](https://github.com/devopsarr/terraform-provider-lidarr/commit/09a1d7d50cddca4a539a91c927819577dd7f9c72))
* add import list exclusion datasource ([a2cf86f](https://github.com/devopsarr/terraform-provider-lidarr/commit/a2cf86ff856acaf336965e3b6b16630f30f073c2))
* add import list exclusion resource ([2c02efc](https://github.com/devopsarr/terraform-provider-lidarr/commit/2c02efc64315c751db38adf0f6974ccbb85e0605))
* add import list exclusions datasource ([e36dd93](https://github.com/devopsarr/terraform-provider-lidarr/commit/e36dd937a8bfcd5565ab09e4b122beb2ace7dba7))
* add import list headphones resource ([7bb04d8](https://github.com/devopsarr/terraform-provider-lidarr/commit/7bb04d8e0b7ff6de081791fb3ea11f62cefa22aa))
* add import list lastfm tag resource ([ae380a2](https://github.com/devopsarr/terraform-provider-lidarr/commit/ae380a2a02390c59b834a9aab13ea9bec7d61c60))
* add import list lastfm user resource ([9909cc3](https://github.com/devopsarr/terraform-provider-lidarr/commit/9909cc305cf8f2576bc4bf7e6d736e5c8366401f))
* add import list lidarr list resource ([b977717](https://github.com/devopsarr/terraform-provider-lidarr/commit/b977717f1e79babd364938c045a8269b37b68520))
* add import list lidarr resource ([114e553](https://github.com/devopsarr/terraform-provider-lidarr/commit/114e55315ee48177965c301816a93c9b29086900))
* add import list music brainz resource ([d669766](https://github.com/devopsarr/terraform-provider-lidarr/commit/d6697668279c3b0c7755818a5559254139ea4f10))
* add import list resource ([a824e3a](https://github.com/devopsarr/terraform-provider-lidarr/commit/a824e3ab1dfe58b2c1f6ea221cca7b90984e1158))
* add import list spotify albums resource ([0f8043c](https://github.com/devopsarr/terraform-provider-lidarr/commit/0f8043cf4a5e4557ada07ce1203c11c49d2bfde1))
* add import list spotify artists resource ([88f3bb6](https://github.com/devopsarr/terraform-provider-lidarr/commit/88f3bb6577ae2331dd5ae77dac9f42f643fe656f))
* add import list spotify playlists resource ([0954ee2](https://github.com/devopsarr/terraform-provider-lidarr/commit/0954ee2dc0137f2cc37105e5bf3b61ebfdbb6969))
* add import lists data source ([d78d705](https://github.com/devopsarr/terraform-provider-lidarr/commit/d78d7056a832a207c40995168be7c1def2e8962d))
* add media management data source ([6bf2611](https://github.com/devopsarr/terraform-provider-lidarr/commit/6bf2611cf89f5706c732242d752a649cb17c0732))
* add media management resource ([bc8929c](https://github.com/devopsarr/terraform-provider-lidarr/commit/bc8929c7190f7342fda35c8b5c80237adf5ecb20))
* add naming data source ([13ad9df](https://github.com/devopsarr/terraform-provider-lidarr/commit/13ad9df8a4521957d860beb3b5692e9721c0e823))
* add naming resource ([cc2ebc6](https://github.com/devopsarr/terraform-provider-lidarr/commit/cc2ebc69a17df93b1072f97b7e839c27c56893a6))
* add root folder data source ([bddeaa0](https://github.com/devopsarr/terraform-provider-lidarr/commit/bddeaa06f8053a6495eca7c33069a56c989b8eb8))
* add root folder resource ([7e6087a](https://github.com/devopsarr/terraform-provider-lidarr/commit/7e6087aba687bef5eb2773d7e36faec4fc59aac6))
* add root folders data source ([076f632](https://github.com/devopsarr/terraform-provider-lidarr/commit/076f632dbd0c34af8dbfe546f9f51ab0174d8c6a))
* improve tags slice allocation ([5e1c304](https://github.com/devopsarr/terraform-provider-lidarr/commit/5e1c304309e08eada3080bf9dcfcd1c716660695))


### Bug Fixes

* remove deprecated indexers ([1dd4d40](https://github.com/devopsarr/terraform-provider-lidarr/commit/1dd4d401e42ff775e009678a3812ce97094734e6))
* rootfolder should have metadata profile id ([f2caa14](https://github.com/devopsarr/terraform-provider-lidarr/commit/f2caa1410ff9a19939f22c30fafcce5c5efe8721))

## [1.4.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.3.0...v1.4.0) (2022-12-15)


### Features

* add download client aria2 ([e76a59f](https://github.com/devopsarr/terraform-provider-lidarr/commit/e76a59f35c4e295d9bd10705bba2d7238bec5566))
* add download client deluge ([8d7ab65](https://github.com/devopsarr/terraform-provider-lidarr/commit/8d7ab65f6ffead8156c692c23cbe908bf768b9cd))
* add download client flood ([98a20b2](https://github.com/devopsarr/terraform-provider-lidarr/commit/98a20b2583736be02cf0f1d8b440ced13c8ad128))
* add download client hadouken ([b9dcaf8](https://github.com/devopsarr/terraform-provider-lidarr/commit/b9dcaf82d7e1f2951356a6768f35eaff2b35758c))
* add download client nzbget ([06c4d72](https://github.com/devopsarr/terraform-provider-lidarr/commit/06c4d72090599f1bed351c1fbe023c5dea3c5bfb))
* add download client nzbvortex ([de8022e](https://github.com/devopsarr/terraform-provider-lidarr/commit/de8022eb8a6c4872f0b5824472c2cf1fb307d0bb))
* add download client pneumatic ([7f1ea9e](https://github.com/devopsarr/terraform-provider-lidarr/commit/7f1ea9e8c6d0fceb6daf966821b64ce6d2da3554))
* add download client qbittorrent ([65f3894](https://github.com/devopsarr/terraform-provider-lidarr/commit/65f389432dae590729f8276614f40cf58b5fac30))
* add download client rtorrent ([3d523f1](https://github.com/devopsarr/terraform-provider-lidarr/commit/3d523f1740e20252447d0a1de2da25945ef876b5))
* add download client sabnzbd ([7fcf89d](https://github.com/devopsarr/terraform-provider-lidarr/commit/7fcf89d40474d93e1056c4676d546db04d56be25))
* add download client torrent blackhole ([516dbf5](https://github.com/devopsarr/terraform-provider-lidarr/commit/516dbf5fc40d35eac77287c3e5d1c443fc28aee9))
* add download client torrent download station ([8c19a82](https://github.com/devopsarr/terraform-provider-lidarr/commit/8c19a8221a79b57a2d639ea76dc16f899190024d))
* add download client usenet blackhole ([0f4413e](https://github.com/devopsarr/terraform-provider-lidarr/commit/0f4413ed66501579deb353cb6f51e394173024fb))
* add download client usenet download station ([f685604](https://github.com/devopsarr/terraform-provider-lidarr/commit/f685604890e359142b3362caca79400d05220223))
* add download client utorrent ([f729f5f](https://github.com/devopsarr/terraform-provider-lidarr/commit/f729f5f91255679fe6002db4a54851cf45d8e881))
* add download client vuze ([4b85640](https://github.com/devopsarr/terraform-provider-lidarr/commit/4b856409af0082af9e1d28e4f5806d758f685761))
* add indexer config datasource ([001ef8a](https://github.com/devopsarr/terraform-provider-lidarr/commit/001ef8aa2d97f20563f65a0165f5faa7b06f6c45))
* add indexer config resource ([f2d89d2](https://github.com/devopsarr/terraform-provider-lidarr/commit/f2d89d26f99d86af10f366c0d3a5d23b6e46d5bb))
* add indexer datasource ([9805152](https://github.com/devopsarr/terraform-provider-lidarr/commit/9805152b917f8dc074b3b92ab6b111c5b5bcfd5b))
* add indexer filelist ([3e7456d](https://github.com/devopsarr/terraform-provider-lidarr/commit/3e7456decfae1cff1033470b5d2d209482534b46))
* add indexer gazelle ([34b78ae](https://github.com/devopsarr/terraform-provider-lidarr/commit/34b78ae7234addb3c5adbc365230494f29da3bbb))
* add indexer headphones ([08132e7](https://github.com/devopsarr/terraform-provider-lidarr/commit/08132e7b209b72d44f1fe3b6aded16329511ec40))
* add indexer iptorrents ([412379b](https://github.com/devopsarr/terraform-provider-lidarr/commit/412379bdec1c12a349342173d5aa1de199e61733))
* add indexer newznab resource ([c514269](https://github.com/devopsarr/terraform-provider-lidarr/commit/c5142696862fb3616fbdd3dc6217889d641e5cba))
* add indexer nyaa ([cc24529](https://github.com/devopsarr/terraform-provider-lidarr/commit/cc2452978391d5a6b06c460ef4015b2655fcdae1))
* add indexer omgwtfnzbs ([bd81590](https://github.com/devopsarr/terraform-provider-lidarr/commit/bd8159083a253e0c77f534c98c7313d947ef7dc5))
* add indexer rarbg resource ([0beef0d](https://github.com/devopsarr/terraform-provider-lidarr/commit/0beef0d09b7475969eba5084c796de8ec8f743a0))
* add indexer redacted ([6a68aff](https://github.com/devopsarr/terraform-provider-lidarr/commit/6a68aff2363688425da462a415df28c610eae041))
* add indexer resource ([96cdca5](https://github.com/devopsarr/terraform-provider-lidarr/commit/96cdca5bdeaee7868f9cffebcacbcfa00653a56d))
* add indexer torrent rss ([6321c9d](https://github.com/devopsarr/terraform-provider-lidarr/commit/6321c9d2c3de52d9a00426f3c72c1db6ba1cda7e))
* add indexer torrentleech ([d80deaf](https://github.com/devopsarr/terraform-provider-lidarr/commit/d80deaf48acabd762925f1bb0ae2dedf6ba8a390))
* add indexer torznab ([b97ffcd](https://github.com/devopsarr/terraform-provider-lidarr/commit/b97ffcd6432a37eb461e3a6ea5f8848933f07b16))
* add indexer waffles ([fd663a1](https://github.com/devopsarr/terraform-provider-lidarr/commit/fd663a1ca34b35c341698fc6b0a36a151cbf02d1))
* add indexers datasource ([331b347](https://github.com/devopsarr/terraform-provider-lidarr/commit/331b3474956022d6042efef3a1449a6e95fb9b5c))
* add notification boxcar ([8e04554](https://github.com/devopsarr/terraform-provider-lidarr/commit/8e0455424d52914f5f98fba90281690deee60f12))
* add notification discord ([dbfad1c](https://github.com/devopsarr/terraform-provider-lidarr/commit/dbfad1c9ea60876f7d68685bb7e435b599211c59))
* add notification email ([a8016b1](https://github.com/devopsarr/terraform-provider-lidarr/commit/a8016b17ea4f094da6623bb6738e7c1fc96a41ee))
* add notification emby ([aa093b5](https://github.com/devopsarr/terraform-provider-lidarr/commit/aa093b5692984f63f1f88cb67521c7e0da541062))
* add notification gotify ([6cf79ab](https://github.com/devopsarr/terraform-provider-lidarr/commit/6cf79ab937068f92a6d7533da184c35b8ba13283))
* add notification join ([65d3cd8](https://github.com/devopsarr/terraform-provider-lidarr/commit/65d3cd844c0e502aa2d2e0f1e72f606ed6270741))
* add notification kodi ([d5ce452](https://github.com/devopsarr/terraform-provider-lidarr/commit/d5ce452d5cd8dd3c90fb2841c8674ac2b51b793f))
* add notification mailgun ([53ca8b4](https://github.com/devopsarr/terraform-provider-lidarr/commit/53ca8b4b79765052f452f175e2e6ef6c10c84a06))
* add notification notifiarr ([32f9a73](https://github.com/devopsarr/terraform-provider-lidarr/commit/32f9a73b41623e912bf528a21f2b0b01339d2794))
* add notification plex ([980f42b](https://github.com/devopsarr/terraform-provider-lidarr/commit/980f42b506dff13f7b17901f6333d86b0896935c))
* add notification prowl ([db26097](https://github.com/devopsarr/terraform-provider-lidarr/commit/db2609700f341e41957dc0eecc0654ba7ce95295))
* add notification pushbullet ([349ebc6](https://github.com/devopsarr/terraform-provider-lidarr/commit/349ebc6f19c10206eed54e26c9820eb64364f24a))
* add notification pushover ([612ecea](https://github.com/devopsarr/terraform-provider-lidarr/commit/612ecea6693abb9bd5db60f41a83e47be9b08ca7))
* add notification sendgrid ([8defd11](https://github.com/devopsarr/terraform-provider-lidarr/commit/8defd11703c4d222ded1c4918c8c7e1174d346c5))
* add notification slack ([df73986](https://github.com/devopsarr/terraform-provider-lidarr/commit/df73986d8a55babfa28b58f05ebef07f013fb5cb))
* add notification subsonic ([702136b](https://github.com/devopsarr/terraform-provider-lidarr/commit/702136b3c5b18845f4e3535423c73a2138ee209e))
* add notification synology indexer ([7742c46](https://github.com/devopsarr/terraform-provider-lidarr/commit/7742c46cb4ef339d2ebab694ec85ebea77976794))
* add notification telegram ([45f5727](https://github.com/devopsarr/terraform-provider-lidarr/commit/45f57272d8da43a1f38acaf97773374e92a2cafe))
* add notification twitter ([54b8b82](https://github.com/devopsarr/terraform-provider-lidarr/commit/54b8b82f219c1f9ad63a70509f9e2f3651e480ff))


### Bug Fixes

* few download client fields ([e9b2c6b](https://github.com/devopsarr/terraform-provider-lidarr/commit/e9b2c6bc5de6ed06a3b37296a52a673092339c43))
* fix few field type ([a492097](https://github.com/devopsarr/terraform-provider-lidarr/commit/a49209781c6dea7c572902f9e827ff8e182f5132))
* fix indexer data source fields ([de6254f](https://github.com/devopsarr/terraform-provider-lidarr/commit/de6254f76d465a227a337d1c88fa171cdd6694cd))
* watch folder type ([9dc1063](https://github.com/devopsarr/terraform-provider-lidarr/commit/9dc106317b5882befa1bda3f45ec9c7bbda26d5b))

## [1.3.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.2.0...v1.3.0) (2022-11-17)


### Features

* add download client config datasource ([300be92](https://github.com/devopsarr/terraform-provider-lidarr/commit/300be924f84d0d8860797b212961f68a5dd240ef))
* add download client config resource ([1829d41](https://github.com/devopsarr/terraform-provider-lidarr/commit/1829d41b8fa4c76c4cb13fe6beed2abe7aa38bf4))
* add download client datasource ([3c5af07](https://github.com/devopsarr/terraform-provider-lidarr/commit/3c5af0782557f99dd54c3b1884fd7f19ec87bb61))
* add download client resource ([eaacf6a](https://github.com/devopsarr/terraform-provider-lidarr/commit/eaacf6aad1094685a15216d489598f445019850d))
* add download client transmission resource ([0197739](https://github.com/devopsarr/terraform-provider-lidarr/commit/0197739b4633419c15c03a82166068e8802a85aa))
* add download clients datasource ([2198afc](https://github.com/devopsarr/terraform-provider-lidarr/commit/2198afc4a55dc332042df35509110f5b296d287f))
* add remote path mapping datasource ([b71473f](https://github.com/devopsarr/terraform-provider-lidarr/commit/b71473ff357478e0b1ed788e82dad8fb2d7c7b84))
* add remote path mapping resource ([0667d4e](https://github.com/devopsarr/terraform-provider-lidarr/commit/0667d4e2c62d1a473991437ee393d73a2db59100))
* add remote path mappings datasource ([c0772c5](https://github.com/devopsarr/terraform-provider-lidarr/commit/c0772c5049195df4f46c1d9bc292260e48455d5d))

## [1.2.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.1.0...v1.2.0) (2022-11-15)


### Features

* add notification custom script resource ([f2871a1](https://github.com/devopsarr/terraform-provider-lidarr/commit/f2871a11a44274bec8609d39c61f1403a167a22c))
* add notification datasource ([a34ef76](https://github.com/devopsarr/terraform-provider-lidarr/commit/a34ef761cee3410c444439a77735150d081cfc2b))
* add notification resource ([814ef1a](https://github.com/devopsarr/terraform-provider-lidarr/commit/814ef1a0edb58966ffcb9d8538d86eb401d8ca2e))
* add notification webhook resource ([1b9b7a5](https://github.com/devopsarr/terraform-provider-lidarr/commit/1b9b7a59cbfa6651e118e0a5ef1cc499b551eede))
* add notifications data source ([eb5430f](https://github.com/devopsarr/terraform-provider-lidarr/commit/eb5430f01cbd18097f1aa47a1df8b08fc985f217))
* add system status datasource ([e73039f](https://github.com/devopsarr/terraform-provider-lidarr/commit/e73039fa5489ecca7d4690bfb26f4842f3275033))
* add tag datasource ([7f5c79c](https://github.com/devopsarr/terraform-provider-lidarr/commit/7f5c79c048a5afa81dca684d6ab026b5c9392f61))


### Bug Fixes

* system status mapping ([931802b](https://github.com/devopsarr/terraform-provider-lidarr/commit/931802b149c2c5453acc4cd0af273179439113cd))

## [1.1.0](https://github.com/devopsarr/terraform-provider-lidarr/compare/v1.0.0...v1.1.0) (2022-08-29)


### Features

* add validators ([0b9896a](https://github.com/devopsarr/terraform-provider-lidarr/commit/0b9896afe948e78db56f45083a09e68965b46955))


### Bug Fixes

* remove set parameter for framework 0.9.0 ([1adf511](https://github.com/devopsarr/terraform-provider-lidarr/commit/1adf511a0a5f73ed893eac38bf65ce8587f74e41))
* repo reference ([ce82856](https://github.com/devopsarr/terraform-provider-lidarr/commit/ce82856cef04b8834a5eb1ad316546b20d69e7e5))

## 1.0.0 (2022-03-15)


### Features

* first configuration ([c927928](https://github.com/devopsarr/terraform-provider-lidarr/commit/c927928b5bedb64510611e468432ca23cfaae68a))
