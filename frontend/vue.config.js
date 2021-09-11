module.exports = {
  // customises settings for manifest.json

  // we can plonk other vue-cli plugin options here
  publicPath: process.env.CI ? "/hansel/" : "/",
  pwa: {
    name: "frontend",
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
};