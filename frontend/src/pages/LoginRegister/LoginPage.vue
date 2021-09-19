<template>
  <h1>Login</h1>
  <div class="container">
    <van-form @submit="handleLogin">
      <van-field
        v-model="username"
        name="Username"
        placeholder="Username"
        :rules="[{ required: true, message: 'Username is required' }]"
      />
      <van-field
        v-model="password"
        type="password"
        name="Password"
        placeholder="Password"
        :rules="[{ required: true, message: 'Password is required' }]"
      />
      <div class="button-container">
        <van-button round block type="primary" native-type="submit" width="80%">
          Login
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import { defineComponent, inject } from "vue";
import { LOGIN } from "@/providers/AuthProvider"
import {LOGIN_ROUTE} from "@/constants";

export default defineComponent({
  setup() {
    const login = inject(LOGIN, () => Promise.resolve(false));
    return {
      login
    };
  },
  data() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
    handleLogin() {
      this.login(this.username, this.password)
        .then(() => this.$router.push(LOGIN_ROUTE))
        .catch((err) => console.log(err, "Failed to log in"));
    }
  }
});
</script>

<style scoped lang="less">
.container .van-field {
  width: 80%;
}
.button-container {
  margin: 16px;
}
</style>
