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
      <CellGroup :border="false" class="content container">
        <Cell title="Edit profile" is-link @click="goToEditProfile" />
        <Cell title="Friends" is-link />
      </CellGroup>
      <Cell class="logout" title="Logout" />
    </div>
    <div class="reverse-background-gradient">
      <SemiCircleBg />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Cell } from "vant";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import { EDIT_PROFILE_ROUTE } from "@/constants";
import ProfileAvatar from "../../components/ProfileAvatar.vue";
import Header from "@/components/Header.vue";
import SemiCircleBg from "@/components/SemiCircleBg.vue";

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
      console.log('hi');
      this.$router.push(EDIT_PROFILE_ROUTE);
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

.bring-to-front {
  z-index: 2;
}

.header {
  text-align: center !important;
  margin: 1em 0 1em 0;
}

.sub-header {
  text-align: center !important;
  padding-bottom: 4em;
  color: gray;
  font-weight: lighter !important;
}

.logout {
  color: @red;
  padding-bottom: 2em;
}

.stack-container {
  z-index: -1;
}

.reverse-background-gradient {
  background: linear-gradient(to bottom, #c0c2e5, #e43d9700);
  height: 100vh;
  position: relative;
  z-index: -2;
}
</style>
