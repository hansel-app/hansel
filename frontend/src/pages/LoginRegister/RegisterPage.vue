<template>
  <div class="background-gradient">
    <div class="drawer">
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
            <van-button
              round
              block
              type="primary"
              native-type="submit"
              width="80%"
            >
              Register
            </van-button>
          </div>
          <span>
            Already have an account?
          </span>
          <span class="login-prompt" @click="goToLoginPage">
            Login
          </span>
        </van-form>
      </div>
    </div>
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
  color: @primary-color;
}

.drawer {
  color: black;
  background-color: rgba(255, 255, 255, 0.473);
  backdrop-filter: blur(5px);
  border-radius: 2em 2em 2em 2em;

  float: left;
  position: fixed;
  z-index: 1;
  left: 0;
  right: 0;
  padding: 0.5em;

  transition: max-height 0.3s;

  position: absolute;
  top: 50%;
  transform: translate(0, -50%);

  margin-right: 5vw;
  margin-left: 5vw;
}
</style>
