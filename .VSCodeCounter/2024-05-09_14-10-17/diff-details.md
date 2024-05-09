# Diff Details

Date : 2024-05-09 14:10:17

Directory /Users/lyt/Desktop/github-repo/nbuctf_go

Total : 66 files,  9559 codes, 287 comments, 1038 blanks, all 10884 lines

[Summary](results.md) / [Details](details.md) / [Diff Summary](diff.md) / Diff Details

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [deploy/docker-compose.yaml](/deploy/docker-compose.yaml) | YAML | 20 | 1 | 0 | 21 |
| [deploy/settings.yaml](/deploy/settings.yaml) | YAML | 51 | 1 | 1 | 53 |
| [deploy/wait_service_create.sh](/deploy/wait_service_create.sh) | Shell Script | 15 | 1 | 5 | 21 |
| [nbuctfVue3/.eslintrc.cjs](/nbuctfVue3/.eslintrc.cjs) | JavaScript | 37 | 1 | 2 | 40 |
| [nbuctfVue3/.prettierrc.json](/nbuctfVue3/.prettierrc.json) | JSON | 8 | 0 | 0 | 8 |
| [nbuctfVue3/README.md](/nbuctfVue3/README.md) | Markdown | 26 | 0 | 15 | 41 |
| [nbuctfVue3/index.html](/nbuctfVue3/index.html) | HTML | 13 | 0 | 1 | 14 |
| [nbuctfVue3/jsconfig.json](/nbuctfVue3/jsconfig.json) | JSON with Comments | 8 | 0 | 1 | 9 |
| [nbuctfVue3/package.json](/nbuctfVue3/package.json) | JSON | 54 | 0 | 1 | 55 |
| [nbuctfVue3/pnpm-lock.yaml](/nbuctfVue3/pnpm-lock.yaml) | YAML | 3,024 | 0 | 457 | 3,481 |
| [nbuctfVue3/src/App.vue](/nbuctfVue3/src/App.vue) | vue | 9 | 1 | 1 | 11 |
| [nbuctfVue3/src/api/announcement.js](/nbuctfVue3/src/api/announcement.js) | JavaScript | 28 | 4 | 4 | 36 |
| [nbuctfVue3/src/api/authority.js](/nbuctfVue3/src/api/authority.js) | JavaScript | 43 | 7 | 6 | 56 |
| [nbuctfVue3/src/api/fileUploadAndDownload.js](/nbuctfVue3/src/api/fileUploadAndDownload.js) | JavaScript | 32 | 9 | 4 | 45 |
| [nbuctfVue3/src/api/game.js](/nbuctfVue3/src/api/game.js) | JavaScript | 20 | 3 | 3 | 26 |
| [nbuctfVue3/src/api/gameadmin.js](/nbuctfVue3/src/api/gameadmin.js) | JavaScript | 42 | 6 | 5 | 53 |
| [nbuctfVue3/src/api/jwt.js](/nbuctfVue3/src/api/jwt.js) | JavaScript | 7 | 1 | 1 | 9 |
| [nbuctfVue3/src/api/k8s.js](/nbuctfVue3/src/api/k8s.js) | JavaScript | 28 | 4 | 4 | 36 |
| [nbuctfVue3/src/api/menu.js](/nbuctfVue3/src/api/menu.js) | JavaScript | 7 | 2 | 1 | 10 |
| [nbuctfVue3/src/api/rank.js](/nbuctfVue3/src/api/rank.js) | JavaScript | 8 | 1 | 2 | 11 |
| [nbuctfVue3/src/api/user.js](/nbuctfVue3/src/api/user.js) | JavaScript | 91 | 13 | 8 | 112 |
| [nbuctfVue3/src/components/challenge/ChallengeCard.vue](/nbuctfVue3/src/components/challenge/ChallengeCard.vue) | vue | 338 | 6 | 19 | 363 |
| [nbuctfVue3/src/components/challenge/index.vue](/nbuctfVue3/src/components/challenge/index.vue) | vue | 64 | 0 | 5 | 69 |
| [nbuctfVue3/src/components/chart/LineChart.vue](/nbuctfVue3/src/components/chart/LineChart.vue) | vue | 124 | 0 | 7 | 131 |
| [nbuctfVue3/src/components/chart/macarons.js](/nbuctfVue3/src/components/chart/macarons.js) | JavaScript | 202 | 21 | 20 | 243 |
| [nbuctfVue3/src/components/chooseImg/index.vue](/nbuctfVue3/src/components/chooseImg/index.vue) | vue | 180 | 1 | 16 | 197 |
| [nbuctfVue3/src/components/customPic/index.vue](/nbuctfVue3/src/components/customPic/index.vue) | vue | 91 | 3 | 8 | 102 |
| [nbuctfVue3/src/components/upload/common.vue](/nbuctfVue3/src/components/upload/common.vue) | vue | 89 | 0 | 9 | 98 |
| [nbuctfVue3/src/components/upload/image.vue](/nbuctfVue3/src/components/upload/image.vue) | vue | 121 | 0 | 13 | 134 |
| [nbuctfVue3/src/core/config.js](/nbuctfVue3/src/core/config.js) | JavaScript | 5 | 1 | 2 | 8 |
| [nbuctfVue3/src/core/global.js](/nbuctfVue3/src/core/global.js) | JavaScript | 52 | 4 | 4 | 60 |
| [nbuctfVue3/src/main.js](/nbuctfVue3/src/main.js) | JavaScript | 22 | 4 | 7 | 33 |
| [nbuctfVue3/src/permission.js](/nbuctfVue3/src/permission.js) | JavaScript | 113 | 18 | 4 | 135 |
| [nbuctfVue3/src/pinia/index.js](/nbuctfVue3/src/pinia/index.js) | JavaScript | 7 | 2 | 4 | 13 |
| [nbuctfVue3/src/pinia/modules/router.js](/nbuctfVue3/src/pinia/modules/router.js) | JavaScript | 104 | 14 | 7 | 125 |
| [nbuctfVue3/src/pinia/modules/user.js](/nbuctfVue3/src/pinia/modules/user.js) | JavaScript | 162 | 18 | 11 | 191 |
| [nbuctfVue3/src/router/index.js](/nbuctfVue3/src/router/index.js) | JavaScript | 24 | 39 | 4 | 67 |
| [nbuctfVue3/src/style/main.scss](/nbuctfVue3/src/style/main.scss) | SCSS | 19 | 1 | 4 | 24 |
| [nbuctfVue3/src/utils/asyncRouter.js](/nbuctfVue3/src/utils/asyncRouter.js) | JavaScript | 23 | 2 | 3 | 28 |
| [nbuctfVue3/src/utils/bus.js](/nbuctfVue3/src/utils/bus.js) | JavaScript | 2 | 2 | 1 | 5 |
| [nbuctfVue3/src/utils/date.js](/nbuctfVue3/src/utils/date.js) | JavaScript | 33 | 6 | 2 | 41 |
| [nbuctfVue3/src/utils/downloadImg.js](/nbuctfVue3/src/utils/downloadImg.js) | JavaScript | 16 | 21 | 3 | 40 |
| [nbuctfVue3/src/utils/fmtRouterTitle.js](/nbuctfVue3/src/utils/fmtRouterTitle.js) | JavaScript | 13 | 0 | 1 | 14 |
| [nbuctfVue3/src/utils/format.js](/nbuctfVue3/src/utils/format.js) | JavaScript | 44 | 7 | 6 | 57 |
| [nbuctfVue3/src/utils/image.js](/nbuctfVue3/src/utils/image.js) | JavaScript | 99 | 10 | 13 | 122 |
| [nbuctfVue3/src/utils/page.js](/nbuctfVue3/src/utils/page.js) | JavaScript | 9 | 0 | 1 | 10 |
| [nbuctfVue3/src/utils/request.js](/nbuctfVue3/src/utils/request.js) | JavaScript | 151 | 7 | 10 | 168 |
| [nbuctfVue3/src/view/about/index.vue](/nbuctfVue3/src/view/about/index.vue) | vue | 33 | 0 | 4 | 37 |
| [nbuctfVue3/src/view/admin/game/index.vue](/nbuctfVue3/src/view/admin/game/index.vue) | vue | 647 | 2 | 32 | 681 |
| [nbuctfVue3/src/view/admin/index.vue](/nbuctfVue3/src/view/admin/index.vue) | vue | 18 | 1 | 3 | 22 |
| [nbuctfVue3/src/view/admin/user/user.vue](/nbuctfVue3/src/view/admin/user/user.vue) | vue | 521 | 1 | 30 | 552 |
| [nbuctfVue3/src/view/announcement/index.vue](/nbuctfVue3/src/view/announcement/index.vue) | vue | 199 | 0 | 12 | 211 |
| [nbuctfVue3/src/view/error/index.vue](/nbuctfVue3/src/view/error/index.vue) | vue | 24 | 0 | 4 | 28 |
| [nbuctfVue3/src/view/error/reload.vue](/nbuctfVue3/src/view/error/reload.vue) | vue | 12 | 0 | 3 | 15 |
| [nbuctfVue3/src/view/file/index.vue](/nbuctfVue3/src/view/file/index.vue) | vue | 227 | 3 | 15 | 245 |
| [nbuctfVue3/src/view/game/index.vue](/nbuctfVue3/src/view/game/index.vue) | vue | 116 | 1 | 9 | 126 |
| [nbuctfVue3/src/view/gameintroduce/index.vue](/nbuctfVue3/src/view/gameintroduce/index.vue) | vue | 292 | 10 | 20 | 322 |
| [nbuctfVue3/src/view/layout/index.vue](/nbuctfVue3/src/view/layout/index.vue) | vue | 226 | 2 | 15 | 243 |
| [nbuctfVue3/src/view/login/LoginForm.vue](/nbuctfVue3/src/view/login/LoginForm.vue) | vue | 401 | 3 | 16 | 420 |
| [nbuctfVue3/src/view/login/index.vue](/nbuctfVue3/src/view/login/index.vue) | vue | 34 | 0 | 5 | 39 |
| [nbuctfVue3/src/view/person/person.vue](/nbuctfVue3/src/view/person/person.vue) | vue | 380 | 1 | 23 | 404 |
| [nbuctfVue3/src/view/rank/index.vue](/nbuctfVue3/src/view/rank/index.vue) | vue | 232 | 3 | 13 | 248 |
| [nbuctfVue3/tailwind.config.js](/nbuctfVue3/tailwind.config.js) | JavaScript | 8 | 1 | 1 | 10 |
| [nbuctfVue3/vite.config.js](/nbuctfVue3/vite.config.js) | JavaScript | 41 | 17 | 4 | 62 |
| [nbuctfVue3/前端设计.md](/nbuctfVue3/%E5%89%8D%E7%AB%AF%E8%AE%BE%E8%AE%A1.md) | Markdown | 249 | 0 | 42 | 291 |
| [开发文档.md](/%E5%BC%80%E5%8F%91%E6%96%87%E6%A1%A3.md) | Markdown | 221 | 0 | 81 | 302 |

[Summary](results.md) / [Details](details.md) / [Diff Summary](diff.md) / Diff Details