<template>
  <div class="overlay">
    <Header title="Edit Profile" />
    <div class="avatar-container">
      <ProfileAvatar :avatarUrl="self.avatar" />
      <van-button class="edit-button overlay"
        ><van-icon name="edit" size="40px"
      /></van-button>
    </div>
    <van-form>
      <van-field
        v-model="username"
        name="Username"
        label="Username"
        placeholder="username"
        readonly
        error-message="Username cannot be changed"
      />
      <van-field
        v-model="displayName"
        name="Display name"
        label="Display name"
        placeholder="Display name"
      />
      <van-button round type="primary" @click="editProfile"> Save </van-button>
    </van-form>
  </div>
  <ProfilePageBg />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ProfileAvatar from "./ProfileAvatar.vue";
import ProfilePageBg from "./ProfilePageBg.vue";
import Header from "@/components/Header.vue";
import { User } from "@/interfaces/user.ts";
import { useStore } from "vuex";

export default defineComponent({
  data() {
    return {
      store: useStore(),
    };
  },
  computed: {
    self(): User {
      return this.store.state.user.self;
    },
    username(): string {
      return this.self.username;
    },
    displayName(): string {
      return this.self.displayName;
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    editProfile() {
      this.store.dispatch("editOwnProfile", {
        newDisplayName: this.displayName,
      });
    },
  },
  components: {
    ProfileAvatar,
    Header,
    ProfilePageBg,
  },
});
</script>

<style scoped lang="less">
.sub-header {
  text-align: left;
}
.avatar-container {
  position: relative;
  width: 5rem;
  left: 0;
  right: 0;
  margin-left: auto;
  margin-right: auto;
}
.edit-button {
  position: absolute;
  bottom: 0;
}
</style>
