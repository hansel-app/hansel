<template>
  <NavBar />
  <AvatarHeader
    :avatarUrl="placeholderAvatarUrl"
    :displayName="mockSelfUser.displayName"
    :username="mockSelfUser.username"
  />
  <h3 class="sub-header">Friends</h3>
  <CellGroup>
    <FriendCell v-for="user in mockFriends" :key="user.id" :friend="user" />
  </CellGroup>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NavBar } from "vant";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import AvatarHeader from "./AvatarHeader.vue";
import FriendCell from "@/components/FriendCell.vue";

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
  components: {
    AvatarHeader,
    FriendCell,
    NavBar,
  },
});
</script>

<style scoped>
.sub-header {
  text-align: left;
}
</style>
