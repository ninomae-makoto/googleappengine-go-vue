# go vue templateについて

GAE go Vueプロジェクトのテンプレート。  
クローンした後にVisualStudioCode上ですぐに開発できるようにセットアップしている。  

## VueCLI Settings

```
? Check the features needed for your project: TS, PWA, Linter, Unit
? Use class-style component syntax? No
? Use Babel alongside TypeScript for auto-detected polyfills? No
? Pick a linter / formatter config: TSLint
? Pick additional lint features: (Press <space> to select, <a> to toggle all, <i>
 to invert selection)Lint on save
? Pick a unit testing solution: Mocha
? Where do you prefer placing config for Babel, PostCSS, ESLint, etc.? In dedicated config files
```

## VueCLIで作成したプロジェクトとの相違点

- tasks.json の追加(VSCode用)
- settings.json の追加(VSCode用)
- vue.config.js の追加(外部ライブラリの指定)
- tsconfig.json書き換え
- axios導入

tslint, tsconfigは独自ルールにしているので適宜置き換える。  


# Front setup

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

or

Command(Ctrl) + Shift + B

### Watch project
```
npm run watch
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```

### Run your unit tests
```
npm run test:unit
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


# Back Setup

## Run App

引数localでローカルかGoogleAppEngineかを判断している。

```
go run main.go local
```

http://localhost:8080/

# Deploy GoogleAppEngine

```
gcloud auth login
gcloud config set project PROJECT_ID
gcloud app deploy
gcloud app browse
```
