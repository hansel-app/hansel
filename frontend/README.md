# Hansel Frontend

## Getting Started

1. Install `nodenv` by following the instructions [here](https://github.com/nodenv/nodenv#installation).
   Then, install Node `16.9.1` via:
   ```
   $ nodenv install 16.9.1
   ```
1. Install the project dependencies:
   ```
   $ npm install
   ```
1. Set up the necesary environment variables by editing `public/env.js` with the relevant values.
   **Make sure not to commit any credentials added here.**
   For `VUE_APP_GOOGLE_API_KEY`, you can get your own key by following the instructions [here](https://cloud.google.com/docs/authentication/api-keys#creating_an_api_key).
   Similarly, for `VUE_APP_GOOGLE_ANALYTICS_ID`, you can get your own tracking ID by following the instructions [here](https://support.google.com/analytics/answer/10269537).

## Updating Protocol Buffer Definitions

After changing the definitions of the Protocol Buffer, run `npm run proto` to regenerate the respective TypeScript / JavaScript files.
Note that `npm run serve` and `npm run build` also regenerate the TypeScript / JavaScript files automatically.

## Useful Commands

### Compiles and hot-reloads for development

```
$ npm run serve
```

### Compiles and minifies for production

```
$ npm run build
```

### Lints and fixes files

```
$ npm run lint
```

### Customize configuration

See [Configuration Reference](https://cli.vuejs.org/config/).
