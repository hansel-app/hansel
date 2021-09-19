<template><slot></slot></template>

<script lang="ts">
import { defineComponent, InjectionKey, provide } from "vue";
import services from "@/api/services";
import { LoginRequest } from "@/protobuf/auth_pb";

export const LOGIN: InjectionKey<(username: string, password: string) => Promise<boolean>> = Symbol("Login");

export default defineComponent({
  setup() {
    const client = services.authClient

    const login = async (username: string, password: string): Promise<boolean> => {
      const request = new LoginRequest();
      request.setUsername(username);
      request.setPassword(password);

      try {
        return await client
            .login(request)
            // If there is no error, then the authentication was successful.
            .then(() => true);
      } catch (err) {
        return Promise.reject(err)
      }
    };

    provide(LOGIN, login);
  },
});
</script>
