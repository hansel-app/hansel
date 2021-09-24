<template>
  <div class="background-gradient-saturated base flex">
    <div class="center">
      <WelcomeGem :gemColor="GemColor.PINK" />
      <div class="drawer">
        <h1>Register</h1>
        <van-form @submit="handleRegister">
          <van-field
            v-model="fullName"
            name="fullName"
            placeholder="Full name"
            :rules="[
              { required: true, message: 'Full name is required' },
              charCountValidator,
            ]"
          />
          <van-field
            v-model="username"
            name="Username"
            placeholder="Username"
            :rules="[
              { required: true, message: 'Username is required' },
              charCountValidator,
            ]"
          />
          <van-field
            v-model="password"
            type="password"
            name="Password"
            placeholder="Password"
            :rules="[
              { required: true, message: 'Password is required' },
              charCountValidator,
            ]"
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
              charCountValidator,
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
import { GemColor } from "@/interfaces";
import WelcomeGem from "./WelcomeGem.vue";

export default defineComponent({
  data() {
    return {
      fullName: "",
      username: "",
      password: "",
      confirmPassword: "",
      store: useStore(),
      GemColor,
      charCountValidator: {
        validator: (val: string) => {
          return val.length >= 4 && val.length <= 16;
        },
        message: "Must be between 4 to 16 characters",
      },
    };
  },
  components: {
    WelcomeGem,
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
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(8px);
  border-radius: 2em 2em 2em 2em;

  float: left;
  position: fixed;
  z-index: 1;
  left: 0;
  right: 0;
  padding: 0.5em;

  transition: max-height 0.3s;

  margin: 0 32px 32px 32px;
  padding-bottom: 48px;
  margin-top: 12px;
}

.flex {
  display: flex;
  justify-content: center;
}

.center {
  position: absolute;
  margin: auto;
  top: 5%;
}
</style>
