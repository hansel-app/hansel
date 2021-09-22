const merge = require("webpack-merge");
const tsImportPluginFactory = require("ts-import-plugin");
const path = require("path");

module.exports = {
  // customises settings for manifest.json

  publicPath: process.env.CI ? "/hansel/" : "/",
  pwa: {
    name: "hansel",
    themeColor: "#4DBA87",
    msTileColor: "#000000",
    appleMobileWebAppCapable: "yes",
    appleMobileWebAppStatusBarStyle: "black",

    // configure the workbox plugin
    //   workboxPluginMode: 'InjectManifest',
    //   workboxOptions: {
    //     // swSrc is required in InjectManifest mode.
    //     swSrc: 'dev/sw.js',
    //     // ...other Workbox options...
    //  }
  },
  css: {
    loaderOptions: {
      less: {
        lessOptions: {
          javascriptEnabled: true,
          modifyVars: {
            // overide with less vars
            hack: `true; @import "${path.resolve(
              __dirname,
              "src",
              "index.less"
            )}";`,
          },
        },
      },
    },
  },
  chainWebpack: (config) => {
    config.module
      .rule("ts")
      .use("ts-loader")
      .tap((options) => {
        options = merge(options, {
          transpileOnly: true,
          getCustomTransformers: () => ({
            before: [
              tsImportPluginFactory({
                libraryName: "vant",
                libraryDirectory: "es",
                style: (name) => `${name}/style/less`,
              }),
            ],
          }),
          compilerOptions: {
            module: "es2015",
          },
        });
        return options;
      });
  },
};
