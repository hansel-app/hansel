<template>
  <!-- TODO: some onboarding screens for user to set displayName and avatar? -->
  <h1>Register</h1>
  <div class="container">
    <van-form @submit="register">
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
    </van-form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      username: "",
      password: "",
      confirmPassword: "",
    };
  },
  methods: {
    register() {
      if (this.password !== this.confirmPassword) {
        // TODO: show some error here.
        return;
      }
      this.$store.dispatch("regsiter", {
        username: this.username,
        password: this.password,
      });
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
</style>
