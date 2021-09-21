<template>
  <h1>Register</h1>
  <div class="container">
    <van-form @submit="onSubmit">
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
        type="confirmPassword"
        name="confirmPassword"
        placeholder="Confirm password"
        :rules="[{ required: true, message: 'Confirm password is required' }]"
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
import { defineComponent } from 'vue';
import { LOGIN_ROUTE } from "@/constants";

export default defineComponent({
  methods: {
    goToLoginPage() {
      this.$router.push(LOGIN_ROUTE);
    }
  }
})
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
