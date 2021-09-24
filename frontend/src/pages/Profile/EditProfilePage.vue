<template>
  <div class="overlay background-gradient">
    <Header title="Edit Profile" :backLink="PROFILE_ROUTE" />
    <van-form class="container">
      <Uploader v-model="fileList" :max-count="1" :preview-size="150">
        <template #default>
          <van-image
            height="150px"
            width="150px"
            :src="`data:image/png;base64,${self.avatar}`"
            round
            fit="cover"
          ></van-image>
        </template>
      </Uploader>
      <van-field
        v-model="displayName"
        name="Display name"
        label="Name"
        placeholder="Enter your name"
      />
      <van-button round type="primary" @click="editProfile"> Save </van-button>
    </van-form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Header from "@/components/Header.vue";
import { User } from "@/interfaces/user.ts";
import { useStore } from "vuex";
import { PROFILE_ROUTE } from "@/constants";
import { Uploader, UploaderFileListItem, Toast } from "vant";

export default defineComponent({
  setup() {
    const store = useStore();
    return { store };
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
    editProfile() {
      let updateProfile;

      if (this.fileList.length != 0) {
        updateProfile = () =>
          this.store.dispatch("editOwnProfile", {
            newAvatar: this.fileList[0].file,
            newDisplayName: this.displayName,
          });
      } else {
        updateProfile = () =>
          this.store.dispatch("editOwnProfile", {
            newDisplayName: this.displayName,
          });
      }
      updateProfile()
        .then(() => {
          Toast.success("Updated profile!");
          this.$router.replace(PROFILE_ROUTE);
        })
        .catch((err) => Toast.fail(`Failed to update profile: ${err.message}`));
    },
  },
  data() {
    return {
      PROFILE_ROUTE,
      fileList: [] as UploaderFileListItem[],
    };
  },
  components: {
    Header,
    Uploader,
  },
});
</script>

<style scoped lang="less">
.container {
  padding-top: 2em;
}

.van-field {
  padding: 3em 0;
}
</style>
