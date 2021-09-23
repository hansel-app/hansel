<template>
  <div class="profile-page">
    <div class="overlay bring-to-front container">
      <Header />
      <ProfileAvatar :avatarUrl="placeholderAvatarUrl" />
      <div class="header">
        {{ mockSelfUser.displayName }}
      </div>
      <div class="sub-header">
        @{{ mockSelfUser.username }}
      </div>
      <CellGroup :border="false" class="content">
        <Cell title="Edit profile" is-link @click="goToEditProfile" />
        <Cell title="Friends" is-link />
      </CellGroup>
      <Cell class="logout" title="Logout" />
    </div>
    <ProfilePageBg />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Cell } from "vant";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import { EDIT_PROFILE_ROUTE } from "@/constants";
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
      color: "#ffffff",
    };
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

.header {
  text-align: center !important;
  margin: 2em 0 1em 0;
}

.sub-header {
  text-align: center !important;
  padding-bottom: 2em;
  color: gray;
}

.logout {
  color: red;
}
</style>
