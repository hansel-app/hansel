<template>
  <div class="profile-page">
    <div class="overlay container">
      <Header title="" />
      <ProfileAvatar :avatarUrl="self.avatar" />
      <div class="header">
        {{ self.displayName }}
      </div>
      <div class="sub-header">@{{ self.username }}</div>
      <CellGroup :border="false" class="content container">
        <Cell title="Edit profile" is-link @click="goToEditProfile" />
        <Cell title="Friends" is-link @click="goToFriendList" />
      </CellGroup>
    </div>
    <div class="reverse-background-gradient">
      <SemiCircleBg />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useStore } from "vuex";
import { CellGroup, Cell } from "vant";
import { EDIT_PROFILE_ROUTE, FRIEND_LIST_ROUTE } from "@/constants";
import { User } from "@/interfaces";
import ProfileAvatar from "@/components/ProfileAvatar.vue";
import Header from "@/components/Header.vue";
import SemiCircleBg from "@/components/SemiCircleBg.vue";

export default defineComponent({
  data() {
    return {
      color: "#ffffff",
      store: useStore(),
    };
  },
  computed: {
    self(): User {
      return this.store.state.user.self;
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    goToEditProfile() {
      this.$router.push(EDIT_PROFILE_ROUTE);
    },
    goToFriendList() {
      this.$router.push(FRIEND_LIST_ROUTE);
    },
  },
  components: {
    ProfileAvatar,
    CellGroup,
    Cell,
    Header,
    SemiCircleBg,
  },
});
</script>

<style scoped lang="less">
.sub-header {
  text-align: left;
}

.header {
  text-align: center !important;
  margin: 1.2em 0 1em 0;
}

.sub-header {
  text-align: center !important;
  padding-bottom: 4em;
  color: gray;
  font-weight: lighter !important;
}
</style>
