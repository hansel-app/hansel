<template>
  <div v-click-away="clickAway" v-touch:swipe.left="clickAway">
    <div class="menu-icon" @click="toggleMenu">
      <img :src="HamburgerMenuIcon" />
    </div>
    <div class="menu" :style="{ width: menuWidth }">
      <div class="menu-contents" v-if="!collapsed">
        <div class="avatar">
          <CircleAvatar
            :avatarUrl="self.avatar"
            :radius="2.5"
            @click="goToProfile"
          >
            Profile
          </CircleAvatar>
        </div>
        <div class="menu" :style="{ width: menuWidth }">
          <div class="menu-contents" v-if="!collapsed">
            <div class="avatar">
              <CircleAvatar
                :avatarUrl="self.avatar"
                :radius="2.5"
                @click="goToProfile"
              >
                Profile
              </CircleAvatar>
            </div>
            <div class="links">
              <van-row tabindex="1" @click="goToLogs">
                <p>Gem logs</p>
              </van-row>
              <van-row tabindex="2" @click="goToAddFriends">
                <p>Add friends</p>
              </van-row>
              <van-row tabindex="3" @click="goToFriendRequests">
                <p>Friend requests</p>
                <van-col v-if="requestCount > 0">
                  <Notification :requestCount="requestCount"/>
                </van-col>
              </van-row>
              <van-row class="log-out" tabindex="4" @click="logOut">
                <p>Log out</p>
              </van-row>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {
  PROFILE_ROUTE,
  ADD_FRIENDS_ROUTE,
  FRIEND_REQUESTS_ROUTE,
  GEM_LOGS_ROUTE,
  LOGIN_ROUTE,
} from "@/constants";
import { defineComponent, ref, computed } from "vue";
import { useStore } from "vuex";
import CircleAvatar from "@/components/CircleAvatar.vue";
import HamburgerMenuIcon from "@/assets/icons/menu-hamburger.svg";
import Notification from "@/components/Notification.vue";
import { Row } from "vant";

export default defineComponent({
  setup() {
    const collapsed = ref(true);
    const toggleMenu = () => (collapsed.value = !collapsed.value);
    const MENU_WIDTH_EXPANDED = 14;
    const MENU_WIDTH_COLLAPSED = 0;
    const menuWidth = computed(
      () => `${collapsed.value ? MENU_WIDTH_COLLAPSED : MENU_WIDTH_EXPANDED}em`
    );

    return {
      collapsed,
      toggleMenu,
      menuWidth,
      HamburgerMenuIcon,
    };
  },
  components: {
    "van-row": Row,
    CircleAvatar,
    Notification,
  },
  props: {
    requestCount: {
      type: Number,
      required: true,
    },
  },
  data() {
    return { store: useStore() };
  },
  computed: {
    self() {
      return this.store.state.user.self;
    },
  },
  methods: {
    goToProfile() {
      this.$router.push(PROFILE_ROUTE);
    },
    goToLogs() {
      this.$router.push(GEM_LOGS_ROUTE);
    },
    goToAddFriends() {
      this.$router.push(ADD_FRIENDS_ROUTE);
    },
    goToFriendRequests() {
      this.$router.push(FRIEND_REQUESTS_ROUTE);
    },
    clickAway() {
      if (!this.collapsed) {
        this.toggleMenu();
      }
    },
    logOut() {
      this.store.dispatch("logout").then(() => this.$router.push(LOGIN_ROUTE));
    },
  },
});
</script>
<style>
:root {
  --menu-bg-color: #5f4bae;
  --menu-item-active: #9c8be0;
}
</style>
<style scoped lang="less">
.menu,
.menu-icon {
  color: white;
  background-color: var(--menu-bg-color);

  float: left;
  position: fixed;
  z-index: 1;
  left: 0;
  display: flex;
}
.menu {
  top: 0;
  bottom: 0;
  transition: 0.3s ease;
  flex-direction: column;
}

.menu-icon {
  top: 2em;
  color: white;
  background-color: var(--menu-bg-color);
  padding: 0.5em;
  border-radius: 0 0.75em 0.75em 0;
  vertical-align: middle;
  text-align: center;
}

.menu-contents {
  padding: 5em 1em;
}

.avatar {
  margin-bottom: 2em;
}

.van-row {
  border-radius: 1em;
  padding: 0 1em;
  white-space: nowrap;
  overflow: hidden;
}

.van-row:hover,
.van-row:focus {
  background-color: var(--menu-item-active);
}

.log-out {
  margin-top: 5em;
  color: @gray-3;
}
</style>
