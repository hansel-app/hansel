<template>
  <NavBar left-arrow @click-left="goBack" />
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
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Cell, NavBar } from "vant";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import { EDIT_PROFILE_ROUTE } from "@/constants";
import AvatarHeader from "./AvatarHeader.vue";

export default defineComponent({
  data() {
    return {
      // TODO: replace with user's own data
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
      mockSelfUser,
      mockFriends,
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
    NavBar,
  },
});
</script>

<style scoped lang="less">
.sub-header {
  text-align: left;
}
</style>
