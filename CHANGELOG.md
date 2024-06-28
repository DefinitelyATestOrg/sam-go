# Changelog

## 3.0.0-beta.3 (2024-06-28)

Full Changelog: [v3.0.0-beta.2...v3.0.0-beta.3](https://github.com/DefinitelyATestOrg/sam-go/compare/v3.0.0-beta.2...v3.0.0-beta.3)

### Features

* **api:** update via SDK Studio ([#16](https://github.com/DefinitelyATestOrg/sam-go/issues/16)) ([b96baa3](https://github.com/DefinitelyATestOrg/sam-go/commit/b96baa334c6ca34beab0e25869a2449e277185b8))
* **api:** update via SDK Studio ([#18](https://github.com/DefinitelyATestOrg/sam-go/issues/18)) ([4726783](https://github.com/DefinitelyATestOrg/sam-go/commit/47267833719a86a605ea0bb89e76a3ebaef8d0a2))
* **api:** update via SDK Studio ([#19](https://github.com/DefinitelyATestOrg/sam-go/issues/19)) ([8a7cfaa](https://github.com/DefinitelyATestOrg/sam-go/commit/8a7cfaaf06ece510c4e8a921d4ab453cc317651b))
* **api:** update via SDK Studio ([#20](https://github.com/DefinitelyATestOrg/sam-go/issues/20)) ([b7b004b](https://github.com/DefinitelyATestOrg/sam-go/commit/b7b004bd793fee39ee88c61024c83f5ab591c0dc))
* **api:** update via SDK Studio ([#21](https://github.com/DefinitelyATestOrg/sam-go/issues/21)) ([5bbfe13](https://github.com/DefinitelyATestOrg/sam-go/commit/5bbfe1338e1b5bdf6ae58a2f0f4e2276f2957917))
* better validation of path params ([#38](https://github.com/DefinitelyATestOrg/sam-go/issues/38)) ([9f031e4](https://github.com/DefinitelyATestOrg/sam-go/commit/9f031e489c1ca6c4ed6c88f8c9debc9a578e3c05))
* propagate resource description field from stainless config to SDK docs ([#33](https://github.com/DefinitelyATestOrg/sam-go/issues/33)) ([2d2ee53](https://github.com/DefinitelyATestOrg/sam-go/commit/2d2ee53873043535f2e52fbc1f7d942d8182d722))


### Bug Fixes

* fix ExtraFields serialization / deserialization ([#42](https://github.com/DefinitelyATestOrg/sam-go/issues/42)) ([1204ddc](https://github.com/DefinitelyATestOrg/sam-go/commit/1204ddc844cb34e13ae016962281bfab04af8907))
* fix reading the error body more than once ([#36](https://github.com/DefinitelyATestOrg/sam-go/issues/36)) ([d79c936](https://github.com/DefinitelyATestOrg/sam-go/commit/d79c936c93cf15fd4d55dbc30f6c34d51eee1249))
* **internal:** fix the way that unions are deserialized in nested arrays ([#40](https://github.com/DefinitelyATestOrg/sam-go/issues/40)) ([8072ade](https://github.com/DefinitelyATestOrg/sam-go/commit/8072adec31e7d62e2991b961ca600b6f6e00fd1b))
* **test:** fix test github actions job ([#31](https://github.com/DefinitelyATestOrg/sam-go/issues/31)) ([6e79405](https://github.com/DefinitelyATestOrg/sam-go/commit/6e7940522feed9a88fb94d875645fa885bd0e561))


### Chores

* **docs:** add SECURITY.md ([#34](https://github.com/DefinitelyATestOrg/sam-go/issues/34)) ([17f6a71](https://github.com/DefinitelyATestOrg/sam-go/commit/17f6a710064a48f372e8baa1b0824d2e10b160bc))
* gitignore test server logs ([#43](https://github.com/DefinitelyATestOrg/sam-go/issues/43)) ([2c2e7ff](https://github.com/DefinitelyATestOrg/sam-go/commit/2c2e7ff6a6f3a817439f5912d9b762d24ad0d3b0))
* go live ([#22](https://github.com/DefinitelyATestOrg/sam-go/issues/22)) ([f218ba8](https://github.com/DefinitelyATestOrg/sam-go/commit/f218ba8b53a86e90bbc639b0790bd73232591c2f))
* **internal:** add scripts/test, scripts/mock and add ci job ([#30](https://github.com/DefinitelyATestOrg/sam-go/issues/30)) ([ac4d86b](https://github.com/DefinitelyATestOrg/sam-go/commit/ac4d86b0659c7c2b03e77d48ada874f3e9b74b66))
* **internal:** add slightly better logging to scripts ([#35](https://github.com/DefinitelyATestOrg/sam-go/issues/35)) ([cc08f72](https://github.com/DefinitelyATestOrg/sam-go/commit/cc08f724de0f12de78b5fe6982c434dfd00ed3b9))
* **internal:** codegen related update ([#27](https://github.com/DefinitelyATestOrg/sam-go/issues/27)) ([beb0968](https://github.com/DefinitelyATestOrg/sam-go/commit/beb09683dc6d9d499bc84e904889689848438589))
* **internal:** codegen related update ([#28](https://github.com/DefinitelyATestOrg/sam-go/issues/28)) ([892c14a](https://github.com/DefinitelyATestOrg/sam-go/commit/892c14aca5f8f5010ef781cb5c733fccf97ad8a5))
* **internal:** fix bootstrap script ([#32](https://github.com/DefinitelyATestOrg/sam-go/issues/32)) ([743d0e0](https://github.com/DefinitelyATestOrg/sam-go/commit/743d0e044802621bef5e5e924b3d37df5bf8e2a0))
* **internal:** fix format script ([#39](https://github.com/DefinitelyATestOrg/sam-go/issues/39)) ([228fdb3](https://github.com/DefinitelyATestOrg/sam-go/commit/228fdb386df9a2ba2a62f2e4ff191d1e42af56fc))
* **internal:** fix Port function for number and boolean enums ([#26](https://github.com/DefinitelyATestOrg/sam-go/issues/26)) ([1ebcf77](https://github.com/DefinitelyATestOrg/sam-go/commit/1ebcf774c08feb4392fdc043d7967443008f6787))
* **internal:** support parsing other json content types ([#37](https://github.com/DefinitelyATestOrg/sam-go/issues/37)) ([9321ac0](https://github.com/DefinitelyATestOrg/sam-go/commit/9321ac010af54e5360c6a15e9129ba9bb9cb40a0))
* **internal:** use actions/checkout@v4 for codeflow ([#25](https://github.com/DefinitelyATestOrg/sam-go/issues/25)) ([17a6ab9](https://github.com/DefinitelyATestOrg/sam-go/commit/17a6ab9039bbf980bccc9346e8fe0ddd4d997f00))
* rebuild project due to codegen change ([#29](https://github.com/DefinitelyATestOrg/sam-go/issues/29)) ([b45699d](https://github.com/DefinitelyATestOrg/sam-go/commit/b45699d5d77d6c21c43568c1af76e379640ddd9d))
* rebuild project due to codegen change ([#41](https://github.com/DefinitelyATestOrg/sam-go/issues/41)) ([29bc1dc](https://github.com/DefinitelyATestOrg/sam-go/commit/29bc1dc690232042ddbdcd24161ea3d93de922dd))


### Build System

* configure UTF-8 locale in devcontainer ([#24](https://github.com/DefinitelyATestOrg/sam-go/issues/24)) ([0d7f755](https://github.com/DefinitelyATestOrg/sam-go/commit/0d7f755e7c405b362f4eeaf7e77d86c302e17273))

## 3.0.0-beta.2 (2024-03-29)

Full Changelog: [v3.0.0-beta.1...v3.0.0-beta.2](https://github.com/DefinitelyATestOrg/sam-go/compare/v3.0.0-beta.1...v3.0.0-beta.2)

### Features

* **api:** update via SDK Studio ([#13](https://github.com/DefinitelyATestOrg/sam-go/issues/13)) ([9c3a5a7](https://github.com/DefinitelyATestOrg/sam-go/commit/9c3a5a7c328cba9187fa1095170ab41d6d4e8757))
* **api:** update via SDK Studio ([#15](https://github.com/DefinitelyATestOrg/sam-go/issues/15)) ([4eb3718](https://github.com/DefinitelyATestOrg/sam-go/commit/4eb371810bdfe699ede722e3bf4056f5bbc7a915))

## 3.0.0-beta.1 (2024-03-08)

Full Changelog: [v3.0.0-beta.0...v3.0.0-beta.1](https://github.com/DefinitelyATestOrg/sam-go/compare/v3.0.0-beta.0...v3.0.0-beta.1)

### Features

* update via SDK Studio ([#10](https://github.com/DefinitelyATestOrg/sam-go/issues/10)) ([f35afed](https://github.com/DefinitelyATestOrg/sam-go/commit/f35afed83dc2284963707e9326a218005d42950a))
* update via SDK Studio ([#12](https://github.com/DefinitelyATestOrg/sam-go/issues/12)) ([5d86b49](https://github.com/DefinitelyATestOrg/sam-go/commit/5d86b494c00440c27d2f2494d6ace297076ad16d))

## 3.0.0-beta.0 (2024-03-08)

Full Changelog: [v2.0.0-beta.0...v3.0.0-beta.0](https://github.com/DefinitelyATestOrg/sam-go/compare/v2.0.0-beta.0...v3.0.0-beta.0)

### Features

* update via SDK Studio ([bfb6a57](https://github.com/DefinitelyATestOrg/sam-go/commit/bfb6a57554cfafe4daf966e61b29e31e8e2d62b2))
* update via SDK Studio ([#8](https://github.com/DefinitelyATestOrg/sam-go/issues/8)) ([ce345ed](https://github.com/DefinitelyATestOrg/sam-go/commit/ce345ed92b7713ed21a20d607010c88555a42e84))

## 2.0.0-beta.0 (2024-03-08)

Full Changelog: [v0.1.0...v2.0.0-beta.0](https://github.com/DefinitelyATestOrg/sam-go/compare/v0.1.0...v2.0.0-beta.0)

### Features

* update via SDK Studio ([#3](https://github.com/DefinitelyATestOrg/sam-go/issues/3)) ([f0ae1ca](https://github.com/DefinitelyATestOrg/sam-go/commit/f0ae1caa3d79ceebc140b9b0eb62fb14b1b6b5ea))
* update via SDK Studio ([#5](https://github.com/DefinitelyATestOrg/sam-go/issues/5)) ([5613903](https://github.com/DefinitelyATestOrg/sam-go/commit/5613903badd9a00d6a312208775fdfc4a637e5fb))

## 0.1.0 (2024-03-08)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/DefinitelyATestOrg/sam-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** OpenAPI spec update ([a86efa3](https://github.com/DefinitelyATestOrg/sam-go/commit/a86efa3660f23952599191a592c5e3389a51d62a))
* **api:** OpenAPI spec update ([1297577](https://github.com/DefinitelyATestOrg/sam-go/commit/129757703ef98892a86464184055d4fd2518e381))
* **api:** OpenAPI spec update ([9cb8fb2](https://github.com/DefinitelyATestOrg/sam-go/commit/9cb8fb22565b4e0643a6a516d2d4cbcc6ca4f9b4))


### Chores

* configure new SDK language ([63501d4](https://github.com/DefinitelyATestOrg/sam-go/commit/63501d4841f507eb03384d0494bb70346452d02a))
* go live ([#1](https://github.com/DefinitelyATestOrg/sam-go/issues/1)) ([c901064](https://github.com/DefinitelyATestOrg/sam-go/commit/c9010641a8d2011d876caf9407d5fb6bf9939fe9))
