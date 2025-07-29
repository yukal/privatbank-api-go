## [v0.15.1](https://github.com/yukal/privatbank-api-go/compare/95c60885b905465fd46162ef7c51add2b76fddb0...e4c0f81a0d76aab650b07ebe45779b90deb9f98e) – 2025-07-29

### Refactors

- rename module  ([a8bf2258](https://github.com/yukal/privatbank-api-go/commit/a8bf2258e6f15e1faa77f6bc4755998e88ad7b14))

### Build

- use go v1.24.5 ([e4c0f81a](https://github.com/yukal/privatbank-api-go/commit/e4c0f81a0d76aab650b07ebe45779b90deb9f98e))



## [v0.15.0](https://github.com/yukal/privatbank-api-go/compare/567141b192ea05d5ea39dbc879e355fded13fb50...2ef89de46854774f537659124c6871ef81138a99) – 2025-07-28

### New Features

- **currency:**  
  - get history ([ec9fb88b](https://github.com/yukal/privatbank-api-go/commit/ec9fb88b6f3e317ce0ff79d1af7f778f8afe2927))

- **statement:**  
  - get settings ([4c49f27f](https://github.com/yukal/privatbank-api-go/commit/4c49f27f94a27eee4dae4fae1313627228457aff))
  - get transactions at ([a5efa358](https://github.com/yukal/privatbank-api-go/commit/a5efa358aaeb6ed95b9c2004556181d747d97e15))
  - get interim transactions ([49b1db57](https://github.com/yukal/privatbank-api-go/commit/49b1db57365c213c7545707b6a00f09dddb338ce))
  - get final transactions ([f550341f](https://github.com/yukal/privatbank-api-go/commit/f550341f97b1e64275e169e5e4a3d04f4382e9a3))
  - get balance at ([902b10bb](https://github.com/yukal/privatbank-api-go/commit/902b10bb470b4576b8e9842e8812005c4f5bc029)), ([a2c5fd29](https://github.com/yukal/privatbank-api-go/commit/a2c5fd2937412e99dc701bf6a446de9d51233e36))
  - get interim balance ([8e1efae3](https://github.com/yukal/privatbank-api-go/commit/8e1efae31fa3dd4a71ec6a447e03e8ef91bad1a4))
  - get final balance ([c3c6858f](https://github.com/yukal/privatbank-api-go/commit/c3c6858fa5c055d9ac773d002199df6c582ca3c0))

- **payment:**  
  - get payment info ([0ccb40e7](https://github.com/yukal/privatbank-api-go/commit/0ccb40e7903d2ab42d4883ce13d84c40a490c754))
  - get payment receipt  ([94252801](https://github.com/yukal/privatbank-api-go/commit/942528019f22f608d536bdf630b1a1774ef4b616))
  - get multiple receipts  ([109dbfd9](https://github.com/yukal/privatbank-api-go/commit/109dbfd9089217839d18acf5cc1acec973f227ed))

- **journal:**  
  - get inbox ([0afef3d8](https://github.com/yukal/privatbank-api-go/commit/0afef3d8f6bf57a73b0640b4d9483e172e08c85c))
  - get outbox ([8e1cf0fd](https://github.com/yukal/privatbank-api-go/commit/8e1cf0fd89a18f75540dc0c25f8fcc24d88ea390))
  - get paysheets ([3f65d41a](https://github.com/yukal/privatbank-api-go/commit/3f65d41ad40bbd4c94a0dae4a7d8054dc2477405))
  - get all ([5a25a139](https://github.com/yukal/privatbank-api-go/commit/5a25a13972c70028f71f52d2bdd6c9729d797877))

- **core:**
  - API instantiation ([7a3a2d61](https://github.com/yukal/privatbank-api-go/commit/7a3a2d614384301da1b210da5ff1ca4f4cfd7fc1)), ([e88a980d](https://github.com/yukal/privatbank-api-go/commit/e88a980daf41dc81a5339bacbe491d95b6a09af7))

### Refactors

- **http agent:**
  - improvements ([80a181e3](https://github.com/yukal/privatbank-api-go/commit/80a181e35291ae564c9b53e62e41a4acad9e6c5e))
  ([8e5692a2](https://github.com/yukal/privatbank-api-go/commit/8e5692a2cafcba8bbcdbd1b22e22b4d501ea0c68))
  ([6dd4401d](https://github.com/yukal/privatbank-api-go/commit/6dd4401deddef859147ece240991ca1c1d200fc7))
  ([f848efa7](https://github.com/yukal/privatbank-api-go/commit/f848efa7a216d360a144c6753776f574bff16c48))
- **other**
  - improve fetching data ([1f6a44d6](https://github.com/yukal/privatbank-api-go/commit/1f6a44d6550068c2ebcb5daf637756ea470a1261))
  - improve descriptions ([433b5d56](https://github.com/yukal/privatbank-api-go/commit/433b5d569e895a4ba68b88d1480fc16c65233515))
  - payment under construction ([795fe3f8](https://github.com/yukal/privatbank-api-go/commit/795fe3f8ba893487b89c84a30c2549a15855dbe2))
  - one instruction per page ([1330720d](https://github.com/yukal/privatbank-api-go/commit/1330720d0d1f7c98669f42574ad9441bc35d2247))
  - add comments ([e9dc3725](https://github.com/yukal/privatbank-api-go/commit/e9dc3725d0af69e423accad5c01960e98393e951))
  - fix payment info api url ([1d99633f](https://github.com/yukal/privatbank-api-go/commit/1d99633faffa5a0c553085ec0f0029659f880287))
  - api response improvement ([a7bd15be](https://github.com/yukal/privatbank-api-go/commit/a7bd15bee5c506bceb7cb14336636d996f078f34))
  - rename entities ([b7ab7521](https://github.com/yukal/privatbank-api-go/commit/b7ab752126c5ab7da236ee3fb2b65e8f036da83f))
  - set limitation ([09366a65](https://github.com/yukal/privatbank-api-go/commit/09366a657b7dd5d376749795b89c8d92c77e1c3b))
  - improve response entity names ([b26b5cc7](https://github.com/yukal/privatbank-api-go/commit/b26b5cc73ad1a30400fbe6dee6e85e4cdd3f1eda))
  - examples improvement ([2e1059ef](https://github.com/yukal/privatbank-api-go/commit/2e1059efcae7e5041a26f2ed71975a7f0a99b202))
  - api improvement ([519e0395](https://github.com/yukal/privatbank-api-go/commit/519e03955b17bc435677d99b503dd3cca3e69d6a))
  - internal code quality improvement ([13275f75](https://github.com/yukal/privatbank-api-go/commit/13275f7521c19d5b61b59606c1996233ea9e7bf7))
  - remove API knowledge ([13bfc95a](https://github.com/yukal/privatbank-api-go/commit/13bfc95a740c1b02b71522feff9520fe7d907041))
  - comb the code ([9d99aab1](https://github.com/yukal/privatbank-api-go/commit/9d99aab10674777b8322d04857a834d404620c3b))
  - use post with headers ([0b3542ae](https://github.com/yukal/privatbank-api-go/commit/0b3542ae9c21ae44d1856429f8bf3484bc07b53c))

### Chores

- add demo ([5b3d20b1](https://github.com/yukal/privatbank-api-go/commit/5b3d20b140c8ed5d2ac0e1848db2538f5a81a153))

### Docs

- **readme:**  add description ([55999dad](https://github.com/yukal/privatbank-api-go/commit/55999dad883ee7b5a148e7d429dec71d6bede83c)), ([669739e7](https://github.com/yukal/privatbank-api-go/commit/669739e7e5c20d037c94a297ede9f00496018210)), ([2ef89de4](https://github.com/yukal/privatbank-api-go/commit/2ef89de46854774f537659124c6871ef81138a99))

### Build

- init package ([31d3f6e4](https://github.com/yukal/privatbank-api-go/commit/31d3f6e4fee9cffe2d8a8acd65e1418eb25f4c8e)), ([567141b1](https://github.com/yukal/privatbank-api-go/commit/567141b192ea05d5ea39dbc879e355fded13fb50))
- use go v1.24.4 ([60248bd0](https://github.com/yukal/privatbank-api-go/commit/60248bd05ae5e4da3bd0b213650177e71f6531e9))
