<template>
  <div class="overlay bring-to-front">
    <Header />
    <ProfileAvatar :avatarUrl="self.avatar" />
    <h2>{{ self.displayName }}</h2>
    <h3>{{ "@" + self.username }}</h3>
    <CellGroup>
      <Cell title="Edit profile" is-link @click="goToEditProfile" />
      <Cell title="Friends" is-link />
    </CellGroup>
    <Cell title="Logout" />
  </div>
  <ProfilePageBg />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useStore } from "vuex";
import { CellGroup, Cell } from "vant";
import { EDIT_PROFILE_ROUTE } from "@/constants";
import { User } from "@/interfaces";
import ProfileAvatar from "./ProfileAvatar.vue";
import ProfilePageBg from "./ProfilePageBg.vue";
import Header from "@/components/Header.vue";

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
  },
  components: {
    ProfileAvatar,
    CellGroup,
    Cell,
    Header,
    ProfilePageBg,
  },
});
</script>

<style scoped lang="less">
.sub-header {
  text-align: left;
}

.bring-to-front {
  z-index: 2;
}
</style>
