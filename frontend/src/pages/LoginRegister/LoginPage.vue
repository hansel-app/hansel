<template>
  <div class="background-gradient-saturated base flex">
    <div class="center">
      <WelcomeGem :gemColor="GemColor.BLUE" />
      <div class="drawer">
        <h1>Login</h1>
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
          <div class="button-container content">
            <van-button
              round
              block
              type="primary"
              native-type="submit"
              size="large"
            >
              Login
            </van-button>
          </div>
          <span class="register" @click="goToRegisterPage">
            Sign up
          </span>
        </van-form>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { HOME_ROUTE, REGISTER_ROUTE } from "@/constants";
import { useStore } from "vuex";
import { Toast } from "vant";
import WelcomeGem from "./WelcomeGem.vue";
import { GemColor } from "@/interfaces";

export default defineComponent({
  data() {
    return {
      username: "",
      password: "",
      store: useStore(),
      GemColor,
    };
  },
  methods: {
    handleLogin() {
      this.store
        .dispatch("login", {
          username: this.username,
          password: this.password,
        })
        .then(() => this.$router.push(HOME_ROUTE))
        .catch(() => Toast.fail(`Failed to log in`));
    },
    goToRegisterPage() {
      this.$router.push(REGISTER_ROUTE);
    },
  },
  components: {
    WelcomeGem,
  },
});
</script>

<style scoped lang="less">
.van-field {
  width: 80%;
}
.button-container {
  margin: 16px;
}
.register {
  padding: 0.5em 1em;
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
  top: 10%;
}
</style>
