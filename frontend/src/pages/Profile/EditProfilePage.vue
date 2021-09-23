<template>
  <div class="overlay">
    <Header title="Edit Profile" :backLink="PROFILE_ROUTE"/>
    <van-form class="container">
      <Uploader
          v-model="fileList"
          :max-count="1"
          :preview-size="150"
      >
        <template #default>
          <van-image height="150px" width="150px" :src="self.avatar"></van-image>
        </template>
      </Uploader>
      <van-field
        v-model="displayName"
        name="Display name"
        label="self.displayName"
        placeholder="Enter your name"
      />
      <van-button round type="primary" @click="editProfile"> Save </van-button>
    </van-form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Header from "@/components/Header.vue";
import { useStore } from "vuex";
import { PROFILE_ROUTE } from "@/constants";
import { Uploader, UploaderFileListItem, Toast } from "vant";
import { User } from "@/interfaces";

export default defineComponent({
  setup() {
    const store = useStore();
    const self: User = store.state.user.self;
    const displayName: string = self.displayName;
    return { store, self, displayName };
  },
  methods: {
    editProfile() {
      let updateProfile;

      if(this.fileList.length != 0) {
        updateProfile = () => this.store.dispatch("editOwnProfile", {
          newAvatar: this.fileList[0].content,
          newDisplayName: this.displayName,
        });
      } else {
        updateProfile = () => this.store.dispatch("editOwnProfile", {
          newDisplayName: this.displayName,
        });
      }
    updateProfile().
      then(() => {
        Toast.success("Updated profile!");
        this.$router.replace(PROFILE_ROUTE);
      })
      .catch((err) => Toast.fail(`Failed to update profile: ${err}`));
    },
  },
  data() {
    return {
      PROFILE_ROUTE,
      // displayName: this.self.displayName as string,
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
