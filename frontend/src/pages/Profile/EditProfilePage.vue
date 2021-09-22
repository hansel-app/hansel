<template>
  <div class="overlay">
    <Header title="Edit Profile" />
    <CircleAvatar :avatarUrl="placeholderAvatarUrl" radius="5" />
    <van-button class="edit-button"
      ><van-icon name="edit" size="40px"
    /></van-button>
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
    </van-form>
  </div>
  <ProfilePageBg />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import CircleAvatar from "@/components/CircleAvatar.vue";
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
    CircleAvatar,
    Header,
    ProfilePageBg,
  },
});
</script>

<style scoped lang="less">
.sub-header {
  text-align: left;
}
</style>
