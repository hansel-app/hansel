<template>
  <!-- TODO: some onboarding screens for user to set displayName and avatar? -->
  <h1>Register</h1>
  <div class="container">
    <van-form @submit="handleRegister">
      <van-field
        v-model="fullName"
        name="fullName"
        placeholder="Full name"
        :rules="[{ required: true, message: 'Full name is required' }]"
      />
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
      <van-field
        v-model="confirmPassword"
        type="password"
        name="confirmPassword"
        placeholder="Confirm password"
        :rules="[
          { required: true, message: 'Confirm password is required' },
          {
            validator: validatePasswordConfirmation,
            message: 'Passwords do not match',
          },
        ]"
      />
      <div class="button-container">
        <van-button round block type="primary" native-type="submit" width="80%">
          Register
        </van-button>
      </div>
      <span class="login-prompt" @click="goToLoginPage">
        Already have an account? Login
      </span>
    </van-form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { HOME_ROUTE, LOGIN_ROUTE } from "@/constants";
import { useStore } from "vuex";
import { Toast } from "vant";

export default defineComponent({
  data() {
    return {
      fullName: "",
      username: "",
      password: "",
      confirmPassword: "",
      store: useStore(),
    };
  },
  methods: {
    handleRegister() {
      this.store
        .dispatch("register", {
          displayName: this.fullName,
          username: this.username,
          password: this.password,
        })
        .then(() => this.$router.push(HOME_ROUTE))
        .catch((err) => Toast.fail(`Failed to register: ${err.message}`));
    },
    goToLoginPage() {
      this.$router.push(LOGIN_ROUTE);
    },
    validatePasswordConfirmation(val: string) {
      return val === this.password;
    },
  },
});
</script>
<style scoped lang="less">
.container .van-field {
  width: 80%;
}
.button-container {
  margin: 16px;
}
.login-prompt {
  padding: 0.5em 1em;
}
</style>
