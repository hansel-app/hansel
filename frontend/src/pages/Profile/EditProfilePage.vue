<template>
  <div class="overlay">
    <Header title="Edit Profile" :backLink="PROFILE_ROUTE"/>
    <!-- <div class="avatar-container">
      <van-button class="edit-button overlay"
        ><van-icon name="edit" size="40px"
      /></van-button>
    </div> -->
    <van-form class="container">
      <ProfileAvatar :avatarUrl="placeholderAvatarUrl" />
      <van-uploader>
        <van-button round icon="edit"></van-button>
      </van-uploader>
      <van-field
        v-model="displayName"
        name="Display name"
        label="Display name"
        placeholder="Enter your name"
      />
      <van-button round type="primary" @click="editProfile"> Save </van-button>
    </van-form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import ProfileAvatar from "../../components/ProfileAvatar.vue";
import Header from "@/components/Header.vue";
import { useStore } from "vuex";
import { PROFILE_ROUTE } from "@/constants";

export default defineComponent({
  data() {
    return {
      PROFILE_ROUTE,
      // TODO: replace with user's own data
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
      mockSelfUser,
      mockFriends,
      username: mockSelfUser.username,
      displayName: mockSelfUser.displayName,
      store: useStore(),
    };
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
  },
});
</script>

<style scoped lang="less">
.sub-header {
  text-align: left;
}

.edit-button {
  position: absolute;
  bottom: 0;
}

.container {
  padding-top: 2em;
}

.van-field {
  padding: 3em 0;
}
</style>
