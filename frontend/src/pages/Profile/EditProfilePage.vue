<template>
  <div class="overlay">
    <Header title="Edit Profile" />
    <div class="avatar-container">
      <ProfileAvatar :avatarUrl="placeholderAvatarUrl" />
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
      />
      <van-field
        v-model="displayName"
        name="Display name"
        label="Display name"
        placeholder="Display name"
      />
      <van-button round type="primary"> Save </van-button>
    </van-form>
  </div>
  <ProfilePageBg />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import ProfileAvatar from "./ProfileAvatar.vue";
import ProfilePageBg from "./ProfilePageBg.vue";
import Header from "@/components/Header.vue";

export default defineComponent({
  data() {
    return {
      // TODO: replace with user's own data
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
      mockSelfUser,
      mockFriends,
      username: mockSelfUser.username,
      displayName: mockSelfUser.displayName,
    };
  },
  methods: {
    goBack() {
      this.$router.back();
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
