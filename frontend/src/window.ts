export type Env = {
  VUE_APP_SERVER_URL: string;
  VUE_APP_SERVER_PORT: number;
  VUE_APP_GOOGLE_API_KEY: string;
  VUE_APP_GOOGLE_ANALYTICS_ID: string;
};

declare global {
  interface Window {
    env: Env;
  }
}

export type WindowWithEnv = Window & {
  env: Env;
};
