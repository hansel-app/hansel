<template>
  <div class="overlay bring-to-front">
    <Header />
    <AvatarHeader
      :avatarUrl="placeholderAvatarUrl"
      :displayName="mockSelfUser.displayName"
      :username="mockSelfUser.username"
    />

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
import { CellGroup, Cell } from "vant";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import { EDIT_PROFILE_ROUTE } from "@/constants";
import AvatarHeader from "./AvatarHeader.vue";
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
    AvatarHeader,
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
